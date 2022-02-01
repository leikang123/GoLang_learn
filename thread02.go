package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printl("go")
		}
	}()
	// 循环
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("main")
	}
}
