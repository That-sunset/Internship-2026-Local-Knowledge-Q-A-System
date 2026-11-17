# API Documentation

Base URL: `http://127.0.0.1:8080`

---

## GET

### Getchat
Retrieves chat data from `record.db` using `recordid`.

* **Endpoint:** `GET /api/chat/:recordid`
* **Database:** `record.db`
* **Frontend Usage:**
  ```typescript
  api.get(`/api/chat/${recordId}`).then(function (data) {
    console.log(data)
  })
  ```

---

### Getrecord
Retrieves record data from `record.db`.

* **Endpoint:** `GET /api/record`
* **Database:** `record.db`
* **Frontend Usage:**
  ```typescript
  api.get(`/api/record/`).then(function (data) {
    console.log(data)
  })
  ```

---

## POST

### Post
Sends new data to `record.db`.

* **Endpoint:** `POST /api/post`
* **Database:** `record.db`
* **Frontend Usage:**
  ```typescript
  api.post('/api/post', payload).then(function (data) {
    console.log(data)
  })
  ```
* **Request Body:**
  ```json
  {
    "text": "<text>"
  }
  ```

---

## DELETE

### Deletebyid
Finds and deletes a record in `record.db` using `recordid`.

* **Endpoint:** `DELETE /api/record/:recordid`
* **Database:** `record.db`
* **Frontend Usage:**
  ```typescript
  api.delete(`/api/record/${recordId}`).then(function (data) {
    console.log(data)
  })
  ```
