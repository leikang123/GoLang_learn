package main

import (
	"errors"
	"fmt"
)

type Database interface {
	Connector() error
	Disconnector() error
}
type Mysql struct {
	DBname      string
	isConnector bool
}

func (mysql *Mysql) Connector() error {
	fmt.Println("Mysql Connector DB =>" + mysql.DBname)
	mysql.isConnector = true

	if mysql.isConnector {
		fmt.Println("Mysql Connector Success!")
		return nil
	} else {
		return errors.New("Connector failure!")
	}
}
func (mysql *Mysql) Disconnector() error {
	fmt.Println("Mysql Disconnector Success!")
	return nil
}
func main() {
	var mysql = Mysql{DBname: "sudent"}
	fmt.Println("start connector")
	mysql.Connector()
	fmt.Println("fail connector")
	mysql.Disconnector()
}
