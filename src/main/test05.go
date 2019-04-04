package main

//反射

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.1415926
	v := reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v)

	////这个是获取结构体
	//v := reflect.TypeOf(x)
	//fmt.Println(v.Kind() == reflect.Float64)
	//fmt.Println(v)
}
