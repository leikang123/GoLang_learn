package main

import "fmt"

func addSub(x int, y int )  (sum int , sub int){

	// 两个变量的计算
	sum = x + y
	sub = x -y
	 return sum, sub 

}

func main (){
	a :=4
	b :=9
	f2 := addSub
	sum,sub := f2(a,b)
	fmt.Println(a,"+",b,"=",sum)
	fmt.Println(a,"-",b,"=",sub)
}