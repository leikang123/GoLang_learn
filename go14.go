package main
import "fmt"
func main(){
  for i :=0;i<4;i++{
    for j :=0;j<4;j++{
      if i == 3{
          break
}
if j == 3{
continue
}
      fmt.Println("i *j",i*j)
}
}
}
