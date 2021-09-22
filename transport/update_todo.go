package transport

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
	v1 "github.com/shandysiswandi/gokit-service/proto"
)

func (s *server) UpdateTodo(ctx context.Context, req *v1.UpdateTodoRequest) (*v1.UpdateTodoResponse, error) {
	_, resp, err := s.updateTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.UpdateTodoResponse), nil
}

// decodeUpdateTodo is a
func decodeUpdateTodo(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*v1.UpdateTodoRequest)
	return entity.UpdateTodoRequest{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Status:      entity.Status(req.Status.String()),
	}, nil
}

// encodeUpdateTodo is a
func encodeUpdateTodo(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(entity.UpdateTodoResponse)
	return &v1.UpdateTodoResponse{
		Code:    resp.Code,
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
