package main

import "fmt"

// Pascal returns the first n rows of Pascal's triangle.
// Each interior entry is the sum of the two entries above it.
func Pascal(n int) [][]int {
	rows := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows = append(rows, row)
	}
	return rows
}

func main() {
	for _, row := range Pascal(6) {
		fmt.Println(row)
	}
}
