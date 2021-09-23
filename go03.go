//指针类型学习
package main
import "fmt"
func main(){
  // 定义指针变量a
var a *int
fmt.Println("a=",a)
num :=1
num1 :=2
fmt.Println("num=",num,"num1=",num1)
var q *int
var p *int
q =&num1
p =&num
fmt.Println("q=",&q)
fmt.Println("num=",p)
fmt.Println("P=",&p)
*p=10
fmt.Println("p=",*p)
var x *int
x =new(int)
fmt.Println("x=",x)
*x =12
fmt.Println("*x=",*x)
}
