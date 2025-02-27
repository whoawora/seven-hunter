package main

import "fmt"

func findMinSum(pattern string) []int {
	n := len(pattern) + 1
	digits := make([]int, n)
	minSum, result := -1, []int{}

	for {
		valid := true
		for i := 0; i < n-1; i++ {
			if (pattern[i] == 'L' && digits[i] <= digits[i+1]) ||
				(pattern[i] == 'R' && digits[i] >= digits[i+1]) ||
				(pattern[i] == '=' && digits[i] != digits[i+1]) {
				valid = false
				break
			}
		}
		if valid {
			sum := 0
			for _, d := range digits {
				sum += d
			}
			if minSum == -1 || sum < minSum {
				minSum = sum
				result = append([]int{}, digits...)
			}
		}

		pos := n - 1
		for pos >= 0 && digits[pos] == 9 {
			digits[pos] = 0
			pos--
		}
		if pos < 0 {
			break
		}
		digits[pos]++
	}
	return result
}

func main() {
	for _, p := range []string{"LLRR=", "==RLL", "=LLRR", "RRL=R"} {
		r := findMinSum(p)
		fmt.Printf("Input: %s Output: ", p)
		for _, d := range r {
			fmt.Print(d)
		}
		fmt.Println()
	}
}
