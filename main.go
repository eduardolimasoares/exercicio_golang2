package main

import "fmt"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {

	a := makeRange(1, 100)

	for _, num := range a {
		if (num % 2) != 0 {
			fmt.Println(num)
		}
	}

}
