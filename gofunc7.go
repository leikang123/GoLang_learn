package main

import "fmt"


// 定义函数名两个变量，然后两个变量的参数返回加法和减法
func  addSub( x int,y int ) (sum int,sub int)  {
	sum = x+ y
	sub = x -y

	return sum,sub
	
}
func main(){
	a :=2
	b :=3
	var f1 func (x int , y int) (sun int,sub int)

	f1 = addSub

	sum,sub :=f1(a,b)
	fmt.Println(sum, "=",a,"+",b )
	fmt.Println(sub,"=",a,"-",b)

}