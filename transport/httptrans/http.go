package httptrans

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/shandysiswandi/gokit-service/endpoint"
	"github.com/shandysiswandi/gokit-service/pkg/jwt"
)

var (
	// ErrBadRouting error programmer
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
	// ErrBadRequest
	ErrBadRequest = errors.New("bad request")
)

// NewServer is
func NewServer(end endpoint.Endpoints, jwtSecret string) http.Handler {
	r := mux.NewRouter()

	// set custom not found router
	r.NotFoundHandler = httpNotFound{}
	// set custom method not allowed
	r.MethodNotAllowedHandler = httpMethodNotAllowed{}

	// options
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(jwt.HTTPToContext),
		httptransport.ServerErrorEncoder(httpServerErrorEncoder),
	}

	// middleware
	middJWT := middlewareJWT(jwtSecret)

	// GET /todos retrieve all todo
	r.Methods(http.MethodGet).Path("/todos").Handler(httptransport.NewServer(
		middJWT(end.GetAllTodo),
		decodeGetAllTodoRequest,
		encodeGetAllTodoResponse,
		options...,
	))

	// GET /todos/:id retrieve all todo
	r.Methods(http.MethodGet).Path("/todos/{id}").Handler(httptransport.NewServer(
		middJWT(end.GetTodoByID),
		decodeGetTodoByIDRequest,
		encodeGetTodoByIDResponse,
		options...,
	))

	// POST /todos create todo
	r.Methods(http.MethodPost).Path("/todos").Handler(httptransport.NewServer(
		middJWT(end.CreateTodo),
		decodeCreateTodoRequest,
		encodeCreateTodoResponse,
		options...,
	))

	// PUT /todos/:id update todo by id
	r.Methods(http.MethodPut).Path("/todos/{id}").Handler(httptransport.NewServer(
		middJWT(end.UpdateTodo),
		decodeUpdateTodoRequest,
		encodeUpdateTodoResponse,
		options...,
	))

	// DELETE /todos/:id delete todo by id
	r.Methods(http.MethodDelete).Path("/todos/{id}").Handler(httptransport.NewServer(
		middJWT(end.DeleteTodo),
		decodeDeleteTodoRequest,
		encodeDeleteTodoResponse,
		options...,
	))

	return r
}

// httpNotFound ... (net/http default)
type httpNotFound struct{}

// ServeHTTP ... (net/http default)
func (h httpNotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "route not found"})
}

// httpMethodNotAllowed ... (net/http default)
type httpMethodNotAllowed struct{}

// ServeHTTP ... (net/http default)
func (h httpMethodNotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "route not found"})
}

// httpServerErrorEncoder use for custom from request (decode request gokit-flow) error
//
// check error code 400, 401, 403, 404, default is 500
func httpServerErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	code := http.StatusInternalServerError
	switch err {
	case sql.ErrNoRows:
		code = http.StatusNotFound
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}
