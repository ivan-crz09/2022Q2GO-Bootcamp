package main

import (
	"fmt"
	"sort"
)

func getMedian(number []float64) float64 {
	sort.Float64s(number)
	totalNumbers := len(number)
	if totalNumbers%2 == 0 {
		currentIndex := totalNumbers / 2
		return (number[currentIndex-1] + number[currentIndex]) / float64(2)
	}

	return number[totalNumbers/2]
}

func getMean(number []float64) float64 {
	totalNumbers := len(number)
	result := 0.0
	for _, v := range number {
		result += v
	}

	return result / float64(totalNumbers)
}

func getMaxValue(number []float64) float64 {
	totalNumbers := len(number)
	sort.Float64s(number)

	return number[totalNumbers-1]
}

func main() {
	numbers := []float64{3, 10, 2, 1, 7}
	fmt.Println("Mean: ", getMean(numbers))
	fmt.Println("Median: ", getMedian(numbers))
	fmt.Println("Max Value: ", getMaxValue(numbers))
}
