package server

import "testProject/repository"

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).DO=o()
}

func NewQueryPageInfoFlow(topicId int64)

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}

	if err := f.prepareInfo(); err != nil {
		return nil, err
	}

	if err := f.checkParam(); err != nil {
		return nil, err
	}
}
