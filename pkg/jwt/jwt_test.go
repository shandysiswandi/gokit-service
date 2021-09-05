package jwt

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPToContext(t *testing.T) {
	testTable := []struct {
		title    string
		inputCtx context.Context
		inputReq func() *http.Request
		want     context.Context
	}{
		{
			title:    "test no authorization header",
			inputCtx: context.Background(),
			inputReq: func() *http.Request {
				return &http.Request{}
			},
			want: context.Background(),
		},
		{
			title:    "test with authorization header",
			inputCtx: context.Background(),
			inputReq: func() *http.Request {
				req := &http.Request{
					Header: make(http.Header),
				}
				req.Header.Add("Authorization", "Bearer token")
				return req
			},
			want: context.WithValue(context.Background(), JWTContextKey, "token"),
		},
	}

	for _, testCase := range testTable {
		actual := HTTPToContext(testCase.inputCtx, testCase.inputReq())

		assert.Equal(t, testCase.want, actual)
	}
}

func Test_extractTokenFromAuthHeader(t *testing.T) {
	testTable := []struct {
		title   string
		input   string
		wantVal string
		wantErr bool
	}{
		{title: "test 1", input: "bearer", wantVal: "", wantErr: true},
		{title: "test 2", input: "beare", wantVal: "", wantErr: true},
		{title: "test 4", input: "bearer token ", wantVal: "", wantErr: true},
		{title: "test 3", input: "bearer ", wantVal: "", wantErr: false},
		{title: "test 4", input: "bearer token", wantVal: "token", wantErr: false},
	}

	for _, testCase := range testTable {
		actVal, actErr := extractTokenFromAuthHeader(testCase.input)

		assert.Equal(t, testCase.wantVal, actVal)
		assert.Equal(t, testCase.wantErr, actErr != nil)
	}
}
