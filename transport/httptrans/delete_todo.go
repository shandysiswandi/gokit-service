package httptrans

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shandysiswandi/gokit-service/entity"
)

func decodeDeleteTodoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	if _, err := strconv.Atoi(id); err != nil {
		return nil, err
	}

	return entity.DeleteTodoRequest{
		ID: id,
	}, nil
}

func encodeDeleteTodoResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}
