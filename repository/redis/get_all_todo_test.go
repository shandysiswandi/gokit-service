package redis

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/shandysiswandi/gokit-service/entity"
	"github.com/stretchr/testify/assert"
)

func Test_redisCache_GetAllTodo(t *testing.T) {
	db, mock := redismock.NewClientMock()
	defer db.Close()

	type args struct {
		ctx context.Context
		key string
	}
	testTable := []struct {
		title   string
		args    args
		mocking func(a args)
		want    []entity.Todo
	}{
		{
			title:   "Negative r.client.Get",
			args:    args{ctx: context.Background(), key: "some"},
			mocking: func(a args) {},
			want:    []entity.Todo{},
		},
		{
			title: "Negative json.Unmarshal",
			args:  args{ctx: context.Background(), key: "some"},
			mocking: func(a args) {
				mock.ExpectGet(a.key).SetVal("error")
			},
			want: []entity.Todo{},
		},
		{
			title: "Positive",
			args:  args{ctx: context.Background(), key: "some"},
			mocking: func(a args) {
				val, _ := json.Marshal([]entity.Todo{})
				mock.ExpectGet(a.key).SetVal(string(val))
			},
			want: []entity.Todo{},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			tc.mocking(tc.args)
			rdb := redisCache{client: db}
			data := rdb.GetAllTodo(tc.args.ctx, tc.args.key)
			assert.Equal(t, tc.want, data)
			mock.ExpectationsWereMet()
		})
	}
}

func Test_redisCache_SetAllTodo(t *testing.T) {
	db, mock := redismock.NewClientMock()
	defer db.Close()

	type args struct {
		ctx   context.Context
		key   string
		value []entity.Todo
	}
	testTable := []struct {
		title   string
		args    args
		mocking func(a args)
		wantErr bool
	}{
		{
			title:   "Negative r.client.Get",
			args:    args{ctx: context.Background(), key: "some"},
			mocking: func(a args) {},
			wantErr: true,
		},
		{
			title: "Positive",
			args:  args{ctx: context.Background(), key: "some", value: make([]entity.Todo, 0)},
			mocking: func(a args) {
				val, _ := json.Marshal(a.value)
				mock.ExpectSet(a.key, val, time.Second*10).SetVal(string(val))
			},
			wantErr: false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.title, func(t *testing.T) {
			tc.mocking(tc.args)
			rdb := redisCache{client: db}
			err := rdb.SetAllTodo(tc.args.ctx, tc.args.key, tc.args.value)
			log.Println(err)
			assert.Equal(t, tc.wantErr, err != nil)
			mock.ExpectationsWereMet()
		})
	}
}
