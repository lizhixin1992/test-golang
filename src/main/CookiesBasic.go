package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/kataras/iris"
	"time"
)

func newApp() *iris.Application {
	app := iris.New()

	// Set A Cookie.
	app.Get("/cookies/{name}/{value}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		value := ctx.Params().Get("value")

		ctx.SetCookieKV(name, Sha1String(value), iris.CookieExpires(time.Duration(30)*time.Second)) // <--
		// Alternatively: ctx.SetCookie(&http.Cookie{...})
		//
		// If you want to set custom the path:
		// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))
		//
		// If you want to be visible only to current request path:
		// (note that client should be responsible for that if server sent an empty cookie's path, all browsers are compatible)
		// ctx.SetCookieKV(name, value, iris.CookieCleanPath /* or iris.CookiePath("") */)
		// More:
		//                              iris.CookieExpires(time.Duration)
		//                              iris.CookieHTTPOnly(false)

		ctx.Writef("cookie added: %s = %s", name, value)
	})

	// Retrieve A Cookie.
	app.Get("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		value := ctx.GetCookie(name) // <--
		// If you want more than the value then:
		// cookie, err := ctx.Request().Cookie(name)
		// if err != nil {
		//  handle error.
		// }

		ctx.WriteString(value)
	})

	// Delete A Cookie.
	app.Delete("/cookies/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")

		ctx.RemoveCookie(name) // <--
		// If you want to set custom the path:
		// ctx.SetCookieKV(name, value, iris.CookiePath("/custom/path/cookie/will/be/stored"))

		ctx.Writef("cookie %s removed", name)
	})

	return app
}

func main() {
	app := newApp()

	// GET:    http://localhost:8080/cookies/my_name/my_value
	// GET:    http://localhost:8080/cookies/my_name
	// DELETE: http://localhost:8080/cookies/my_name
	app.Run(iris.Addr(":8080"))
}

func Md5String(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

func Md5(data []byte) string {
	md5 := md5.New()
	md5.Write(data)
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

func HmacString(key string, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte(nil)))
}

func HmacByStringKey(key string, data []byte) string {
	return Hmac([]byte(key), data)
}

func Hmac(key []byte, data []byte) string {
	hmac := hmac.New(md5.New, key)
	hmac.Write(data)
	return hex.EncodeToString(hmac.Sum([]byte(nil)))
}

func Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

func Sha1(data []byte) string {
	sha1 := sha1.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
