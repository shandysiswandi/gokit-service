package redis

import (
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestNewRedis(t *testing.T) {
	db, _ := redismock.NewClientMock()
	defer db.Close()

	type args struct {
		conf Configuration
	}
	testTable := []struct {
		title   string
		args    args
		wantErr bool
	}{
		{
			title:   "Negative Ping",
			args:    args{conf: Configuration{}},
			wantErr: true,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			_, err := NewRedis(tc.args.conf)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func Test_redisCache_Close(t *testing.T) {
	db, _ := redismock.NewClientMock()
	defer db.Close()

	type args struct {
		client *redis.Client
	}
	testTable := []struct {
		title   string
		args    args
		wantErr bool
	}{
		{
			title:   "Positive",
			args:    args{client: db},
			wantErr: false,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			rdb := &redisCache{client: tc.args.client}
			err := rdb.Close()
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
