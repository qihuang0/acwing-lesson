package main

import "fmt"

func main() {
	input := []int{2, 3, 4, 1}
	mergeSort(input, 0, 3)
	fmt.Println(input)
}
