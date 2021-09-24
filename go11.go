package main
import "fmt"
func main(){
	str :="go language"
	bytes :=[]byte(str)
	for i :=0;i<2;i++{
		bytes[i] = ''
	}
	fmt.Println(string(str))
}
