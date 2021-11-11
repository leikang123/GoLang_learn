// define interface and method name & speak
type Animal interface {
    Name() string
    Speak() string

}
// define struct name cat implement interface
type Cat struct{

}
func (Cat Cat)Name() string{
    return "Cat"
}
func (cat Cat)Speak() string{
    return "miaomiao“

}