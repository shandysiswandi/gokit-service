package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/service"
)

func makeDeleteTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(entity.DeleteTodoRequest)
			response entity.DeleteTodoResponse
		)

		if err := s.DeleteTodo(ctx, req); err != nil {
			return nil, err
		}

		response.Code = http.StatusOK
		response.Status = http.StatusText(http.StatusOK)
		response.Message = "Success delete todo"

		return response, nil
	}
}
