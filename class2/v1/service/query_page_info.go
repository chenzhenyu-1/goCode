package service

import (
	"class2/repository"
	"errors"
	"sync"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

type QueryPageInfoFlow struct {
	pageInfo *PageInfo
	topicId  int64
	topic    *repository.Topic
	posts    []*repository.Post
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	var topicErr, postErr error
	go func() {
		defer wg.Done()
		topic, err := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts, err := repository.NewPostDaoInstance().QueryPostById(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.posts = posts
	}()
	wg.Wait()
	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		PostList: f.posts,
		Topic:    f.topic,
	}
	return nil
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	queryPageInfoFlow := &QueryPageInfoFlow{
		topicId: topicId,
	}
	pageInfo, err := queryPageInfoFlow.Do()
	return pageInfo, err
}
