package main
// 倒入的包名或者库
import "fmt"
// 主函数
func main(){
// 定义变量名的简单写法:变量名 := 数值
x,y :=7,9
fmt.Println(" x + y ",x+y)
fmt.Println(" x - y ",x-y)
fmt.Println(" x * y ",x*y)
fmt.Println(" x / y ",x/y)
fmt.Println(" x % y ",x%y)
isbool :=true
hate :=false
fmt.Println(isbool && hate)
fmt.Println(isbool || hate)
fmt.Println(isbool)
a,b,c :="A","中国",98
fmt.Println("a=",a,"b=",b,"c=",c)
}

