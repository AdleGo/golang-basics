package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	reverse(&arr)
	fmt.Println("arr:", arr)
}

func reverse(arrParam *[4]int) {
	for i, val := range *arrParam {
		(*arrParam)[len(arrParam)-1-i] = val
	}
}
