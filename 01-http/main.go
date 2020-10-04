package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

// 原生micro http
func main01() {
	service := web.NewService(web.Address(":8001"))

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	err := service.Run()
	if err != nil {
		fmt.Println(err)
	}
}



// 使用gin框架
func main() {
	r := gin.Default()
	r.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	})
	service := web.NewService(
		web.Name("demo_service"),
		web.Address(":8000"),
		web.Handler(r),
	)
	service.Run()
}