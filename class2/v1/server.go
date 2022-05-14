package main

import (
	"class2/controller"
	"class2/repository"
	"class2/tool"
	"encoding/json"
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
	defer tool.NewIdInstance().SaveId()
	if err := Init("./data/"); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/page/post", func(c *gin.Context) {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		var page controller.Page
		json.Unmarshal(buf[0:n], &page)
		err := controller.CreatePageInfo(&page)
		if err != nil {
			fmt.Println(err.Error())
		}
		resp := map[string]string{"msg": "ok"}
		c.JSON(200, resp)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
