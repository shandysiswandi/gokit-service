package httptrans

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shandysiswandi/gokit-service/entity"
)

func decodeCreateTodoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req entity.CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrBadRequest
	}

	return req, nil
}

func encodeCreateTodoResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}
