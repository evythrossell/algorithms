package main

import (
	"log"
	"math"
)

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func maximumContainerArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		currentHeight := min(height[left], height[right])
		currentArea := width * currentHeight

		log.Printf("Current Area: %d, Left: %d, Right: %d\n", currentArea, left, right)

		maxArea = max(maxArea, currentArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea := maximumContainerArea(height)

	log.Printf("Maximum Area: %d\n", maxArea)
}
