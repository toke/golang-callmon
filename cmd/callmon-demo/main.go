package main

import (
	"./fritzbox/"
	"fmt"
)

type Test1 struct {
	name string
}

type Test2 struct {
	header Test1
	value  int
}

type Test3 struct {
	Test1
	value int
}

func main() {

	e := fritzbox.EventFromString("06.08.14 14:52:26;CALL;1;10;50000001;012344567;SIP1;\r\n")
	fmt.Println(e)

	m := fritzbox.MessageFromString("06.08.14 14:52:26;CALL;1;10;50000001;012344567;SIP1;\r\n")
	fmt.Println(m)

	header := Test1{name: "HI"}

	message := Test2{header: header, value: 3}
	message2 := Test3{Test1: header, value: 4}

	fmt.Println(message)

	fmt.Println(message.header.name)

	fmt.Println(message2)

}
