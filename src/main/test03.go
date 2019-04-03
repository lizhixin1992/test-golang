package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years - ✆ " + h.phone + "❱"
}

func main() {
	//mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "tow"}
	//	//sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	//	//mark.SayHi()
	//	//sam.SayHi()

	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
