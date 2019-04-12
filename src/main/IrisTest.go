package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", index)

	party := app.Party("/party", partyTest)

	party.Get("/ping", ping)

	party.Get("/hello", hello)

	users := app.Party("/users", myAuthMiddlewareHandler)
	// http://localhost:8080/users/42/profile
	users.Get("/{id:int}/profile", userProfileHandler)
	// http://localhost:8080/users/inbox/1
	users.Get("/inbox/{id:int}", userMessageHandler)

	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./src/main/iris.yml")))
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>Welcome</h1>")
}

func partyTest(ctx iris.Context) {
	ctx.WriteString("Authentication failed")
	ctx.Next()
}

func ping(ctx iris.Context) {
	ctx.WriteString("111")
}

func hello(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
}

func myAuthMiddlewareHandler(ctx iris.Context) {
	ctx.WriteString("Authentication failed")
	//允许运行后续的逻辑
	ctx.Next()
}
func userProfileHandler(ctx iris.Context) { //
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
func userMessageHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.WriteString(id)
}
