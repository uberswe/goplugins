# Performance of go plugins

I created this project to test the performance of go plugins.

To build the plugin run the following
```bash
go build -buildmode=plugin -o plugin.so plugin/main.go
```

Currently, there does not seem to be much difference in performance.

| Benchmark                    | Operations | Operation Time |
| ---------------------------- |------------| ---------------|
| BenchmarkRandInt-12          | 87147016   | 13.5 ns/op     |
| BenchmarkPluginRandInt-12    | 79713106   | 14.9 ns/op     |
| BenchmarkRandString-12       | 132092     | 8427 ns/op     |
| BenchmarkPluginRandString-12 | 129799     | 8296 ns/op     |

To run the benchmark on your own system run:
```bash
go test -bench=.
```

If you would like to run the tests to ensure that the setup works as expected then run the following:
```bash
go test
```
