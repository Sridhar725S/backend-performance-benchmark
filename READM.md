🚀 Backend Performance Benchmark (Go vs Node.js vs Java)

This project compares the performance of Go, Node.js, and Java (Spring Boot) APIs using k6 load testing with MongoDB as the database.

It evaluates both:

📖 Read performance (GET /users)
✍️ Write performance (POST /users)
📦 Empty database scenario (GET /event/recent)
🗄️ Large dataset scenario (~5000+ records)
🧪 Tech Stack
Layer	Technology
Backend 1	Go (Gin + MongoDB)
Backend 2	Node.js (Express + MongoDB)
Backend 3	Java (Spring Boot + MongoDB)
Database	MongoDB
Load Testing	k6
📁 Project Structure
benchmark/
 ├── go-service/
 ├── node-service/
 ├── java-service/
 ├── k6-scripts/
 │     ├── read-test.js
 │     ├── write-test.js
 │     ├── empty-test.js
 ├── results/
 │     ├── go-read.json
 │     ├── node-read.json
 │     ├── java-read.json
 │     ├── go-write.json
 │     ├── node-write.json
 │     ├── java-write.json
 │     ├── empty.json
⚙️ How to Run Backend
🐹 Go Service
go run main.go
🟩 Node.js Service
npm install
node app.js
☕ Java Spring Boot
mvn spring-boot:run
🧪 How to Run Load Tests (k6)
📖 Read Test
k6 run read-test.js --out json=go-read.json
k6 run read-test.js --out json=node-read.json
k6 run read-test.js --out json=java-read.json
✍️ Write Test
k6 run write-test.js --out json=go-write.json
k6 run write-test.js --out json=node-write.json
k6 run write-test.js --out json=java-write.json
📦 Empty DB Test
k6 run empty-test.js --out json=go-empty.json
k6 run empty-test.js --out json=node-empty.json
k6 run empty-test.js --out json=java-empty.json
📊 Summary Results
📖 READ PERFORMANCE (5000+ MongoDB records)
Stack	Avg (ms)	P95 (ms)	Req/sec	Result
🐹 Go	~1359 ms	~2676 ms	41.9	🥇 Best throughput
🟩 Node.js	~8141 ms	~12363 ms	9.0	❌ Slow under load
☕ Java	~2282 ms	~4730 ms	29.4	⚖️ Stable
✍️ WRITE PERFORMANCE
Stack	Avg (ms)	P95 (ms)	Req/sec	Result
🐹 Go	~25 ms	~84 ms	95.9	🥇 Fastest
🟩 Node.js	~59 ms	~488 ms	92.7	🥈 Balanced
☕ Java	~152 ms	~1925 ms	83.5	🥉 Slowest
📦 EMPTY DB PERFORMANCE
Stack	Avg (ms)	P95 (ms)	Req/sec	Notes
🐹 Go	~21 ms	~107 ms	96.9	🥇 Best
🟩 Node.js	~31 ms	~218 ms	95.7	Stable
☕ Java	~365 ms	~4512 ms	71.2	High tail latency
📉 Key Observations
🐹 Go performs best overall in latency + throughput
🟩 Node.js performs well but degrades under heavy read load
☕ Java is stable but slower in most cases
⚠️ Write tests show issues in status validation (check API response codes)
🚨 Important Notes
All services must return correct HTTP status codes (200 / 201)
Ensure MongoDB indexes are enabled for better read performance
Use connection pooling in all services
Fix failing checks in k6 (currently some write tests show 0% success)
📊 Conclusion

👉 Best for high performance APIs: Go (Gin)
👉 Best for balanced development speed: Node.js
👉 Best for enterprise + structure: Java (Spring Boot)