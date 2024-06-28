package main

import "fmt"

func isHappy(n int) bool {
	seen := map[int]bool{}
	for n != 1 {
		sum := calculateSquareSum(n)
		if seen[sum] {
			return false
		}
		seen[sum] = true
		n = sum
	}
	return true
}

func calculateSquareSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

func main() {
	var number int

	fmt.Println("enter a positive integer:")
	fmt.Scanf("%d", &number)

	if isHappy(number) {
		fmt.Println(number, "is a happy number.")
	} else {
		fmt.Println(number, "is not a happy number.")
	}
}