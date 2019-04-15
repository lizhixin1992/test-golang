package main

import (
	"net/http"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	irisMiddleware := iris.FromStd(nativeTestMiddleware)
	app.Use(irisMiddleware)

	// Method GET: http://localhost:8080/
	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Home")
		ctx.WriteString("11111111")
	})

	// Method GET: http://localhost:8080/ok
	app.Get("/ok", func(ctx iris.Context) {
		ctx.HTML("<b>Hello world!</b>")
		ctx.WriteString("11111111")
	})

	// http://localhost:8080
	// http://localhost:8080/ok
	app.Run(iris.Addr(":8080"))
}

//直接在这里停住了，没有往之后的handler进行
func nativeTestMiddleware(w http.ResponseWriter, r *http.Request) {
	println("Request path: " + r.URL.Path)
}

// Look "routing/custom-context" if you want to convert a custom handler with a custom Context
// to a context.Handler.
