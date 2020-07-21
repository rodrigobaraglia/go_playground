package main

import "fmt"

var a = [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
var b = [][]int{{4, 5, 6}, {4, 5, 6}, {4, 5, 6}}

func sum(a [][]int, b [][]int) [][]int {
	for i := range a {
		for j := range a[i] {
			a[i][j] += b[i][j]
		}
	}
	return a
}

func transpose(m [][]int) [][]int {
	mt := fromMatrix(m)
	for i := range m {
		for j := range m[i] {
			mt[j][i] = m[i][j]
		}
	}
	return mt
}

func mult(a [][]int, b [][]int) [][]int {

	for i := range a {
		for j := range a[i] {
			a[i][j] *= b[j][i]
		}
	}
	return a
}

func copyMatrix(m1 [][]int) (m2 [][]int) {
	m2 = fromMatrix(m1)
	for i := range m1 {
		for j := range m1[i] {
			m2[i][j] = m1[i][j]
		}
	}
	return
}

func fromMatrix(m [][]int) [][]int {
	rows, columns := len(m), len(m[0])
	copy := matrix(rows, columns)
	return copy
}

func matrix(rows int, cols int) [][]int {
	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}
	return m
}

func main() {
	fmt.Println(matrix(2, 2))
	fmt.Println(sum(copyMatrix(a), b))
	fmt.Println(mult(copyMatrix(a), b))
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println(transpose(a))
	fmt.Println(copyMatrix(a))

}
