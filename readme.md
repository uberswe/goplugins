# Performance of go plugins

I created this project to test the performance of go plugins.

To build the plugin run the folowing
```bash
go build -buildmode=plugin -o plugin.so plugin/main.go
```

Currently, there does not seem to be much difference in performance.

BenchmarkRandInt-12             87147016                13.5 ns/op
BenchmarkPluginRandInt-12       79713106                14.9 ns/op