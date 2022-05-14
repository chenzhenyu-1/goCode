package service

import (
	"class2/repository"
	"class2/tool"
)

type PostData struct {
	User_id     int64  `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Parent_id   int    `json:"parent_id"`
	Create_time int64  `json:"create_time"`
}

func CreatePost(post *PostData) error {
	var newPost repository.Post
	newPost.User_id = post.User_id
	newPost.Content = post.Content
	newPost.Create_time = post.Create_time
	newPost.Parent_id = post.Parent_id
	newPost.Id = tool.NewIdInstance().GetID()
	err := repository.NewPostDaoInstance().AppendPost(&newPost)
	if err != nil {
		return err
	}
	return nil
}
