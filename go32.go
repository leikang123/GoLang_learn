package main

import "fmt"

func devaiv(a int, b int) (dev int, aiv int) {
	dev = a * b
	aiv = a / b
	return dev, aiv

}
func main() {
	var a int = 10
	var b int = 20
	var f func(a int, b int) (dev int, aiv int)
	f = devaiv
	dev, aiv := f(a, b)
	//dev, aiv := devaiv(a, b)
	fmt.Println(a, "*", b, "=", dev)
	fmt.Println(a, "/", b, "=", aiv)
}
