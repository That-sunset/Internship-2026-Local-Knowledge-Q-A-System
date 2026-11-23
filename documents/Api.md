# API Documentation

Base URL: `http://127.0.0.1:8080`

---

## GET

### Getchat
Retrieves all chat messages from `chat.db` (`chats` table) belonging to a session using `id` (`session_id`).

- **Endpoint:** `GET /api/chat/:id`
- **Database:** `chat.db`
- **Frontend Usage:**
  ```typescript
  api.get(`/api/chat/${sessionId}`).then(function (data) {
    console.log(data)
  })
  ```

---

### Getrecord
Retrieves all session records from `chat.db` (`records` table).

- **Endpoint:** `GET /api/records`
- **Database:** `chat.db`
- **Frontend Usage:**
  ```typescript
  api.get('/api/records').then(function (data) {
    console.log(data)
  })
  ```

---

## POST

### Postrecord
Creates a new chat session in `chat.db` (`records` table).

- **Endpoint:** `POST /api/records`
- **Database:** `chat.db`
- **Frontend Usage:**
  ```typescript
  api.post('/api/records', {}).then(function (data) {
    console.log(data)
  })
  ```
- **Request Body:**
  ```json
  {}
  ```

---

### Postchat
Adds a new message (`user` or `model`) to a session in `chat.db` (`chats` table) using `id` (`session_id`).

- **Endpoint:** `POST /api/chat/:id`
- **Database:** `chat.db`
- **Frontend Usage:**
  ```typescript
  api.post(`/api/chat/${sessionId}`, payload).then(function (data) {
    console.log(data)
  })
  ```
- **Request Body:**
  ```json
  {
    "content": "<text>",
    "source": "user"
  }
  ```

---

## DELETE

### Deletebyid
Finds and deletes a session record in `chat.db` (`records` table) using `id`.

- **Endpoint:** `DELETE /api/records/:id`
- **Database:** `chat.db`
- **Frontend Usage:**
  ```typescript
  api.delete(`/api/records/${recordId}`).then(function (data) {
    console.log(data)
  })
  ```