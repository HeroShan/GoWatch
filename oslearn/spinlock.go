package main

import (
 "runtime"
 "sync"
 "sync/atomic"
 "fmt"
)

type SpinLock struct {
 state int64
}

func (s *SpinLock) Lock() {
	for {
			if atomic.CompareAndSwapInt64(&s.state, 0, 1) {
			return
			}
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	if !atomic.CompareAndSwapInt64(&s.state, 1, 0) {
	panic("unlock of unlocked spinlock")
	}
}
// go run -race
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var spin SpinLock
	x := 0
	inc := func() {
	defer wg.Done()
		spin.Lock()
		defer spin.Unlock()
			for i := 0; i < 10; i++ {
			x++
			fmt.Println(x)
			}
	}
	for n := 0; n < 4; n++ {
		go inc()
		fmt.Println("inc======",n)
	}

 	wg.Wait()
}
