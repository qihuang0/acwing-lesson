package main

import "fmt"

func main() {
	input := []int{2, 3, 4, 1}
	mergeSort(input, 0, 3)
	fmt.Println(input)
}

func quickSort2(q []int, l int, r int) {
	if l >= r {
		return
	}
	i, j, mid := l, r, (l+r)/2
	for i < j {
		for q[i] < q[mid] {
			i++
		}
		for q[j] > q[mid] {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}

	}
	quickSort2(q, l, mid)
	quickSort2(q, mid+1, r)
}
