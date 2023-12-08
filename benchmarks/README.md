# Benchmarks

This directory contains a simple round-trip marshalling benchmark to compare against the string variant.

# Current results

```
goos: linux
goarch: amd64
pkg: github.com/exaring/ulid-protobuf/benchmarks
cpu: 12th Gen Intel(R) Core(TM) i7-1250U
BenchmarkString
BenchmarkString-12      55256742               202.2 ns/op            96 B/op          3 allocs/op
BenchmarkNative
BenchmarkNative-12      46060738               255.0 ns/op           152 B/op          3 allocs/op
```