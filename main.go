package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	n := randomArray(10000, 10000)

	fmt.Println("Array to be sorted: ", n)

	sleepSort(n, &wg)

	wg.Wait()
}

func sleepSort(n []int, wg *sync.WaitGroup) {

	for _, x := range n {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			time.Sleep(time.Duration(x) * time.Millisecond)
			fmt.Println(x)
		}(x)
	}

}

func randomArray(n, x int) []int {
	var nums []int
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < n; i++ {
		r := random.Intn(x)
		nums = append(nums, r)
	}

	return nums
}
