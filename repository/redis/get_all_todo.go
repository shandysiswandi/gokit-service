package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/shandysiswandi/gokit-service/entity"
)

func (r *redisCache) GetAllTodo(ctx context.Context, k string) []entity.Todo {
	data := make([]entity.Todo, 0)
	r.mu.RLock()
	value, err := r.client.Get(ctx, k).Result()
	r.mu.RUnlock()
	if err != nil {
		return data
	}

	err = json.Unmarshal([]byte(value), &data)
	if err != nil {
		return data
	}

	return data
}

func (r *redisCache) SetAllTodo(ctx context.Context, k string, v []entity.Todo) error {
	bytes, _ := json.Marshal(v)
	r.mu.Lock()
	err := r.client.Set(ctx, k, bytes, 10*time.Second).Err()
	r.mu.Unlock()
	return err
}
