package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("temp:", a)

	a[4] = 100
	fmt.Print("set a as ", a)
	fmt.Println("get a[4] is ", a[4])

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println(b)
}
