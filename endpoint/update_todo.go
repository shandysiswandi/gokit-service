package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/service"
)

func makeUpdateTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(entity.UpdateTodoRequest)
			response entity.UpdateTodoResponse
		)

		if err := s.UpdateTodo(ctx, req); err != nil {
			return nil, err
		}

		response.Code = http.StatusOK
		response.Status = http.StatusText(http.StatusOK)
		response.Message = "Success update todo"

		return response, nil
	}
}
