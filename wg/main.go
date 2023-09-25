package main

import (
	"fmt"
	"sync"
)

func calculateSum(numbers []int, startIndex, endIndex int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done() // Notify WaitGroup when the goroutine is done

	sum := 0
	for i := startIndex; i < endIndex; i++ {
		sum += numbers[i]
	}

	resultChan <- sum // Send the partial sum to the result channel
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 4 // Number of goroutines to use
	listSize := len(numbers)

	var wg sync.WaitGroup
	resultChan := make(chan int, numWorkers)

	// Divide the list into segments for each worker
	segmentSize := listSize / numWorkers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		startIndex := i * segmentSize
		endIndex := (i + 1) * segmentSize

		// For the last segment, include any remaining elements
		if i == numWorkers-1 {
			endIndex = listSize
		}

		go calculateSum(numbers, startIndex, endIndex, &wg, resultChan)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	close(resultChan) // Close the result channel since we're done sending values

	totalSum := 0
	for partialSum := range resultChan {
		totalSum += partialSum
	}

	fmt.Printf("Total Sum: %d\n", totalSum)
}
