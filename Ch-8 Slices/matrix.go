package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	result := [][]int{}
	for i := 0; i < rows; i++ {
		k := i
		for j := 0; j < cols; j++ {
			result[i] = append(result[i], k+i)
		}
		k++
	}
	return result
}

