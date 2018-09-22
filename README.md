# Spinlock 

[![Go Report Card](https://goreportcard.com/badge/github.com/barryz/spinlock)](https://goreportcard.com/report/github.com/barryz/spinlock)
[![Build Status](https://travis-ci.org/barryz/spinlock.svg?branch=master)](https://travis-ci.org/barryz/spinlock)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/oklog/run/master/LICENSE)

Spinlock is a [spin-lock](https://en.wikipedia.org/wiki/Spinlock) go language implementation.


## Usage

```go
import "github.com/barryz/spinlock"

lock := spinlock.New()

lock.Lock()     // Try to acquire the lock.
lock.Unlock()   // Release the lock.
lock.String()   // Print the lock status.

// embedding
type Foo struct {
    lock *spinlock.SpinLock // always use pointer to avoid copy
}
```

## Benchmark:

- Go 1.11.0
- OS X 10.13.6

```shell
BenchmarkSpinLock       100000000               14.5 ns/op             0 B/op          0 allocs/op 
BenchmarkSpinLock-2     100000000               14.4 ns/op             0 B/op          0 allocs/op
BenchmarkSpinLock-4     100000000               14.2 ns/op             0 B/op          0 allocs/op
BenchmarkMutex          100000000               16.1 ns/op             0 B/op          0 allocs/op
BenchmarkMutex-2        100000000               16.3 ns/op             0 B/op          0 allocs/op
BenchmarkMutex-4        100000000               18.1 ns/op             0 B/op          0 allocs/op
BenchmarkRWMutexW       50000000                33.5 ns/op             0 B/op          0 allocs/op
BenchmarkRWMutexW-2     50000000                33.4 ns/op             0 B/op          0 allocs/op
BenchmarkRWMutexW-4     50000000                33.5 ns/op             0 B/op          0 allocs/op
BenchmarkRWMutexR       100000000               15.7 ns/op             0 B/op          0 allocs/op
BenchmarkRWMutexR-2     100000000               15.9 ns/op             0 B/op          0 allocs/op
BenchmarkRWMutexR-4     100000000               18.1 ns/op             0 B/op          0 allocs/op
```
