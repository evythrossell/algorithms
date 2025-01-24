package main

import (
	"fmt"
	"math"
)

func minimumValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

const (
	MaximumArraySize       = 201
	MaximumNumberOfResizes = 201
)

var t [MaximumArraySize][MaximumNumberOfResizes]int

func calculateMinimumWastedSpace(currentIndex, k int, nums []int) int {
	n := len(nums)
	if currentIndex == n {
		return 0
	}
	if k < 0 {
		return math.MaxInt32
	}
	if t[currentIndex][k] != -1 {
		return t[currentIndex][k]
	}

	maximumRequiredSize := nums[currentIndex]
	totalElements := 0
	minimumWastedSpace := math.MaxInt32
	for endIndex := currentIndex; endIndex < n; endIndex++ {
		maximumRequiredSize = maximumValue(maximumRequiredSize, nums[endIndex])
		totalElements += nums[endIndex]
		minimumWastedSpace = minimumValue(minimumWastedSpace, maximumRequiredSize*(endIndex-currentIndex+1)-totalElements+calculateMinimumWastedSpace(endIndex+1, k-1, nums))
	}
	t[currentIndex][k] = minimumWastedSpace
	return minimumWastedSpace
}

func minimumSpaceWastedResizing(nums []int, k int) int {
	for currentIndex := 0; currentIndex < MaximumArraySize; currentIndex++ {
		for endIndex := 0; endIndex < MaximumNumberOfResizes; endIndex++ {
			t[currentIndex][endIndex] = -1
		}
	}
	return calculateMinimumWastedSpace(0, k, nums)
}

func main() {
	nums := []int{10, 20}
	k := 0
	result := minimumSpaceWastedResizing(nums, k)
	fmt.Println("Minimum total space wasted to the first example:", result)

	nums = []int{10, 20, 30}
	k = 1
	result = minimumSpaceWastedResizing(nums, k)
	fmt.Println("Minimum total space wasted to the second example:", result)

	nums = []int{10, 20, 15, 30, 20}
	k = 2
	result = minimumSpaceWastedResizing(nums, k)
	fmt.Println("Minimum total space wasted to the third example:", result)
}
