package main

//并发 channel 无缓存 有缓存

import (
	"fmt"
)

func sumNum(a []int, c chan int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	c <- sum
}

func main() {
	//a := []int{7, 2, 8, -9, 4, 0}

	////channel无缓存
	//c := make(chan int)
	//
	//go sumNum(a[len(a)/2:], c)
	//time.Sleep(1 * time.Second)
	//go sumNum(a[:len(a)/2], c)
	//
	//x, y := <-c, <-c
	//fmt.Println(x, y)
	//close(c)

	//channel 有缓存
	c := make(chan int, 2)
	c <- 1
	c <- 2
	for i := range c {
		fmt.Println(i)
	}
	close(c)
}
