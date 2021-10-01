package main

import "fmt"
/**语法格式：
func 函数名(参数列表) (返回值类表){}
	函数体
参数列表中的每个参数都由参数名称和参数类型两部分组成，参数变量为函数
的局部变量。如果函数的参数数量不固定，Go语言函数还支持可变参数。
返回参数列表:返回参数列表中的每个参数由返回的参数名称和参数类型组成，也可简写
为返回值类型列表。
*/
func addSum(x int, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub

}
func main() {
	a := 2
	b := 3
	sum, sub := addSum(a, b)
	fmt.Println(a, "+", b, "=", sum)
	fmt.Println(a, "-", b, "=", sub)

}
