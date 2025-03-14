# Benchmark Results: Go vs Node vs Python

## Summary

We benchmarked three simple applications generating random numbers, one built with Go, another with Node.js, and the last with Python. Below are the key results from the Apache Benchmark tests.

### Go Application (Port 8080)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 2.531 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 3,951.42
- **Transfer Rate**: 474.63 KB/sec
- **Document Length**: 7 bytes

### Node Application (Port 3000)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 3.313 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 3,017.97
- **Transfer Rate**: 604.18 KB/sec
- **Document Length**: 7 bytes

### Python Application (Port 8085)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 17.737 seconds
- **Failed Requests**: 0 (0% failure rate)
- **Requests per Second**: 563.78
- **Transfer Rate**: 75.98 KB/sec
- **Document Length**: 13 bytes

---

## Performance Comparison

| Metric                       | **Go**                         | **Node**                       | **Python**                     |
|------------------------------|--------------------------------|--------------------------------|--------------------------------|
| **Total Requests**            | 10,000                         | 10,000                         | 10,000                         |
| **Time Taken for Tests**      | 2.531 seconds                 | 3.313 seconds                 | 17.737 seconds                 |
| **Failed Requests**           | 0 (0% failure rate)           | 0 (0% failure rate)           | 0 (0% failure rate)           |
| **Requests per Second**       | 3,951.42                      | 3,017.97                      | 563.78                        |
| **Transfer Rate**             | 474.63 KB/sec                 | 604.18 KB/sec                 | 75.98 KB/sec                  |
| **Document Length**           | 7 bytes                       | 7 bytes                       | 13 bytes                      |
| **Mean Time per Request**     | 25.307 ms                     | 33.135 ms                     | 177.375 ms                    |
| **Mean Time (Concurrent)**    | 0.253 ms                      | 0.331 ms                      | 1.774 ms                      |
| **Median Processing Time**    | 24 ms                          | 31 ms                          | 150 ms                        |
| **90th Percentile (Time)**    | 30 ms                          | 47 ms                          | 243 ms                        |

---

## Key Insights

- **Throughput**: Go performs the best with 3,951.42 requests per second, followed by Node (3,017.97) and Python (563.78).
- **Response Time**: Go is the fastest with a mean time per request of 25.307 ms, while Python has the slowest response time with 177.375 ms.
- **Transfer Rate**: Node transfers data the fastest at 604.18 KB/sec, compared to Go (474.63 KB/sec) and Python (75.98 KB/sec).
- **Failure Rate**: All three applications had zero failures during the test.

In summary, **Go** delivers the highest performance in terms of throughput and response time. **Node** performs well in terms of transfer rate, while **Python** significantly lags behind in both throughput and response time.
