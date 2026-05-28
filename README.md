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
|   ├── output-images/
|   ├── results/
│       ├── go-read.json
│       ├── node-read.json
│       ├── java-read.json
│       ├── go-write.json
│       ├── node-write.json
│       ├── java-write.json
│       ├── go-empty.json
│       ├── node-empty.json
│       ├── java-empty.json

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
k6 run read-test.js  --summary-export=go-read.json
k6 run read-test.js  --summary-export=node-read.json
k6 run read-test.js  --summary-export=java-read.json
```

---

### ✍️ Write Test

```bash
k6 run write-test.js  --summary-export=go-write.json
k6 run write-test.js  --summary-export=node-write.json
k6 run write-test.js  --summary-export=java-write.json
```

---

### 📦 Empty DB Test

```bash
k6 run empty-test.js  --summary-export-empty.json
k6 run empty-test.js  --summary-export=node-empty.json
k6 run empty-test.js  --summary-export=java-empty.json
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

## 📊 📁 Output Benchmark Visual Results

---

## 🖥️ Backend Code Snapshots

### ☕ Java Backend
![Java Code](benchmark/output-images/java-code.jpg)

### 🟩 Node.js Backend
![Node Code](benchmark/output-images/node-code.jpg)

### 🐹 Go Backend
![Go Code](benchmark/output-images/go-code.jpg)

---

## 📖 Read Performance Benchmark

![Read Performance](benchmark/output-images/read-performance.jpg)

### Individual Results
- ☕ Java Read: ![Java Read](benchmark/output-images/java-read.jpg)
- 🟩 Node Read: ![Node Read](benchmark/output-images/node-read.jpg)
- 🐹 Go Read: ![Go Read](benchmark/output-images/go-read.jpg)

---

## ✍️ Write Performance Benchmark

![Write Performance](benchmark/output-images/write-performance.jpg)

### Individual Results
- ☕ Java Write: ![Java Write](benchmark/output-images/java-write.jpg)
- 🟩 Node Write: ![Node Write](benchmark/output-images/node-write.jpg)
- 🐹 Go Write: ![Go Write](benchmark/output-images/go-write.jpg)

---

## 📦 Empty Database Benchmark (GET /events/recent)

![Empty DB Performance](benchmark/output-images/emptydb-performance.jpg)

### Individual Results
- ☕ Java Empty: ![Java Empty](benchmark/output-images/java-empty.jpg)
- 🟩 Node Empty: ![Node Empty](benchmark/output-images/node-empty.jpg)
- 🐹 Go Empty: ![Go Empty](benchmark/output-images/go-empty.jpg)

---

## 📊 Final Comparison Summary Graphs

- 📈 Read Comparison:  
  ![Read Comparison](benchmark/output-images/read-performance.jpg)

- 📈 Write Comparison:  
  ![Write Comparison](benchmark/output-images/write-performance.jpg)

- 📈 Empty DB Comparison:  
  ![Empty DB Comparison](benchmark/output-images/emptydb-performance.jpg)

---
