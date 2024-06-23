package main

import "fmt"

func main() {
	num := []int{}

	for i := range 11 {
		num = append(num, i)
	}

	for _, n := range num {
		if n%2 == 0 {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
	}
}
