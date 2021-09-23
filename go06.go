package main
import "fmt"
func main(){
// 遍历数组
EvenNum :=[5]int {1,2,3,4,5}
for i,value := range EvenNum{
fmt.Println(value,i)
}
//定义切片的demo
sliceNum :=[]int {9,8,7,6,5}
sliced :=sliceNum[0:]
fmt.Println(sliced)

}
