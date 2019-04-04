package main

//测试Web

import (
	. "fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//Println(r.Form)
	//Println("Host",r.Host)
	for k := range r.Header {
		Println(k + ":" + r.Header.Get(k))
	}
	//Println("URL", r.URL)

	//for k, v := range r.Form {
	//	Println("key", k)
	//	Println("val", strings.Join(v, " "))
	//}
	Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
