package spinlock

import (
	"sync"
	"testing"
)

func BenchmarkSpinLock(b *testing.B) {
	var lock SpinLock

	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func BenchmarkMutex(b *testing.B) {
	var lock sync.Mutex

	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func BenchmarkRWMutexW(b *testing.B) {
	var lock sync.RWMutex

	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}

func BenchmarkRWMutexR(b *testing.B) {
	var lock sync.RWMutex

	for i := 0; i < b.N; i++ {
		lock.RLock()
		lock.RUnlock()
	}
}
