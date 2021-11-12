package main

import "fmt"

type User struct {
	name    string
	age     int
	mail    int
	address string
}

// 方法
func (u *User) changeName() {
	u.name = "Tom"
	u.address = "wanghu"
}
func (b *User) method() {
	b.name = "bobo"
	b.age = 88
	b.mail = 8766567
	b.address = "jijulhkjljhsdg"

}
func (s *User) str() {
	s.name = "lilei"
	s.age = 34
	s.mail = 2345
	s.address = "jikljhhjfjhv"
}
func (d User) asd() {
	// 初始化值不传递，相当于没写，只有在主类初始化才可以传递value 
	d.name = "asd"
	d.age = 23
	d.mail = 8765
	d.address = "bjhdgfh"
}

// 初始化
func main() {
	user := &User{
		"leikang",
		33,
		10134565,
		"hubeiwuhan",
	}
	fmt.Println("name:", user.name)
	fmt.Println("age:", user.age)
	fmt.Println("mail:", user.mail)
	fmt.Println("address:", user.address)
	user.changeName()
	fmt.Println("name:", user.name, "address:", user.address)
	user.method()
	fmt.Println("name:", user.name, "mail:", user.mail, "age:", user.age, "address:", user.address)
	user.str()
	fmt.Println("name:", user.name, "mail:", user.mail, "age:", user.age, "address:", user.address)
	d := &User{}
	d.asd()
	fmt.Println("name:", d.name, "age:", d.age, "mail:", d.mail, "address:", d.address)
}
