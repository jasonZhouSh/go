package main

import "fmt"

func main() {
	i := 1

	for {
		i = i + 1
		if i == 5 {
			fmt.Println(i)
			break
		}

	}

	go func() {
		fmt.Println("start go")
		for i < 10 {
			fmt.Println(i)
			i = i + 1
		}
		fmt.Println("end go")
	}()

	for j := 7; j < 1000; j++ {
		fmt.Println(j)
	}

	//var a1 string
	//fmt.Scan(&a1)

}
