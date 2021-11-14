package main
import "fmt"


// 使用可变参数实现任意数进行计算
func add(slice ... int) ( int ){
	sum :=0
	// 遍历数字
	for _,value :=range slice{
		sum = sum + value
	}
	return sum
}
func main(){
	fmt.Println("1+2+3+...+10=",add(1,2,3,4,5,6,7,8,9,10))
}