package main
import "fmt"
func main(){
var num int =10
var p * int
p = &num
fmt.Println("num is value ",num)
fmt.Println("num p is value",p)
* p =14
fmt.Println("&p is value",*p)
}

