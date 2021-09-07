package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/service"
)

func makeGetAllTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(entity.GetAllTodoTodoRequest)
			response entity.GetAllTodoTodoResponse
		)

		data, err := s.GetAllTodo(ctx, req)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Status = http.StatusText(http.StatusInternalServerError)
			response.Message = err.Error()
			return response, nil
		}

		response.Code = http.StatusOK
		response.Status = http.StatusText(http.StatusOK)
		response.Message = "Success get all todo"
		response.Data = data
		return response, nil
	}
}
