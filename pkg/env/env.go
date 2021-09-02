package env

import (
	"os"
	"strconv"
)

//Get get string configuration by defined environtment variable key and default value
func Get(k string, dValue ...string) string {
	val := os.Getenv(k)
	if val == "" && len(dValue) > 0 {
		val = dValue[0]
	}
	return val
}

//GetInt get string configuration by defined environtment variable key and default value
func GetInt(k string, dValue ...int) int {
	str := os.Getenv(k)
	if str == "" && len(dValue) > 0 {
		return dValue[0]
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		if len(dValue) > 0 {
			return dValue[0]
		}
		return 0
	}
	return val
}

//GetBool get boolean value configuration
func GetBool(k string, dValue ...bool) bool {
	str := os.Getenv(k)
	if str == "" && len(dValue) > 0 {
		return dValue[0]
	}
	val, err := strconv.ParseBool(str)
	if err != nil {
		if len(dValue) > 0 {
			return dValue[0]
		}
		return false
	}
	return val
}
