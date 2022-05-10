package repository

import "sync"

type Post struct {
	Id          int64  `json:"id"`
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
