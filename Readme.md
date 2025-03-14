# Benchmark Comparison Report: Go vs Node.js vs Python

## Overview
This report compares the performance of three simple applications—one built with Go, another with Node.js, and the third with Python. The applications were benchmarked using Apache Benchmark for three different scenarios: a "Hello World" endpoint, a random number generator, and a database connection simulation.

---

## Benchmark 1: **Hello World**

### Key Results:

#### Go Application (Port 8080)
- **Total Requests**: 10,000
- **Time Taken**: 2.531 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 3,951.42
- **Transfer Rate**: 474.63 KB/sec
- **Document Length**: 7 bytes

#### Node Application (Port 3000)
- **Total Requests**: 10,000
- **Time Taken**: 3.313 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 3,017.97
- **Transfer Rate**: 604.18 KB/sec
- **Document Length**: 7 bytes

#### Python Application (Port 8085)
- **Total Requests**: 10,000
- **Time Taken**: 17.737 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 563.78
- **Transfer Rate**: 75.98 KB/sec
- **Document Length**: 13 bytes

### Performance Comparison:

| Metric                       | **Go**          | **Node**         | **Python**        |
|------------------------------|-----------------|------------------|-------------------|
| **Total Requests**            | 10,000          | 10,000           | 10,000            |
| **Time Taken for Tests**      | 2.531 seconds   | 3.313 seconds    | 17.737 seconds    |
| **Requests per Second**       | 3,951.42        | 3,017.97         | 563.78            |
| **Transfer Rate**             | 474.63 KB/sec   | 604.18 KB/sec    | 75.98 KB/sec      |
| **Document Length**           | 7 bytes         | 7 bytes          | 13 bytes          |
| **Mean Time per Request**     | 25.307 ms       | 33.135 ms        | 177.375 ms        |
| **90th Percentile (Time)**    | 30 ms           | 47 ms            | 243 ms            |

### Key Insights:
- **Go** has the highest throughput and the lowest response time.
- **Node** transfers data faster than Go, thanks to a larger document size.
- **Python** lags in both throughput and response times, making it the slowest in this test.

---

## Benchmark 2: **Random Number Generator**

### Key Results:

#### Go Application (Port 8080)
- **Total Requests**: 10,000
- **Time Taken**: 2.238 seconds
- **Failed Requests**: 5,398 (54% failure rate)
- **Requests per Second**: 4,468.29
- **Transfer Rate**: 569.21 KB/sec
- **Document Length**: 13 bytes

#### Node Application (Port 3000)
- **Total Requests**: 10,000
- **Time Taken**: 2.864 seconds
- **Failed Requests**: 3,900 (39% failure rate)
- **Requests per Second**: 3,491.49
- **Transfer Rate**: 805.58 KB/sec
- **Document Length**: 29 bytes

#### Python Application (Port 8085)
- **Total Requests**: 10,000
- **Time Taken**: 7.876 seconds
- **Failed Requests**: 1,090 (10% failure rate)
- **Requests per Second**: 1,269.64
- **Transfer Rate**: 170.00 KB/sec
- **Document Length**: 12 bytes

### Performance Comparison:

| Metric                       | **Go**          | **Node**         | **Python**        |
|------------------------------|-----------------|------------------|-------------------|
| **Total Requests**            | 10,000          | 10,000           | 10,000            |
| **Time Taken for Tests**      | 2.238 seconds   | 2.864 seconds    | 7.876 seconds     |
| **Requests per Second**       | 4,468.29        | 3,491.49         | 1,269.64          |
| **Transfer Rate**             | 569.21 KB/sec   | 805.58 KB/sec    | 170.00 KB/sec     |
| **Document Length**           | 13 bytes        | 29 bytes         | 12 bytes          |
| **Mean Time per Request**     | 22.380 ms       | 28.641 ms        | 78.762 ms         |
| **90th Percentile (Time)**    | 27 ms           | 36 ms            | 70 ms             |

### Key Insights:
- **Go** has the highest throughput and the fastest response time.
- **Node** is faster in terms of data transfer but has a higher failure rate than Go.
- **Python** is significantly slower than both Go and Node, particularly in terms of throughput.

---

## Benchmark 3: **DB Connection**

### Key Results:

#### Go Application (Port 8080)
- **Total Requests**: 10,000
- **Time Taken**: 2.837 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 3,524.97
- **Transfer Rate**: 716.01 KB/sec
- **Document Length**: 91 bytes

#### Node Application (Port 3000)
- **Total Requests**: 10,000
- **Time Taken**: 11.334 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 882.26
- **Transfer Rate**: 251.58 KB/sec
- **Document Length**: 85 bytes

#### Python Application (Port 8085)
- **Total Requests**: 10,000
- **Time Taken**: 24.260 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 412.20
- **Transfer Rate**: 80.51 KB/sec
- **Document Length**: 75 bytes

### Performance Comparison:

| Metric                       | **Go**          | **Node**         | **Python**        |
|------------------------------|-----------------|------------------|-------------------|
| **Total Requests**            | 10,000          | 10,000           | 10,000            |
| **Time Taken for Tests**      | 2.837 seconds   | 11.334 seconds   | 24.260 seconds    |
| **Requests per Second**       | 3,524.97        | 882.26           | 412.20            |
| **Transfer Rate**             | 716.01 KB/sec   | 251.58 KB/sec    | 80.51 KB/sec      |
| **Document Length**           | 91 bytes        | 85 bytes         | 75 bytes          |
| **Mean Time per Request**     | 28.369 ms       | 113.345 ms       | 242.603 ms        |
| **90th Percentile (Time)**    | 30 ms           | 177 ms           | 317 ms            |

### Key Insights:
- **Go** performs best in throughput and response time.
- **Node** is slower than Go but more efficient than Python.
- **Python** is the slowest across all metrics.

---

## Overall Summary

### **Throughput**:
- **Go** consistently outperforms both Node and Python across all benchmarks, delivering the highest requests per second.

### **Response Time**:
- **Go** offers the fastest response time, significantly outperforming Node and Python.

### **Transfer Rate**:
- **Node** excels in transfer rate, particularly when handling larger document sizes, though Go is still competitive in this area.

### **Failure Rate**:
- **Go** had the highest failure rate in the random number generator test (~54%), followed by Node (~39%) and Python (~10%). Despite the higher failure rate, Go maintained superior throughput and response times.

### **Python Performance**:
- **Python** lags behind in all metrics, with the lowest throughput, highest response times, and lowest transfer rates.

---

## Conclusion
- **Go** offers the best overall performance across all scenarios, providing high throughput, low response times, and solid transfer rates.
- **Node** performs well but generally trails Go in throughput and response times while excelling in transfer rate.
- **Python** is the slowest across all benchmarks, making it less ideal for performance-sensitive applications.

This benchmarking clearly highlights **Go** as the top performer for high-throughput applications. **Node** is a good alternative if data transfer speed is a higher priority, while **Python** is best suited for less performance-critical tasks.
