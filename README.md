# Go URL Shortener 🚀

A minimal URL shortener built with Go, Redis, and Gorilla Mux. It allows you to shorten URLs and redirect users using a simple and scalable backend.

---

## 🎓 Features

- Shorten any valid URL to a custom short code
- Redirect using `GET /r/{id}`
- JSON-based API
- Stores data in Redis
- Simple and clean project structure

---

## 🌐 API Endpoints

### `GET /`
**Description:** Health check or welcome message.

**Response:**
```text
Welcome to Go URL Shortener 🚀
```

### `POST /shorten`
**Description:** Shortens a given URL.

**Request Body:**
```json
{
  "url": "https://example.com"
}
```

**Response:**
```json
{
  "short_url": "http://localhost:8080/r/a1B2c3"
}
```

### `GET /r/{id}`
**Description:** Redirects to the original URL using the short ID.

**Response:**
- `302 Found` if the ID exists
- `404 Not Found` if not

---

## 🚀 Getting Started

### Prerequisites
- Go 1.18+
- Redis (use [Memurai](https://www.memurai.com/) on Windows)

### Run Redis (Memurai on Windows)
Ensure Redis is running on `localhost:6379`
```bash
memurai-cli
127.0.0.1:6379> ping
PONG
```

### Install Dependencies
```bash
go mod tidy
```

### Run the Server
```bash
go run main.go
```

Server runs at: `http://localhost:8080`

---

## 🌟 Project Structure
```bash
go-url-shortener/
├── go.mod
├── main.go
├── handlers/
│   ├── home.go
│   ├── shorten.go
│   └── redirect.go
├── storage/
│   ├── redis.go
│   └── client.go
├── utils/
│   └── id.go
```

---

## 💡 Tech Stack
- Go (Golang)
- Gorilla Mux (routing)
- Redis (storage)
- JSON for communication

---

## 📅 Future Improvements
- Add rate limiting middleware
- Click analytics and usage stats
- Docker support
- Authentication

---

## ✨ Credits
Made with ❤️ by [Your Name Here] for learning and fun.

---

## 🛌 License
MIT License

