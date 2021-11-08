//结构体
package main

import "fmt"

type Book struct {
	title  string
	prize  int
	author string
}

func main() {
	var book Book
	fmt.Println(book)

}
