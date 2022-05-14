package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Post struct {
	Id          int64  `json:"id"`
	User_id     int64  `json:"user_id"`
	Parent_id   int    `json:"parent_id"`
	Content     string `json:"content"`
	Create_time int64  `json:"create_time"`
}

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostById(id int64) ([]*Post, error) {
	return postIndexMap[id], nil
}

func (*PostDao) AppendPost(post *Post) error {
	filePath := ".\\data\\post"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	text, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, item := range text {
		write.WriteByte(item)
	}
	write.WriteByte(byte('\n'))
	write.Flush()
	return nil
}

func (f *PostDao) GetPostMaxId() int64 {
	n := postIndexMap
	var counter int64 = 0
	for _, item := range n {
		k := len(item)
		counter = counter + int64(k)
	}
	return counter + 1
}
