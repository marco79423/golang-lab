package main

import (
	"embed"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

/**
例子：
public/
	index.html
static/
	1.png

main.go
*/

//go:embed static/*
//go:embed public/*
var rayFS embed.FS

//go:embed public/index.html
var indexFs []byte

func main() {
	router := gin.Default()

	// 嘗試 1：
	// * /way-1 失敗
	// * /way-1/index.html 失敗
	// * /way-1/1.png 成功
	router.GET("/way-1/index.html", func(c *gin.Context) {
		c.File("./public/index.html")
	})
	router.GET("/way-1/1.png", func(c *gin.Context) {
		c.File("./static/1.png")
	})

	// 嘗試 2：
	// * /way-2 成功
	// * /way-2/index.html 失敗
	// * /way-2/1.png 成功
	router.GET("/way-2", func(c *gin.Context) {
		c.File("./public/index.html")
	})
	router.GET("/way-2/1.png", func(c *gin.Context) {
		c.File("./static/1.png")
	})

	// 嘗試 3：
	// * /way-3 成功 (會讀到 index.html)
	// * /way-3/index.html 成功
	// * /way-3/1.png 成功
	router.Use(static.Serve("/way-3", static.LocalFile("./static", false)))
	router.Use(static.Serve("/way-3", static.LocalFile("./public", false)))

	// 嘗試 4：
	// * /way-4 成功 (會讀到檔案索引) -> 如果改注冊順序的話，會變成 index.html
	// * /way-4/index.html 成功 (會讀到檔案索引)
	// * /way-4/1.png 成功
	router.Use(static.Serve("/way-4", static.LocalFile("./static", true)))
	router.Use(static.Serve("/way-4", static.LocalFile("./public", true)))

	// 嘗試 5：
	// * /way-5 失敗
	// * /way-5/index.html 失敗
	// * /way-5/1.png 成功
	router.Static("/way-5", "./static")
	// router.Static("/way-5", "./public") (不能重覆注冊)

	// 嘗試 6：
	// * /way-6 成功
	// * /way-6/index.html 成功
	// router.Static("/way-6", "./static") (不能重覆注冊)
	router.Static("/way-6", "./public")

	// 嘗試 7：
	// * /way-7 成功 (會讀到檔案索引)
	// * /way-7/public 失敗
	// * /way-7/public/index.html 失敗
	// * /way-7/static 失敗
	// * /way-7/static/index.html 失敗
	// * /way-7/static/1.png 成功
	router.StaticFS("/way-7", http.FS(rayFS))

	// 嘗試 8：
	// * /way-8 成功
	// * /way-8/index.html 失敗
	router.GET("/way-8", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", indexFs)
	})

	// 嘗試 9：
	// * /way-9 成功 失敗
	// * /way-9/index.html  成功
	router.GET("/way-9/index.html", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html", indexFs)
	})

	// 嘗試 10：
	// * /way-10 失敗
	// * /way-10/index.html  成功
	router.GET("/way-10/index.html", func(ctx *gin.Context) {
		file, _ := rayFS.Open("public/index.html")
		data, _ := ioutil.ReadAll(file)
		ctx.Data(http.StatusOK, "text/html", data)
	})

	router.Run()
}
