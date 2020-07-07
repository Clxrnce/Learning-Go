package main

import "fmt"

func main() {
	array := []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("unsorted %v\n", array)
	bubblesort(array)
	fmt.Printf("sorted %v\n", array)
}

func bubblesort(a []int) {
	for x := 0; x < len(a)-1; x++ {
		for y := x + 1; y < len(a); y++ {
			if a[y] < a[x] {
				a[x], a[y] = a[y], a[x]
			}
		}
	}
}
