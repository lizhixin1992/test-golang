package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式, 此处时间格式是当前时间　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.JSONFormatter{})

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

func Info(value string) {
	log.Info(value)
}

func Debug(value string) {
	log.Debug(value)
}

func Error(value string) {
	log.Error(value)
}

func Warn(value string) {
	log.Warn(value)
}

func Fatal(value string) {
	log.Fatal(value)
}

func Panic(value string) {
	log.Panic(value)
}

func main() {
	log.Info("22")
}
