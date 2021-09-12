package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/shandysiswandi/gokit-service/entity"
)

func (r *redisCache) GetTodoByID(ctx context.Context, k string) entity.Todo {
	var data entity.Todo

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

func (r *redisCache) SetTodoByID(ctx context.Context, k string, v entity.Todo) error {
	bytes, _ := json.Marshal(v)
	r.mu.Lock()
	err := r.client.Set(ctx, k, bytes, time.Second*10).Err()
	r.mu.Unlock()
	return err
}
