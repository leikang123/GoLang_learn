package main
import "fmt"
func main(){
var student =[3]int{12,34,39}
for k,v :=range student{
fmt.Println("k",k,"v",v)
var student2 = student[0:3]
fmt.Println("student",student)
fmt.Println("student2",student2)
fmt.Println("student is len",len(student))
fmt.Println("student2 is len",len(student2))
fmt.Println("studentdizhi",&student[1])
fmt.Println("cap",cap(student2))

}
}
