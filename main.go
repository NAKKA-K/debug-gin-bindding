package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Text string `json:"text"`
}

func main() {
	r := gin.Default()
	r.GET("/", get)
	r.POST("/post", post)

	r.Run() //portを指定する場合はr.Run(":10000")みたいに書きます
}

func get(c *gin.Context) {
	fmt.Println("hello")
}

func post(c *gin.Context) {
	fmt.Println("hello")
	var r Request
	if err := c.BindJSON(&r); err != nil {
		panic(err)
	}

	fmt.Printf("r = %+v\n", r)
	fmt.Println("world")
	c.JSON(http.StatusOK, gin.H{"text": r.Text})
}
