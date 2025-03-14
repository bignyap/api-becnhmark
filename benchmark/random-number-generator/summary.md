Here is the benchmark comparison in the format you requested, with the results from Go, Node, and Python:

# Benchmark Results: Go vs Node vs Python

## Summary

We benchmarked three simple applications generating random numbers, one built with Go, another with Node.js, and the last with Python. Below are the key results from the Apache Benchmark tests.

### Go Application (Port 8080)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 2.238 seconds
- **Failed Requests**: 5,398 (54% failure rate)
- **Requests per Second**: 4,468.29
- **Transfer Rate**: 569.21 KB/sec
- **Document Length**: 13 bytes

### Node Application (Port 3000)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 2.864 seconds
- **Failed Requests**: 3,900 (39% failure rate)
- **Requests per Second**: 3,491.49
- **Transfer Rate**: 805.58 KB/sec
- **Document Length**: 29 bytes

### Python Application (Port 8085)
- **Total Requests**: 10,000
- **Time Taken for Tests**: 7.876 seconds
- **Failed Requests**: 1,090 (10% failure rate)
- **Requests per Second**: 1,269.64
- **Transfer Rate**: 170.00 KB/sec
- **Document Length**: 12 bytes

---

## Performance Comparison

| Metric                       | **Go**                         | **Node**                       | **Python**                     |
|------------------------------|--------------------------------|--------------------------------|--------------------------------|
| **Total Requests**            | 10,000                         | 10,000                         | 10,000                         |
| **Time Taken for Tests**      | 2.238 seconds                 | 2.864 seconds                 | 7.876 seconds                 |
| **Failed Requests**           | 5,398 (54% failure rate)      | 3,900 (39% failure rate)      | 1,090 (10% failure rate)      |
| **Requests per Second**       | 4,468.29                      | 3,491.49                      | 1,269.64                      |
| **Transfer Rate**             | 569.21 KB/sec                 | 805.58 KB/sec                 | 170.00 KB/sec                 |
| **Document Length**           | 13 bytes                      | 29 bytes                      | 12 bytes                      |
| **Mean Time per Request**     | 22.380 ms                     | 28.641 ms                     | 78.762 ms                     |
| **Mean Time (Concurrent)**    | 0.224 ms                      | 0.286 ms                      | 0.788 ms                      |
| **Median Processing Time**    | 22 ms                          | 28 ms                          | 56 ms                          |
| **90th Percentile (Time)**    | 27 ms                          | 36 ms                          | 70 ms                          |

---

## Key Insights

- **Throughput**: Go outperforms both Node and Python in terms of requests per second (4,468.29 vs. 3,491.49 vs. 1,269.64).
- **Response Time**: Go delivers the fastest mean time per request and the quickest processing times across percentiles.
- **Failure Rate**: Go has the highest failure rate (~54%), compared to Node (~39%) and Python (~10%).
- **Transfer Rate**: Node transfers data faster due to the larger document size (805.58 KB/sec vs. 569.21 KB/sec for Go and 170.00 KB/sec for Python).
- **Python Performance**: Python is significantly slower in terms of throughput and response times compared to both Go and Node.

In summary, **Go** provides the best performance for throughput and response time, while **Node** is faster in data transfer and has a lower failure rate. **Python** lags in all aspects, with the slowest throughput, highest time per request, and lowest transfer rate.