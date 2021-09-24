package main
import (
"fmt"
"strings"
)
func main(){
str :="go语言"
index :=strings.Index(str,"言")
fmt.Println(index)
fmt.Println(str[index:])
}
