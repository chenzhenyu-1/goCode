package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

func InitTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func InitPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post []*Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}
		postTmpMap[post[0].Id] = post
	}
	postIndexMap = postTmpMap
	return nil
}
