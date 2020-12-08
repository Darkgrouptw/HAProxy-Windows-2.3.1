package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 最大支援多少大小的檔案 8 MB (Default 32 MB)
	router.MaxMultipartMemory = 8 << 20

	router.GET("/", func(c *gin.Context) {
		c.String(200, "測試1")
	})

	router.Run(":10001")
}
