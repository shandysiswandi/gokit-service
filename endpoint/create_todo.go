package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/service"
)

func makeCreateTodo(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(entity.CreateTodoRequest)
			response entity.CreateTodoResponse
		)

		if err := s.CreateTodo(ctx, req); err != nil {
			return nil, err
		}

		response.Code = http.StatusOK
		response.Status = http.StatusText(http.StatusOK)
		response.Message = "Success create todo"

		return response, nil
	}
}
