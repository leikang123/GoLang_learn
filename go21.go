package main
 import(
 "fmt"
"sync"
)
var lock sync.RWMutex
func main(){
Gmap :=make(map[int] int)
for i :=0; i<1000;i++{
go writerMap(Gmap,i,i)
go readMap(Gmap,i)
}
fmt.Println("Done")
}
func readMap(Gmap map[int] int,key int) int {
lock.Lock()
m :=Gmap[key]
lock.Unlock()
return m
}
func writerMap(Gmap map[int] int,key int,value int){
lock.Lock()
Gmap[key]= value
lock.Unlock()
}
