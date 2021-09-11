package redis

import "github.com/shandysiswandi/gokit-service/entity"

func (r *redisCache) GetTodoByID() entity.Todo {
	return entity.Todo{}
}

func (r *redisCache) SetTodoByID() error {
	return nil
}
