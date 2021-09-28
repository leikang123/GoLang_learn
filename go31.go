package main

import "fmt"

func addSum(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub

}
func main() {
	a := 2
	b := 3
	sum, sub := addSum(a, b)
	fmt.Println(a, "+", b, "=", sum)
	fmt.Println(a, "-", b, "=", sub)

}
