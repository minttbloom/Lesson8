package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var numbers = [10_000_000]int{}

func generateNumbers() {
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}
}

func main() {
	generateNumbers()
	runInMainThread()
	runInGoThreads()
}

func runInMainThread() {
	start := time.Now()

	for _, v := range numbers {
		Sin3x(v)
	}

	fmt.Printf("[MAIN] Time: %s\n", time.Since(start))
}

func runInGoThreads() {
	var wg sync.WaitGroup

	start := time.Now()

	maxThreads := 5
	lastPosition := 0
	for i := 1; i <= maxThreads; i++ {
		to := (len(numbers) / maxThreads) * i
		calculateFor(numbers[lastPosition:to], &wg)
		lastPosition = to
	}
	wg.Wait()

	fmt.Printf("[GO] Time: %s\n", time.Since(start))
}

func calculateFor(slice []int, wg *sync.WaitGroup) {
	go func(slice []int) {
		defer wg.Done()

		wg.Add(1)

		for _, v := range slice {
			Sin3x(v)
		}
	}(slice)
}

func Sin3x(x int) {
	_ = math.Sin(math.Sin(math.Sin(math.Sin(math.Sin(math.Sin(math.Sin(math.Sin(float64(x)))))))))
}
