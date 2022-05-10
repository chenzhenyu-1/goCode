package repository

import "sync"

type Topic struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Create_time int64  `json:"create_time"`
}

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) (*Topic, error) {
	return topicIndexMap[id], nil
}
