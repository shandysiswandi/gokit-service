package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/shandysiswandi/gokit-service/service"
)

func makeGetTodoByID(s service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(entity.GetTodoByIDTodoRequest)
			response entity.GetTodoByIDTodoResponse
		)

		data, err := s.GetTodoByID(ctx, req)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Status = http.StatusText(http.StatusInternalServerError)
			response.Message = err.Error()
			return response, nil
		}

		response.Code = http.StatusOK
		response.Status = http.StatusText(http.StatusOK)
		response.Message = "Success get todo by id"
		response.Data = data
		return response, nil
	}
}
