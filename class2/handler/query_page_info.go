package handler

import (
	"class2/service"
	"strconv"
)

type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func QueryPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: 400,
			Msg:  err.Error(),
		}
	}
	pageInfo, err := service.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{
			Code: 400,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 200,
		Msg:  "success",
		Data: pageInfo,
	}
}
