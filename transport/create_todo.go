package transport

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
	v1 "github.com/shandysiswandi/gokit-service/proto"
)

// CreateTodo is a
func (s *server) CreateTodo(ctx context.Context, req *v1.CreateTodoRequest) (*v1.CreateTodoResponse, error) {
	_, resp, err := s.deleteTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.CreateTodoResponse), nil
}

// decodeCreateTodo is a
func decodeCreateTodo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*v1.CreateTodoRequest)
	return entity.CreateTodoRequest{
		Title:       req.Title,
		Description: req.Description,
	}, nil
}

// encodeCreateTodo is a
func encodeCreateTodo(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(entity.CreateTodoResponse)
	return &v1.CreateTodoResponse{
		Code:    resp.Code,
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
