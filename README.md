# [Note] 在 Golang 中 server SPA 的方式（React, Vue）

由於前後端分離的專案中，React 或 Vue 的路由在換頁時並不是真的向 server 發送請求來換頁，所以在 server 端沒設定路由的情況下，當你從首頁進入 App 後進行頁面切換都沒問題，但若在不是從首頁進來，或到其他頁面後按下重新整理時，因為會真的像伺服器發送請求，所以沒特定設定的話，就會得到 404 Not Found。

最常見的作法就是把除了 API 之外的所有路由，都轉回前端 Web App 的檔案，前端 Web App 會再自行根據使用者的網址決定要呈現什麼頁面，如此使用者不論在任何頁面重新整理，都不會出現從 server 端回傳的 404 Not Found，若該路由真的不存在，也是由前端判斷後，顯示 404 NotFound 的頁面。

這裡提供的方式是搭配 GIN 的 `router.NoRoute()` 方法和 [`gin-contrib/static`](https://github.com/gin-contrib/static) middleware，其中的寫法主要是參考自 [Catch all route with static middle](https://github.com/gin-gonic/contrib/issues/90#issuecomment-381546856)：

```go
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
```

## 套用 middleware

在 gin 的 `router.NoRoute()` 中可以放入多個 middleware，例如這裡的 `RequestIDMiddleware()` 和 `AppleCache`：

```go
router.NoRoute(RequestIDMiddleware(), ApplyCache, func(ctx *gin.Context) {
  // ...
}
```

## 使用

```bash
# 打包前端專案
$ cd client
$ npm run build

# 後端啟動 GIN server
$ cd server
$ go run main.go
```

## 專案程式碼

> [Template: Golang Gin Spa](https://github.com/pjchender/golang-gin-spa-template) @ Github

## 參考資料

- [ ] [Catch all route with static middle](https://github.com/gin-gonic/contrib/issues/90#issuecomment-381546856) @ Github Issues
- [ ] [go-ggz/ggz/api/index.go](https://github.com/go-ggz/ggz/blob/master/api/index.go) @ GitHub
- [ ] [go-ggz/ggz/pkg/router/routes.go](https://github.com/go-ggz/ggz/blob/master/pkg/router/routes/routes.go#L91) @ Github
- [ ] [gin-spa](https://github.com/mandrigin/gin-spa) @ Github
