package bcrypt

import (
	"testing"
)

func TestHash(t *testing.T) {
	testTable := []struct {
		title   string
		plain   string
		wantErr error
	}{
		{
			title:   "test hash 1",
			plain:   "test",
			wantErr: nil,
		},
		{
			title:   "test hash 2",
			plain:   "testing",
			wantErr: nil,
		},
	}

	for _, testcase := range testTable {
		t.Run(testcase.title, func(t *testing.T) {
			_, err := Hash(testcase.plain)

			if err != testcase.wantErr {
				t.Errorf("Expected %v, but got %v", testcase.wantErr, err)
			}
		})
	}
}

func BenchmarkHash(b *testing.B) {
	p := "testing"
	for i := 0; i < b.N; i++ {
		Hash(p)
	}
}

func TestCompare(t *testing.T) {
	testTable := []struct {
		title   string
		hash    string
		plain   string
		wantErr error
	}{
		{
			title:   "test compare 1",
			hash:    "$2a$14$7RrBIDYd4HGumuHIGZPkUOHxaM4he1GTdPtl0vbSbvwHkI2a8U0JK",
			plain:   "test",
			wantErr: nil,
		},
		{
			title:   "test compare 2",
			hash:    "$2a$14$T75vyxNVifUMOv9sdgzvoO5jRlN.1PpuVZpj2FNRZLjroNBsZJrBO",
			plain:   "testing",
			wantErr: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.title, func(t *testing.T) {
			err := Compare(testCase.hash, testCase.plain)

			if err != testCase.wantErr {
				t.Errorf("Expected %v, but got %v", testCase.wantErr, err)
			}
		})
	}
}

func BenchmarkCompare(b *testing.B) {
	h := "$2a$14$T75vyxNVifUMOv9sdgzvoO5jRlN.1PpuVZpj2FNRZLjroNBsZJrBO"
	p := "testing"
	for i := 0; i < b.N; i++ {
		Compare(h, p)
	}
}

func TestIsValid(t *testing.T) {
	testTable := []struct {
		title string
		hash  string
		plain string
		want  bool
	}{
		{
			title: "test isValid 1",
			hash:  "$2a$14$7RrBIDYd4HGumuHIGZPkUOHxaM4he1GTdPtl0vbSbvwHkI2a8U0JK",
			plain: "test",
			want:  true,
		},
		{
			title: "test isValid 2",
			hash:  "$2a$14$T75vyxNVifUMOv9sdgzvoO5jRlN.1PpuVZpj2FNRZLjroNBsZJrBO",
			plain: "testing",
			want:  true,
		},
	}

	for _, testcase := range testTable {
		t.Run(testcase.title, func(t *testing.T) {
			isVal := IsValid(testcase.hash, testcase.plain)

			if isVal != testcase.want {
				t.Errorf("Expected %v, but got %v", testcase.want, isVal)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	h := "$2a$14$T75vyxNVifUMOv9sdgzvoO5jRlN.1PpuVZpj2FNRZLjroNBsZJrBO"
	p := "testing"
	for i := 0; i < b.N; i++ {
		IsValid(h, p)
	}
}
