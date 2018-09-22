package spinlock

import (
	"testing"
	"time"
)

func TestSpinLock(t *testing.T) {
	lock := new(SpinLock)

	testMap := make(map[int]int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 1000; i++ {
			lock.Lock()
			testMap[i] = i + 1             // writer
			<-time.After(time.Millisecond) // spin-lock used which blocked for only short periods.
			lock.Unlock()
		}

		done <- struct{}{}
	}()

	for i := 0; i < 1000; i++ {
		lock.Lock()
		_ = testMap[i] // reader
		<-time.After(time.Millisecond)
		lock.Unlock()
	}

	<-done
}
