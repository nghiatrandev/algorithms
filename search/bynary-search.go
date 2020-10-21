package main

import "fmt"

func main() {
	findN := 36
	l := initLst()
	isExist := false
	for isExist == false {
		f := l[0 : len(l)/2]
		s := l[len(l)/2 : len(l)]

		if l[len(l)/2] == findN {
			isExist = true
			fmt.Println(l[len(l)/2])
			break
		} else if l[len(l)/2] > findN {
			l = f
		} else if l[len(l)/2] < findN {
			l = s
		}
	}

	fmt.Printf("isExist: %v", isExist)

}

func initLst() []int {

	l := make([]int, 0)

	for i := 1; i < 101; i++ {
		l = append(l, i)
	}
	fmt.Printf("%v ", l)
	return l

}
