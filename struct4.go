package main

import "fmt"

type Book struct {
	title  string
	author string
	id     int
	num    int
}

func main() {
	// 初始化
	book := &Book{
		"wangwangdui",
		"leikang",
		02033400,
		345,
	}
	fmt.Println("title:", book.title)
	fmt.Println("author:", book.author)
	fmt.Println("id:", book.id)
	fmt.Println("num:", book.num)
}
