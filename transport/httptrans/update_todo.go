package httptrans

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shandysiswandi/gokit-service/entity"
)

func decodeUpdateTodoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	var req entity.UpdateTodoRequest
	req.ID = id
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrBadRequest
	}

	return req, nil
}

func encodeUpdateTodoResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}
