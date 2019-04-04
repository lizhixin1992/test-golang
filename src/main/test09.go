package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	//log.SetOutput(os.Stdout)
}

func main() {
	c := make(chan int)
	o := make(chan bool)

	go func() {
		log.Debug("start")
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				log.Debug("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
