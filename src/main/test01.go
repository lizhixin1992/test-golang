package main

import (
	"fmt"
	"utils"
)

func main() {
	//fmt.Println("It is a test01!")
	//var testModel1 TestModel1 = TestModel1{1,"test"}
	//fmt.Println(testModel1)

	////var array = []int{1,2,3,4}
	//var array = make([]int,10,10)
	//array = []int{1,2,3,4}
	//for num := range array {
	//	fmt.Println(array[num])
	//}

	//var testMap = make(map[string]string)
	//testMap["1"] = "1"
	//testMap["2"] = "1"
	//testMap["3"] = "1"
	//testMap["4"] = "1"
	//for key := range testMap {
	//	fmt.Println(testMap[key])
	//}

	//var a int  = 1
	//var b uint = 2
	//fmt.Println(unsafe.Sizeof(a))
	//fmt.Println(unsafe.Sizeof(b))

	//var createTime = time.Now()
	//fmt.Println(createTime)

	//createTime := time.Now().Unix()
	//fmt.Println(createTime)
	//createTimeNano := time.Now().UnixNano()
	//fmt.Println(createTimeNano)

	//var createTime = time.Now().Unix()
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Unix(createTime,0).Format("2006-01-02 15:04:05"))

	//createTime,err := time.Parse("2006-01-02 15:04:05","2019-04-01 12:00:01")
	//fmt.Println(createTime,err)
	//
	//fmt.Println(createTime.Unix())

	fmt.Println(utils.GetNow())

}

type TestModel1 struct {
	id   int
	name string
}

type TestModel2 struct {
	createTime string
	updateTime string
}

type TestModel3 struct {
}
