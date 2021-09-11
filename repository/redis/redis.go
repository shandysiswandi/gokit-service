package redis

type redisCache struct{}

func NewRedis() (*redisCache, error) {
	return nil, nil
}

func (r *redisCache) Close() error {
	return nil
}
