package main

import "fmt"

var lst = []int{8, 22, 7, 9, 31, 19, 5, 13}

func main() {
	count := 0
	for i := 0; i < len(lst)-1; i++ {
		for j := 1; j < len(lst); j++ {
			if lst[j] < lst[j-1] {
				lst[j], lst[j-1] = lst[j-1], lst[j]
				count++
			}
		}
	}

	fmt.Println(count)
	fmt.Println(lst)

}
