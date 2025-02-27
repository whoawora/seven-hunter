package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMaxPathSumFromLevels(levels [][]int) int {
	if len(levels) == 0 {
		return 0
	}

	n := len(levels)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, len(levels[i]))
	}

	copy(dp[n-1], levels[n-1])

	// คำนวณจากล่างขึ้นบน
	for level := n - 2; level >= 0; level-- {
		for i, v := range levels[level] {
			dp[level][i] = v + Max(dp[level+1][i], dp[level+1][i+1])
		}
	}

	return dp[0][0]
}

func main() {
	file, err := os.Open("hard.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var levels [][]int
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&levels)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("ผลรวมสูงสุด: %d\n", FindMaxPathSumFromLevels(levels))
}
