package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Topic struct {
	Id          int64  `json:"id"`
	User_id     int64  `json:"user_id"`
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
func (*TopicDao) CreateTopic(topic *Topic) error {
	filepath := ".\\data\\topic"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	text, err := json.Marshal(topic)
	if err != nil {
		return err
	}
	for _, item := range text {
		write.WriteByte(item)
	}
	write.WriteByte(byte('\n'))
	write.Flush()
	return nil
}
