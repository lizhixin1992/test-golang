package main

import "fmt"

/*
Go中聚合跟继承唯一的不同在于，继承自其他结构体的struct类型可以直接访问父类结构体的字段和方法
*/

type TokenType uint16

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	LBRACKET
	RBRACKET
	INT
)

type Token struct {
	Type   TokenType
	Lexeme string
}

type IntegerConstant struct {
	Token *Token
	Value uint64
}

type Pet struct {
	name string
}

type Dog struct {
	Pet
	Breed string
}

func (p *Pet) Play() {
	fmt.Println(p.Speak())
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
}

func main() {
	token := IntegerConstant{&Token{11, "2222"}, 11}
	fmt.Println(token.Token.Lexeme)

	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	fmt.Println(d.name)
	fmt.Println(d.Speak())

	//Dog类型重载了Speak()方法。然而如果Pet有另外一个方法Play()被调用，但是Dog没有实现Play()的时候，Dog类型的Speak()方法则不会被调用
	d.Play()
}
