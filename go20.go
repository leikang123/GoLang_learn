package main
import "fmt"
func main(){
// 定义映射，var 变量名 map[键的类型]值类型
var str =map[string]int{
// 初始化
"leikang":34,
"uer" :12,
"junzhang": 28,
"jixaojixao" :76,
}
fmt.Println("str=",str," len is ",len(str))
for k,v :=range str{
fmt.Println("k=",k,"v=",v)
}
fmt.Println("===================")
delete(str,"uer")
for k,v :=range str{
fmt.Println("k=",k,"v=",v)
}
}
