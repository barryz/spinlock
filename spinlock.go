package spinlock

import (
	"runtime"
	"sync"
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

// Lock will try to acquire the spin-lock, if the lock is already in use, the caller blocks until Unlock is called
func (l *SpinLock) Lock() {
	for !l.TryLock() {
		runtime.Gosched() // interrupt current goroutine execution and allows other goroutines to do some stuff.
	}
}

// Unlock will release the spin-lock.
// Calling a Unlock on a spin-lock there is no harmful.
func (l *SpinLock) Unlock() {
	atomic.StoreUint32(&l.s, unlocked)
}

// String returns the lock status string literal.
func (l *SpinLock) String() string {
	if atomic.LoadUint32(&l.s) == 1 {
		return "Locked"
	}

	return "Unlocked"
}

// NewSpinLock creates an new spin-lock.
func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}
