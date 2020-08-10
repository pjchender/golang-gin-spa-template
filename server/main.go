package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"path/filepath"
)

func main() {
	router := gin.Default()

	// 讓所有 SPA 中的檔案可以在正確的路徑被找到
	router.Use(static.Serve("/", static.LocalFile("./../client/build", true)))

	// API
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 除了有定義路由的 API 之外，其他都會到前端框架
	router.NoRoute(func(ctx *gin.Context) {
		dir, file := path.Split(ctx.Request.RequestURI)

		ext := filepath.Ext(file)
		fmt.Printf("dir is %v; file is %v, ext is %v", dir, file, ext)

		if file == "" || ext == "" {
			// 如果網址最後沒有檔名或副檔名，則提供 SPA 的靜態網站首頁
			ctx.File("./../client/build/index.html")
		} else {
			// 如果有檔案和檔名，則提供該檔案回去
			ctx.File("./../client/public" + path.Join(dir, file))
		}
	})

	router.Run(":3000") // listen and serve on 0.0.0.0:3000
}
