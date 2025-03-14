# Benchmark Results: Go vs Node vs Python

## Summary

We benchmarked three simple applications generating random numbers, one built with Go, another with Node.js, and the last with Python. Below are the key results from the Apache Benchmark tests.

### Go Application (Port 8080)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 2.837 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 3,524.97
- **Transfer Rate**: 716.01 KB/sec
- **Document Length**: 91 bytes

### Node Application (Port 3000)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 11.334 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 882.26
- **Transfer Rate**: 251.58 KB/sec
- **Document Length**: 85 bytes

### Python Application (Port 8085)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 24.260 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 412.20
- **Transfer Rate**: 80.51 KB/sec
- **Document Length**: 75 bytes

---

## Performance Comparison

| Metric                       | **Go**                         | **Node**                       | **Python**                     |
|------------------------------|--------------------------------|--------------------------------|--------------------------------|
| **Total Requests**            | 10,000                         | 10,000                         | 10,000                         |
| **Time Taken for Tests**      | 2.837 seconds                 | 11.334 seconds                | 24.260 seconds                 |
| **Failed Requests**           | 0 (0% failure rate)           | 0 (0% failure rate)           | 0 (0% failure rate)           |
| **Requests per Second**       | 3,524.97                      | 882.26                         | 412.20                         |
| **Transfer Rate**             | 716.01 KB/sec                 | 251.58 KB/sec                 | 80.51 KB/sec                  |
| **Document Length**           | 91 bytes                      | 85 bytes                       | 75 bytes                      |
| **Mean Time per Request**     | 28.369 ms                     | 113.345 ms                     | 242.603 ms                    |
| **Mean Time (Concurrent)**    | 0.284 ms                      | 1.133 ms                      | 2.426 ms                      |
| **Median Processing Time**    | 25 ms                          | 110 ms                         | 221 ms                        |
| **90th Percentile (Time)**    | 30 ms                          | 177 ms                         | 317 ms                        |

---

## Key Insights

- **Throughput**: Go delivers the highest throughput with 3,524.97 requests per second, followed by Node (882.26 requests per second) and Python (412.20 requests per second).
- **Response Time**: Go is the fastest with a mean time per request of 28.369 ms, while Python has the slowest response time at 242.603 ms.
- **Transfer Rate**: Go also leads in transfer rate at 716.01 KB/sec, with Node following at 251.58 KB/sec, and Python lagging behind at 80.51 KB/sec.
- **Failure Rate**: All three applications handled the load without any failures (0% failure rate).

In conclusion, **Go** outperforms both Node and Python in terms of both throughput and response time. **Node** has a slower performance than Go but is more efficient than Python in terms of both throughput and transfer rate. **Python** performs the slowest across all benchmarks.