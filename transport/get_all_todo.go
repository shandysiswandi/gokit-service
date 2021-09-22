package transport

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
	v1 "github.com/shandysiswandi/gokit-service/proto"
)

// GetAllTodo is a
func (s *server) GetAllTodo(ctx context.Context, req *v1.GetAllTodoRequest) (*v1.GetAllTodoResponse, error) {
	_, resp, err := s.getAllTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.GetAllTodoResponse), nil
}

// decodeGetAllTodo is a
func decodeGetAllTodo(ctx context.Context, request interface{}) (interface{}, error) {
	// req := request.(*v1.GetAllTodoRequest)
	return entity.GetAllTodoRequest{}, nil
}

// encodeGetAllTodo is a
func encodeGetAllTodo(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(entity.GetAllTodoResponse)
	var data []*v1.Todo

	for _, todo := range resp.Data {
		data = append(data, &v1.Todo{
			Id:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Status:      v1.Status(todo.Status.ToNumber()),
		})
	}

	return &v1.GetAllTodoResponse{
		Code:    resp.Code,
		Status:  resp.Status,
		Message: resp.Message,
		Data:    data,
	}, nil
}
