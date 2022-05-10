package main

import (
	"class2/handler"
	"class2/repository"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Init(filePath string) error {
	err := repository.InitPostIndexMap(filePath)
	if err != nil {
		return err
	}
	err = repository.InitTopicIndexMap(filePath)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Init("./data/"); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := handler.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
