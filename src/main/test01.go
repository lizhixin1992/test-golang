package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式, 此处时间格式是当前时间　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	//log.SetFormatter(&log.JSONFormatter{})

	// 设置日志格式为TextFormatter,设置输出的时间为当前时间（不设置的话就是服务启动的时间，比如0000）
	//log.SetFormatter(&log.TextFormatter{
	//	TimestampFormat: "2006-01-02T15:04:05.000",
	//	FullTimestamp: true,
	//})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.DebugLevel)

	//设置输出文件，行号，方法
	log.SetReportCaller(true)
}

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
	//createTime, err := time.Parse("2006-01-02 15:04:05", "2019-04-01 12:00:01")
	//fmt.Println(createTime, err)
	//
	//fmt.Println(createTime.Unix())

	//fmt.Println(utils.GetUnixToHourString(time.Now().Unix()))

	//log.Debug("Useful debugging information.")
	//log.Info("Something noteworthy happened!")
	//log.Warn("You should probably take a look at this.")
	//log.Error("Something failed but I'm not quitting.")
	//log.Fatal("Bye.")         //log之后会调用os.Exit(1)
	//log.Panic("I'm bailing.") //log之后会panic()

	requestLogger := log.WithFields(log.Fields{"request_id": 1, "user_ip": 1}) //对于固定不变的可以这样直接写死
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")

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
