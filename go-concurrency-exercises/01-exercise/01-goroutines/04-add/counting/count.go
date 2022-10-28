package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	// Utilize all cores on machine
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Divide the input into parts
	var wg sync.WaitGroup
	sizeOfPart := len(numbers) / runtime.NumCPU()

	for i:=0; i < runtime.NumCPU(); i++{
		l :=i * sizeOfPart
		r :=l + sizeOfPart
		part:= numbers[l:r]

		wg.Add(1)
		go func(nums []int){
			defer wg.Done()

			var tempSum int64
			for _, num:= range nums{
				tempSum += int64(num)
			}
			atomic.AddInt64(&sum, tempSum)
		} (part)
	}

	// Run computation for each part in seperate goroutine.

	// Add part sum to cummulative sum
	wg.Wait()
	return sum
}
