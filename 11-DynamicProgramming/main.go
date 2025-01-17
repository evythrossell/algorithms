package main

import "fmt"

type stone struct {
	position int
}

type key struct {
	index int
	jump  int
}

func canCrossRecursive(stones []stone, index int, jump int, memory map[key]bool) bool {
	if index == len(stones)-1 {
		return true
	}

	k := key{index, jump}
	if result, ok := memory[k]; ok {
		return result
	}

	result := false
	for j := index + 1; j < len(stones); j++ {
		newJump := stones[j].position - stones[index].position
		if newJump >= jump-1 && newJump <= jump+1 && canCrossRecursive(stones, j, newJump, memory) {
			result = true
			break
		}
	}

	memory[k] = result
	return result
}

func canCross(stones []stone) bool {
	if stones[1].position != 1 {
		return false
	}

	memory := make(map[key]bool)

	return canCrossRecursive(stones, 1, 1, memory)
}

func main() {
	stoneSequency1 := []stone{{0}, {1}, {3}, {5}, {6}, {8}, {12}, {17}}
	fmt.Println(canCross(stoneSequency1))

	stoneSequency2 := []stone{{0}, {1}, {2}, {3}, {4}, {8}, {9}, {11}}
	fmt.Println(canCross(stoneSequency2))
}
