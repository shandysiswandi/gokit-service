package entity

import (
	"strings"
)

type Status string

const (
	DONE  Status = "done"
	DRAFT Status = "draft"
)

func (s Status) ToUpper() string {
	return strings.ToUpper(string(s))
}

func (s Status) ToLower() string {
	return strings.ToLower(string(s))
}

func (s Status) ToNumber() int32 {
	if s == DONE {
		return 4
	}
	return 0
}

type Todo struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      Status `json:"status,omitempty"`
}

type GetAllTodoTodoRequest struct{}
type GetAllTodoTodoResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetTodoByIDTodoRequest struct {
	ID string `json:"id"`
}
type GetTodoByIDTodoResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type CreateTodoResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UpdateTodoRequest struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      Status `json:"status,omitempty"`
}
type UpdateTodoResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DeleteTodoRequest struct {
	ID string `json:"id"`
}
type DeleteTodoResponse struct {
	Code    int32       `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
