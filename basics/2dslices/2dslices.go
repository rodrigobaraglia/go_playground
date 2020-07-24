package main

import "fmt"

var a = [][]float64{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
var b = [][]float64{{4, 5, 6}, {4, 5, 6}, {4, 5, 6}}

var c = [][]float64{{34, 45, 7}, {4, 23, 53}, {13, 14, 12}}
var d = [][]float64{{1, 2}, {3, 4}, {5, 6}}

//Matrix sum
func sum(a [][]float64, b [][]float64) [][]float64 {
	for i := range a {
		for j := range a[i] {
			a[i][j] += b[i][j]
		}
	}
	return a
}

//Matrix transposition
func transpose(m [][]float64) [][]float64 {
	mt := fromMatrix(m)
	for i := range m {
		for j := range m[i] {
			mt[j][i] = m[i][j]
		}
	}
	return mt
}

//Matrix multiplication
func mult(a [][]float64, b [][]float64) [][]float64 {

	for i := range a {
		for j := range a[i] {
			a[i][j] *= b[j][i]
		}
	}
	return a
}

//Self-Explanatory
func copyMatrix(m1 [][]float64) (m2 [][]float64) {
	m2 = fromMatrix(m1)
	for i := range m1 {
		for j := range m1[i] {
			m2[i][j] = m1[i][j]
		}
	}
	return
}

//Generates a slice of empty slices with the same rows and columns as m
func fromMatrix(m [][]float64) [][]float64 {
	rows, columns := len(m), len(m[0])
	copy := matrix(rows, columns)
	return copy
}

//Generates an empty slice of slices with a given number of rows and columns
func matrix(rows int, cols int) [][]float64 {
	m := make([][]float64, rows)
	for i := range m {
		m[i] = make([]float64, cols)
	}
	return m
}

//Zips the columns of a given matrix into the rows of a new matrix
func zipMatrix(m [][]float64) [][]float64 {
	rows, columns := len(m[0]), len(m)
	z := matrix(rows, columns)
	for i := range z {
		for j := range m {
			z[i][j] = m[j][i]
		}
	}
	return z
}

//Brute force non-DRY Gauss Jordan elimination method with O(SHIT) complexity.
func reduce(m [][]float64, img [][]float64) ([][]float64, [][]float64) {
	for i := range m {
		if m[i][i] != 0 {
			pivot := m[i][i]
			for j := range m {
				if j != i && m[j][i] != 0 {
					base := m[j][i]

					go func() {
						for k := range m[j] {

							m[j][k] = (m[j][k] * pivot) - (m[i][k] * base)
						}
					}()
					go func() {
						for k := range img[j] {
							img[j][k] = (img[j][k] * pivot) - (img[i][k] * base)
						}
					}()

				}
			}
		} else {
			for k := range img[i] {
				img[i][k] = 0
			}
		}

	}
	for i := range m {
		if m[i][i] != 0 {
			divisor := m[i][i]
			fmt.Println("Divisor: ", divisor)
			go func() {
				for j := range img[i] {
					img[i][j] /= divisor
				}
			}()
			go func() {
				for k := range m[i] {
					m[i][k] /= divisor
				}
			}()

		}
	}
	return m, zipMatrix(img)
}

func main() {
	fmt.Println(sum(copyMatrix(a), b))
	fmt.Println(mult(copyMatrix(a), b))
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println(transpose(a))
	fmt.Println(copyMatrix(a))
	fmt.Println(reduce(c, d))

}
