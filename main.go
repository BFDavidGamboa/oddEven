package main

import "fmt"

func main() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range num {
		if val%2 != 0 {
			fmt.Println(val, "is odd")
		} else if val%2 != 1 {
			fmt.Println(val, "is even")
		}

	}
}
