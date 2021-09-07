package httptrans

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shandysiswandi/gokit-service/entity"
)

func decodeGetAllTodoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return entity.GetAllTodoTodoRequest{}, nil
}

func encodeGetAllTodoResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}
