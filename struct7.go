package main

import "fmt"

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
type Book3 struct {
	Book
	Book1
	Book2
	page int
}

func main() {
	book := &Book{
		"love is it",
		"peterds",
		98231,
		987,
	}
	fmt.Println("title:", book.title, "author:", book.author, "id:", book.id, "num:", book.num)
	b2 := &Book1{
		Book{"love is it",
			"peterds",
			98231,
			987},
		87,
	}
	fmt.Println("title:", b2.title, "author:", b2.author, "id:", b2.id, "num:", b2.num, "price:", b2.price)
	b3 := &Book2{
		Book{
			"love is it",
			"peterds",
			98231,
			987},
		12,
	}
	fmt.Println("title:", b3.title, "author:", b3.author, "id:", b3.id, "num:", b3.num, "bookSize:", b3.bookSize)
	b4 := &Book3{
		Book{
			"love is it",
			"peterds",
			98231,
			987},
		Book1{},
		Book2{},
		140,
	}
	fmt.Println("title:", b4.title, "author:", b4.author, "id:", b4.id, "num:", b4.num, "bookSize:", b4.bookSize)
}
