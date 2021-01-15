package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

type Total struct {
	sync.Mutex
	value int
}

var total = Total{}

func Test_Mu(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
}
func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

var sum uint64

func Test_Mu2(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()
	fmt.Println(sum)

}
func worker2(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&sum, i)
	}
}
