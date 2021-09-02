package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	testTable := []struct {
		title    string
		key      string
		defaults []string
		want     string
	}{
		{
			title:    "test get 1",
			key:      "USER",
			defaults: nil,
			want:     "shandy",
		},
		{
			title:    "test get 1",
			key:      "USERS",
			defaults: []string{"me"},
			want:     "me",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.title, func(t *testing.T) {
			actual := Get(testCase.key, testCase.defaults...)

			if testCase.want != actual {
				t.Errorf("Get Error => want %v, but got %v", testCase.want, actual)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("TEST_NUMBER", "10")

	testTable := []struct {
		title    string
		key      string
		defaults []int
		want     int
	}{
		{
			title:    "test getInt 1",
			key:      "TEST_NUMBER",
			defaults: nil,
			want:     10,
		},
		{
			title:    "test getInt 2",
			key:      "",
			defaults: []int{10},
			want:     10,
		},
		{
			title:    "test getInt 3",
			key:      "USER",
			defaults: nil,
			want:     0,
		},
		{
			title:    "test getInt 4",
			key:      "USER",
			defaults: []int{10},
			want:     10,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.title, func(t *testing.T) {
			actual := GetInt(testCase.key, testCase.defaults...)

			if testCase.want != actual {
				t.Errorf("GetInt Error => want %v, but got %v", testCase.want, actual)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	os.Setenv("TEST_BOOLEAN_1", "true")
	os.Setenv("TEST_BOOLEAN_2", "false")
	os.Setenv("TEST_BOOLEAN_3", "0")
	os.Setenv("TEST_BOOLEAN_4", "1")

	testTable := []struct {
		title    string
		key      string
		defaults []bool
		want     bool
	}{
		{
			title:    "test getBool 1",
			key:      "TEST_BOOLEAN_1",
			defaults: nil,
			want:     true,
		},
		{
			title:    "test getBool 2",
			key:      "TEST_BOOLEAN_2",
			defaults: nil,
			want:     false,
		},
		{
			title:    "test getBool 3",
			key:      "TEST_BOOLEAN_3",
			defaults: nil,
			want:     false,
		},
		{
			title:    "test getBool 4",
			key:      "TEST_BOOLEAN_4",
			defaults: nil,
			want:     true,
		},
		{
			title:    "test getBool 5",
			key:      "USERS",
			defaults: []bool{true},
			want:     true,
		},
		{
			title:    "test getBool 6",
			key:      "USER",
			defaults: []bool{true},
			want:     true,
		},
		{
			title:    "test getBool 7",
			key:      "USERS",
			defaults: nil,
			want:     false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.title, func(t *testing.T) {
			actual := GetBool(testCase.key, testCase.defaults...)

			if testCase.want != actual {
				t.Errorf("GetBool Error => want %v, but got %v", testCase.want, actual)
			}
		})
	}
}
