package service

import (
	"errors"
	"sync"
	"testProject/repository"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topId,
	}
}

// QueryPageInfoFlow 页面信息流结构体
type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *repository.Topic
	posts []*repository.Post
}

// Do 结构体函数流程编排 参数校验 -> 准备数据 -> 组装实体
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

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	// 获取topic信息储存到f.topic里面
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicByID(f.topicId)
		f.topic = topic
	}()
	// 获取posts信息储存到f.posts里面
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostByID(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}
func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}
