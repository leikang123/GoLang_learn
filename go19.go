package main
import "fmt"
func main(){
var str = [...]string{"leikang","asd","zxc","wert","iuo"}
var str2 = str[0:3]
fmt.Println("str=:",str)
fmt.Println("str2=:",str2)
for k,v :=range str{
fmt.Println("k=:",k,"v=:",v)
}
for k2,v2 :=range str2{
fmt.Println("k2=:",k2,"v2=:",v2)
}
fmt.Println("str is len ",len(str))
fmt.Println("str is cap ",cap(str))
fmt.Println("str is & ",& str)
fmt.Println(" str2 is &", & str2)
fmt.Println("str2 is len",len(str2))
fmt.Println("str2 is cap",cap(str2))
str2 = append(str2,"yuli","lijun","juju","jinjin","dgfhfhjf")
fmt.Println("str2 is ",str2,"len is ",len(str2),"cap is ",cap(str2))
for k3,v3 :=range str2{
fmt.Println("k3=",k3,"v3=",v3)
// 删除元素
str  = append(str[0:1],str[2:]...)
fmt.Println("str",str)
}
}

