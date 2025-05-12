package util

import "strconv"

func NotNilValue[T interface{}](data *T) T {
	if data != nil {
		return *data
	}
	var defaultVal T
	return defaultVal
}

func ParseFloatOrZero(str string) float32 {
	if str == "" {
		return 0
	}
	val, _ := strconv.ParseFloat(str, 32)
	return float32(val)
}
