
package main
import (
"fmt"
"reflect"
"bytes"
) 
func main(){
 a :="123"
 b :="987"
var c bytes.Buffer
c.WriteString(a)
c.WriteString(b)
fmt.Println(c.String())
fmt.Println(reflect.TypeOf(c))
}

