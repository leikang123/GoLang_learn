package main

import "fmt"

type Book struct {
	title  string
	aythor string
	price  int
	id     int
}

func main() {
	book := &Book{
		"leikanguu",
		"leikang",
		218,
		32,
	}
	fmt.Println(book.title)
	fmt.Println(book.aythor)
	fmt.Println(book.price)
	fmt.Println(book.id)

}
