package util

import (
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
