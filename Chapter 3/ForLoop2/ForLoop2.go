package main

import "fmt"

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
		print(i)
	}

}

func print(x int) {
	fmt.Printf("%d\n", x)
}
