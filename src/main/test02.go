package main

import "fmt"

func main() {
	// var a int
	// var b = true;
	// var c = "this a string type"
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// e := false
	// var f,g = 1,2
	// fmt.Println(e)
	// fmt.Println(f,g)

	// var r1 = 1
	// var r2 = 1
	// fmt.Println(&r1,&r2)

	// a,b,c := numbers()
	// fmt.Println(a,b,c)

	// const s = 1
	// s = s + 1
	// fmt.Println(s)

	// const s = iota
	// fmt.Println(s)

	// const (
	// 	a = iota
	// 	b = iota
	// 	c = iota
	// )
	// fmt.Println(a, b, c)
	// const (
	// 	d = 0
	// 	e = 1
	// 	f = 2
	// )
	// fmt.Println(d, e, f)

	// const (
	//            a = iota   //0
	//            b          //1
	//            c          //2
	//            d = "ha"   //独立值，iota += 1
	//            e          //"ha"   iota += 1
	//            f = 100    //iota +=1
	//            g          //100  iota +=1
	//            h = iota   //7,恢复计数
	//            i          //8
	//    )
	//    fmt.Println(a,b,c,d,e,f,g,h,i)
	//    fmt.Printf("this is a %d \n",a)

	// var a [10] int
	// fmt.Println(len(a))

	// var a = [5] int {1,2,3,4,5}
	// a[0] = 22
	// fmt.Println(a)

	// for i := 0; i < len(a); i++{
	// 	fmt.Println(a[i])
	// }

	// var i = 1
	// var a *int = &i
	// fmt.Println(a)
	// fmt.Println(a == nil)

	// a := AnyModel{1,"name","subTitle","createTime"}
	// a = printlnModel(a)
	// printlnModel(a)

	// s := make([]int,5,5)
	// fmt.Println(len(s),cap(s))
	// var a []int
	// fmt.Println(a == nil)

	// s := []int{1,2,3,4,5,6,7,8,9}
	// // fmt.Println(s)
	// fmt.Println(s[0:])
	// s = append(s,0)
	// fmt.Println(s)

	// s1 := make([]int, len(s), cap(s)*2)
	// fmt.Println(len(s1),cap(s1))
	// fmt.Println(s1)
	// copy(s1,s)
	// fmt.Println(s1)

	// s := []int{11,12,13,14,15,16,17,18,19}
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }
	// for num := range s{
	// 	fmt.Println(num,s[num])
	// }

	// var testMap map[string]string
	// testMap := make(map[string]string)
	// fmt.Println(len(testMap))
	// testMap["1"] = "1"
	// testMap["2"] = "1"
	// testMap["3"] = "1"
	// testMap["4"] = "1"
	// // for i, num := range testMap{
	// // 	fmt.Println(i,testMap[num])
	// // }
	// delete(testMap,"4")
	// for i, num := range testMap{
	// 	fmt.Println(i,testMap[num])
	// }

	// var i uint64 = 15
	// fmt.Println(OneTo(i))

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d\t", fibonacci(i))
	// 	if(i == 9){
	// 		fmt.Println("\n")
	// 	}
	// }

	// var a int = 5
	// var b float64 = 11.11
	// var c int = 2
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println("-----------------------")
	// fmt.Println(float64(a)/float64(c))

	// var phone Phone
	//    phone = new(NokiaPhone)
	//    phone.call()

	//s := [] int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c
	//fmt.Println(x,y,x + y)

}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

// func OneTo(i int)(int){
// 	var a int = 1
// 	if(i > 0){
// 		a = i * OneTo(i-1)
// 		return a
// 	}
// 	return 1
// }

func OneTo(n uint64) (result uint64) {
	if n > 0 {
		result = n * OneTo(n-1)
		return result
	}
	return 1
}

func printlnModel(a AnyModel) AnyModel {
	fmt.Println(a.id)
	fmt.Println(a.name)
	fmt.Println(a.subTitle)
	fmt.Println(a.createTime)
	a.id = 9999
	return a
}

type AnyModel struct {
	id         int
	name       string
	subTitle   string
	createTime string
}

func numbers() (int, int, int) {
	return 1, 2, 3
}
