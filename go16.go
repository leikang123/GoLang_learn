package main
import "fmt"
func fibonacci(n int) (res int){
a :=1
b :=1
for i :=2;i<n;i++{
c :=b
b =a+b
a =c
}
return b
}
func main(){
n :=9 
fmt.Println("n=9:",n,fibonacci(n))
}
