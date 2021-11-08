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
		title :"leikang",
		aythor :"leikang",
		price:218,
		id:32,
	}
	fmt.Println(book.title)
	fmt.Println(book.aythor)
	fmt.Println(book.price)
	fmt.Println(book.id)


