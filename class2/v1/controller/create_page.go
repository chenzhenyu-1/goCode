package controller

import (
	"class2/service"
	"fmt"
)

type Page struct {
	Data service.PostData
}

func CreatePageInfo(page *Page) error {
	fmt.Println(page)
	post := page.Data
	err := service.CreatePost(&post)
	if err != nil {
		return err
	}
	return nil
}
