package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// STEP 1：讓所有 SPA 中的檔案可以在正確的路徑被找到
	router.Use(static.Serve("/", static.LocalFile("./../client/build", true)))

	// STEP 2： serve 靜態檔案
	router.Static("/public", "./public")

	// STEP 3：API
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// STEP 4：除了有定義路由的 API 之外，其他都會到前端框架
	// https://github.com/go-ggz/ggz/blob/master/api/index.go
	router.NoRoute(func(ctx *gin.Context) {
		file, _ := ioutil.ReadFile("./../client/build/index.html")
		etag := fmt.Sprintf("%x", md5.Sum(file)) //nolint:gosec
		ctx.Header("ETag", etag)
		ctx.Header("Cache-Control", "no-cache")

		if match := ctx.GetHeader("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				ctx.Status(http.StatusNotModified)

				//這裡若沒 return 的話，會執行到 ctx.Data
				return
			}
		}

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	err := router.Run(":3000") // listen and serve on 0.0.0.0:3000
	if err != nil {
		log.Fatalln("Route can not be run", err)
	}
}
