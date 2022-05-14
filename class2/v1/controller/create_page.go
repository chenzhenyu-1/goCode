package controller

import "fmt"

type Page struct {
	Id          int64  `json:"id"`
	User_id     int64  `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Create_time int64  `json:"create_time"`
}

func CreatePageInfo(page *Page) error {
	fmt.Println(page)
	return nil
}
