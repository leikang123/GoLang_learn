package main

// 线程的并发方法

import (
	"fmt"
	"time"
)

func task1() {
	for {
		fmt.Println(time.Now().Format("15:09:09"), "正在处理task1的任务！")
		time.Sleep(time.Second * 3)
	}
}
func task2() {
	for {
		fmt.Println(time.Now().Format("15:11:13"), "正在处理task2的任务")
		time.Sleep(time.Second * 1)
	}
}
func main() {
	go task1()
	go task2()
	for {
		fmt.Println(time.Now().Format("17:01:01"), "正在处理住线程的任务")
		time.Sleep(time.Second * 2)
	}
}
