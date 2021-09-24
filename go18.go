package main
import "fmt"
func main(){
str :=make([]int,1,1)
for i :=0;i<16;i++{
str = append(str,i)
fmt.Println("当前切片长度:",len(str),"当前切片容量:",cap(str))
}
var s =[] int{11,23,45,41}
for k,v :=range s {
fmt.Println("k",k,"v",v)
}
fmt.Println("s切片:",s)
fmt.Println("s的长度:",len(s))
fmt.Println("s的容量:",cap(s))
fmt.Println("s切片是否为空:",s == nil)
}
