package main

import "fmt"
// 结构体的继承
type Book struct {
	title  string
	author string
	id     int
	num    int
}
type Book1 struct {
	Book
	price int
}
type Book2 struct {
	Book
	bookSize int
}

func main() {
	b := &Book{}
	b1 := &Book1{}
	b2 := &Book2{}
	fmt.Println("b:", b)
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)

}
