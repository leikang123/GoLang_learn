package main
import (
"fmt"
// 引入runtime包
"runtime"
)
func main(){
// 调用runtime。NumCPU()方:法
if num :=runtime.NumCPU();
   num >1{
fmt.Println("cpu=:",num)
}
}
