package main

import "fmt"

type Book struct {
	title  string
	author string
	id     int
	price  int
}

func main() {
	book := new(Book)
	fmt.Println(book)
}
