package todo_queries_getById

import (
	"github.com/crosstech/xblocks-golang"
)

type RequestModel struct {
	xblocks.IRequestModel
	Id int `json:"id"`
}

type ResponseModel struct {
	xblocks.IResponseModel
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
