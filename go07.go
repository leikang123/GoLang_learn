package main
import  ("fmt" 
      "reflect"
)
func main(){
a :=1
b :="test"
c :=true
fmt.Println(reflect.TypeOf(a))
fmt.Println(reflect.TypeOf(b))
fmt.Println(reflect.TypeOf(c))
}

