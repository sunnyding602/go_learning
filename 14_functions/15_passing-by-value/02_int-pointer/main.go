package main

import "fmt"

func main() {
	n := 44
	changeMe(&n)
	fmt.Println(&n)
	fmt.Println(n)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 88
	fmt.Println(z)
	fmt.Println(*z)
}
