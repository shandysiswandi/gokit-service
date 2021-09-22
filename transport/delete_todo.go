package transport

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
	v1 "github.com/shandysiswandi/gokit-service/proto"
)

func (s *server) DeleteTodo(ctx context.Context, req *v1.DeleteTodoRequest) (*v1.DeleteTodoResponse, error) {
	_, resp, err := s.deleteTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.DeleteTodoResponse), nil
}

// decodeDeleteTodo is a
func decodeDeleteTodo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*v1.DeleteTodoRequest)
	return entity.DeleteTodoRequest{
		ID: req.Id,
	}, nil
}

// encodeDeleteTodo is a
func encodeDeleteTodo(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(entity.DeleteTodoResponse)
	return &v1.DeleteTodoResponse{
		Code:    resp.Code,
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
