# 🧾 Receipt Processor API

A web service that processes shopping receipts and awards points based on defined rules. Built in **Go**, containerized using **Docker**, and fully compliant with the provided API spec.

---

## 📦 Features

- Processes receipts via `POST /receipts/process`
- Retrieves earned points via `GET /receipts/{id}/points`
- In-memory storage (no database needed)
- Follows exact rules as specified in the prompt
- Thread-safe (uses mutex for concurrency)
- Clean project structure: `handlers`, `models`, `utils`, `store`

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/) 1.20+ (optional if using Docker)
- [Docker](https://www.docker.com/) (✅ recommended)

---

### 🔧 Run with Docker

```bash
# 1. Build the image
docker build -t receipt-processor .

# 2. Run the container
docker run -p 8080:8080 receipt-processor
````

> Service runs on: `http://localhost:8080`

---

## 🔌 API Endpoints

### 📥 POST `/receipts/process`

**Description:** Accepts a receipt in JSON format and returns a unique ID.

**Request Example:**

```bash
curl -X POST http://localhost:8080/receipts/process \
  -H "Content-Type: application/json" \
  -d @testdata/simple-receipt.json
```

**Response:**

```json
{ "id": "abc123-..." }
```

---

### 📤 GET `/receipts/{id}/points`

**Description:** Returns the number of points earned for a previously processed receipt.

**Request Example:**

```bash
curl http://localhost:8080/receipts/id/points
```

**Response:**

```json
{ "points": 28 }
```

---

## 🧠 Rules Used for Scoring

* +1 point per alphanumeric character in the retailer name
* +50 points if total is a whole dollar
* +25 points if total is a multiple of 0.25
* +5 points for every 2 items
* +\[ceil(price × 0.2)] for item descriptions with length divisible by 3
* +6 points if purchase day is odd
* +10 points if time is between 2:00pm and 4:00pm

📊 The terminal will also log a full breakdown of how points were calculated.

---

## 🧪 Sample Test Data

Test receipts are available in the `testdata/` directory:

* `mnm.json`
* `morning-receipt.json`
* `simple-receipt.json`
* `target.json`

---

## 🗂 Project Structure

```
receipt-processor/
├── main.go               # App entry point
├── handlers/             # API route logic
├── models/               # Receipt + Item structs
├── store/                # In-memory data store (with mutex)
├── utils/                # Point calculation logic
├── testdata/             # Example JSON receipts
├── go.mod / go.sum       # Module and dependencies
├── Dockerfile            # Container setup
└── README.md             # You're reading it!
```

---

## 📄 Notes

* No database needed — uses in-memory storage
* All dependencies are managed with Go modules
* Point calculation is logged in the console for transparency