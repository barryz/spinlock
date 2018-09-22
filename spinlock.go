package spinlock

import (
	"runtime"
	"sync/atomic"
)

const (
	unlocked uint32 = iota
	locked
)

// SpinLock implements a atomic spin lock, the zero value for a SpinLock is an unlocked spin-lock.
// After the first use, we SHOULD NOT copy the value.
type SpinLock struct {
	s uint32
}

// TryLock will try to acquire the lock and return it succeed or not.
func (l *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapUint32(&l.s, unlocked, locked)
}

// TryUnlock will try to release the lock and return it succeed or not.
func (l *SpinLock) TryUnlock() bool {
	return atomic.CompareAndSwapUint32(&l.s, locked, unlocked)

}

// Lock acquire the spin-lock, if the lock is already in use, the caller blocks until Unlock is called
func (l *SpinLock) Lock() {
	for !l.TryLock() {
		runtime.Gosched() // interrupt current goroutine execution and allows other goroutines to do some stuff.
	}

}

func (l *SpinLock) Unlock() {
	for !l.TryUnlock() {
		runtime.Gosched()
	}
}

// String returns the lock status string literal.
func (l *SpinLock) String() string {
	if atomic.LoadUint32(&l.s) == 1 {
		return "Locked"
	}

	return "Unlocked"
}
