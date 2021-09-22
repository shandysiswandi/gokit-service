package transport

import (
	"context"

	"github.com/shandysiswandi/gokit-service/entity"
	v1 "github.com/shandysiswandi/gokit-service/proto"
)

func (s *server) GetTodoByID(ctx context.Context, req *v1.GetTodoByIDRequest) (*v1.GetTodoByIDResponse, error) {
	_, resp, err := s.getTodoByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*v1.GetTodoByIDResponse), nil
}

// decodeGetTodoByID is a
func decodeGetTodoByID(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*v1.GetTodoByIDRequest)
	return entity.GetTodoByIDRequest{
		ID: req.Id,
	}, nil
}

// encodeGetTodoByID is a
func encodeGetTodoByID(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(entity.GetTodoByIDResponse)

	return &v1.GetTodoByIDResponse{
		Code:    resp.Code,
		Status:  resp.Status,
		Message: resp.Message,
		Data: &v1.Todo{
			Id:          resp.Data.ID,
			Title:       resp.Data.Title,
			Description: resp.Data.Description,
			Status:      v1.Status(resp.Data.Status.ToNumber()),
		},
	}, nil
}
