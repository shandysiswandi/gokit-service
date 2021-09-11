package service

import "github.com/go-kit/log"

type Middleware func(TodoService) TodoService

type middleware struct {
	next   TodoService
	logger log.Logger
}

func NewMiddleware(logger log.Logger) Middleware {
	return func(ts TodoService) TodoService {
		return &middleware{
			next:   ts,
			logger: logger,
		}
	}
}
