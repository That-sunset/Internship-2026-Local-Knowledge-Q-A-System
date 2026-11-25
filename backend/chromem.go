package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/philippgille/chromem-go"
)

var vectorCollection *chromem.Collection

// InitVectorDB initializes persistent chromem-go storage using Ollama's nomic-embed-text.
func InitVectorDB() {
	persistentDB, err := chromem.NewPersistentDB("./vector_db", false)
	if err != nil {
		log.Fatalf("failed to initialize persistent chromem DB: %v", err)
	}

	embedFunc := chromem.NewEmbeddingFuncOllama("nomic-embed-text", "http://127.0.0.1:11434/api")

	col, err := persistentDB.GetOrCreateCollection("knowledge-base", nil, embedFunc)
	if err != nil {
		log.Fatalf("failed to get or create vector collection: %v", err)
	}

	vectorCollection = col
	log.Printf("Vector DB initialized (%d existing chunks loaded)", vectorCollection.Count())
}

// ingestFileToVectorDB follows the flowchart:
// 1. Check .csv vs .txt -> convert CSV rows to sentences or read plain text
// 2. Chunk into pieces
// 3. Embed via nomic-embed-text & store in chromem-go
func ingestFileToVectorDB(ctx context.Context, filename string, rawBytes []byte) (int, error) {
	if vectorCollection == nil {
		return 0, fmt.Errorf("vector database is not initialized")
	}

	ext := strings.ToLower(filepath.Ext(filename))
	var plainText string

	switch ext {
	case ".csv":
		converted, err := csvToSentences(rawBytes)
		if err != nil {
			return 0, fmt.Errorf("failed to parse CSV %s: %w", filename, err)
		}
		plainText = converted
	case ".txt":
		plainText = string(rawBytes)
	default:
		return 0, fmt.Errorf("unsupported file type: %s (only .txt and .csv are supported)", ext)
	}

	chunks := chunkText(plainText, 500)
	if len(chunks) == 0 {
		return 0, fmt.Errorf("no valid text content found in %s", filename)
	}

	docs := make([]chromem.Document, 0, len(chunks))
	for i, chunk := range chunks {
		docs = append(docs, chromem.Document{
			ID:      fmt.Sprintf("%s-%s-%d", filename, uuid.NewString(), i),
			Content: chunk,
			Metadata: map[string]string{
				"source": filename,
			},
		})
	}

	if err := vectorCollection.AddDocuments(ctx, docs, runtime.NumCPU()); err != nil {
		return 0, fmt.Errorf("failed to embed and store documents: %w", err)
	}

	return len(docs), nil
}

// csvToSentences converts each CSV row into a readable sentence using column headers if present.
func csvToSentences(data []byte) (string, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	reader.FieldsPerRecord = -1 // allow variable column counts safely

	var rows [][]string
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		rows = append(rows, rec)
	}

	if len(rows) == 0 {
		return "", nil
	}

	headers := rows[0]
	var sentences []string

	// If there is a header row + data rows, pair header: value per row
	if len(rows) > 1 {
		for _, row := range rows[1:] {
			var parts []string
			for i, val := range row {
				val = strings.TrimSpace(val)
				if val == "" {
					continue
				}
				if i < len(headers) && strings.TrimSpace(headers[i]) != "" {
					parts = append(parts, fmt.Sprintf("%s is %s", strings.TrimSpace(headers[i]), val))
				} else {
					parts = append(parts, val)
				}
			}
			if len(parts) > 0 {
				sentences = append(sentences, strings.Join(parts, ", ")+".")
			}
		}
	} else {
		// Single-row CSV fallback
		var parts []string
		for _, val := range headers {
			if strings.TrimSpace(val) != "" {
				parts = append(parts, strings.TrimSpace(val))
			}
		}
		if len(parts) > 0 {
			sentences = append(sentences, strings.Join(parts, ", ")+".")
		}
	}

	return strings.Join(sentences, "\n\n"), nil
}

// chunkText splits text by paragraphs/lines first, grouping pieces up to maxChars.
func chunkText(text string, maxChars int) []string {
	normalized := strings.ReplaceAll(text, "\r\n", "\n")
	paragraphs := strings.Split(normalized, "\n\n")

	var chunks []string
	var current strings.Builder

	for _, p := range paragraphs {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		// If a single paragraph is longer than maxChars, split by lines/words
		if len(p) > maxChars {
			if current.Len() > 0 {
				chunks = append(chunks, strings.TrimSpace(current.String()))
				current.Reset()
			}
			words := strings.Fields(p)
			var piece strings.Builder
			for _, w := range words {
				if piece.Len()+len(w)+1 > maxChars && piece.Len() > 0 {
					chunks = append(chunks, strings.TrimSpace(piece.String()))
					piece.Reset()
				}
				if piece.Len() > 0 {
					piece.WriteByte(' ')
				}
				piece.WriteString(w)
			}
			if piece.Len() > 0 {
				chunks = append(chunks, strings.TrimSpace(piece.String()))
			}
			continue
		}

		if current.Len()+len(p)+2 > maxChars && current.Len() > 0 {
			chunks = append(chunks, strings.TrimSpace(current.String()))
			current.Reset()
		}
		if current.Len() > 0 {
			current.WriteString("\n\n")
		}
		current.WriteString(p)
	}

	if current.Len() > 0 {
		chunks = append(chunks, strings.TrimSpace(current.String()))
	}

	return chunks
}

// queryVectorContext searches chromem-go for the topK most relevant chunks.
func queryVectorContext(ctx context.Context, query string, topK int) ([]string, error) {
	if vectorCollection == nil {
		return nil, nil
	}

	count := vectorCollection.Count()
	if count == 0 {
		return nil, nil
	}
	if topK > count {
		topK = count
	}

	results, err := vectorCollection.Query(ctx, query, topK, nil, nil)
	if err != nil {
		return nil, err
	}

	var contexts []string
	for _, res := range results {
		// Filter out completely unrelated chunks
		if res.Similarity >= 0.35 {
			source := res.Metadata["source"]
			if source != "" {
				contexts = append(contexts, fmt.Sprintf("[%s] %s", source, res.Content))
			} else {
				contexts = append(contexts, res.Content)
			}
		}
	}

	return contexts, nil
}
