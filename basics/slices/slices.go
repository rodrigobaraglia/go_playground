package main

import "fmt"

func printWords(wrds []string) {
	for i := 0; i < len(wrds); i++ {
		fmt.Println(wrds[i])
	}
	return
}
func printWords2(wrds []string) {
	for i, value := range wrds {
		fmt.Println(i, value)
	}
	return
}
func printWords3(wrds []string) {
	for _, value := range wrds {
		fmt.Println(value)
	}
	return
}

func thrice(nums []int) []int {

	for i, value := range nums {
		nums[i] = value * 3
	}
	return nums
}

var names = []string{"Pepe", "Juan", "Carlos"}
var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	fmt.Println(names)
	printWords(names)
	printWords2(names)
	printWords3(names)
	fmt.Println(thrice(nums))
}
