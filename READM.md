# 🚀 Backend Performance Benchmark (Go vs Node.js vs Java)

This project compares the performance of **Go, Node.js, and Java (Spring Boot)** APIs using **k6 load testing** with **MongoDB** as the database.

It evaluates:

- 📖 Read performance (`GET /users`)
- ✍️ Write performance (`POST /users`)
- 📦 Empty database scenario (`GET /events/recent`)
- 🗄️ Large dataset scenario (~5000+ records)

---

## 🧪 Tech Stack

| Layer        | Technology                     |
|--------------|-------------------------------|
| Backend 1    | Go (Gin + MongoDB)            |
| Backend 2    | Node.js (Express + MongoDB)   |
| Backend 3    | Java (Spring Boot + MongoDB)  |
| Database     | MongoDB                       |
| Load Testing | k6                            |

---

## 📁 Project Structure

```

backend-performace-benchmark/
├── go-backend/
├── node-backend/
├── spring-backend/
├── benchmark/
│   ├── read-test.js
│   ├── write-test.js
│   ├── empty-test.js
|   ├── results/
│       ├── go-read.json
│       ├── node-read.json
│       ├── java-read.json
│       ├── go-write.json
│       ├── node-write.json
│       ├── java-write.json
│       ├── empty.json

````

---

## ⚙️ How to Run Backend

### 🐹 Go Service
```bash
go run main.go
````

### 🟩 Node.js Service

```bash
npm install
node app.js
```

### ☕ Java Spring Boot

```bash
mvn spring-boot:run
```

---

## 🧪 How to Run Load Tests (k6)

### 📖 Read Test

```bash
k6 run read-test.js --out json=go-read.json
k6 run read-test.js --out json=node-read.json
k6 run read-test.js --out json=java-read.json
```

---

### ✍️ Write Test

```bash
k6 run write-test.js --out json=go-write.json
k6 run write-test.js --out json=node-write.json
k6 run write-test.js --out json=java-write.json
```

---

### 📦 Empty DB Test

```bash
k6 run empty-test.js --out json=go-empty.json
k6 run empty-test.js --out json=node-empty.json
k6 run empty-test.js --out json=java-empty.json
```

---

## 📊 Summary Results

### 📖 READ PERFORMANCE (5000+ MongoDB records)

| Stack      | Avg (ms) | P95 (ms) | Req/sec | Result             |
| ---------- | -------- | -------- | ------- | ------------------ |
| 🐹 Go      | 1359     | 2676     | 41.9    | 🥇 Best throughput |
| 🟩 Node.js | 8141     | 12363    | 9.0     | ❌ Slow under load  |
| ☕ Java     | 2282     | 4730     | 29.4    | ⚖️ Stable          |

---

### ✍️ WRITE PERFORMANCE

| Stack      | Avg (ms) | P95 (ms) | Req/sec | Result      |
| ---------- | -------- | -------- | ------- | ----------- |
| 🐹 Go      | 25       | 84       | 95.9    | 🥇 Fastest  |
| 🟩 Node.js | 59       | 488      | 92.7    | 🥈 Balanced |
| ☕ Java     | 152      | 1925     | 83.5    | 🥉 Slowest  |

---

### 📦 EMPTY DB PERFORMANCE

| Stack      | Avg (ms) | P95 (ms) | Req/sec | Notes             |
| ---------- | -------- | -------- | ------- | ----------------- |
| 🐹 Go      | 21       | 107      | 96.9    | 🥇 Best           |
| 🟩 Node.js | 31       | 218      | 95.7    | Stable            |
| ☕ Java     | 365      | 4512     | 71.2    | High tail latency |

---

## 📉 Key Observations

* 🐹 Go performs best overall in latency + throughput
* 🟩 Node.js performs well but degrades under heavy read load
* ☕ Java is stable but slower in most cases
* ⚠️ Write tests show issues in status validation (check API responses)

---

## 🚨 Important Notes

* Ensure correct HTTP status codes (200 / 201)
* Use MongoDB indexing for better read performance
* Enable connection pooling in all services
* Fix failing k6 assertions (some write tests show 0% success)

---

## 📊 Conclusion

* 🥇 **Best for high performance APIs:** Go (Gin)
* 🥈 **Best for balanced development:** Node.js
* 🥉 **Best for enterprise structure:** Java (Spring Boot)

---


