package main
import "fmt"

func add(slice ... int) ( int ){
	sum :=0
	for _,value :=range slice{
		sum = sum + value
	}
	return sum
}
func main(){
	fmt.Println("1+2+3+...+10=",add(1,2,3,4,5,6,7,8,9,10))
}