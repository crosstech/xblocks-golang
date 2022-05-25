package todoController

import (
	"net/http"
	"strconv"

	"github.com/crosstech/xblocks-golang"
	todo_queries_getById "github.com/crosstech/xblocks-golang/example/handlers/todo/queries/getById"
	"github.com/gorilla/mux"
)

func GetById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	id, _ := strconv.ParseInt(params["id"], 10, 32)

	endpoint := todo_queries_getById.Endpoint{
		RequestPayload: todo_queries_getById.RequestModel{
			Id: int(id),
		},
	}

	xblocks.OperateHttpRequest(response, endpoint)
}
