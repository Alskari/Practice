package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	//var m sync.Mutex
	var counter int64

	x := 10

	var wg sync.WaitGroup
	wg.Add(x)

	for i := 0; i < x; i++ {
		go func() {
			//m.Lock()
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
			//m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("done")
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
