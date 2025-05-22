package util

func TernaryCompare[T interface{}](isOk bool, resTrue, resFalse T) T {
	if isOk {
		return resTrue
	}
	return resFalse
}

func Filter[T interface{}](data []T, fn func(T) bool) []T {
	result := make([]T, 0, len(data))
	for _, item := range data {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func ModifySlice[T interface{}, K interface{}](data []T, fn func(T) K) []K {
	result := make([]K, 0, len(data))
	for _, item := range data {
		result = append(result, fn(item))
	}
	return result
}

func SliceToMap[T comparable, K interface{}](data []T, val K) map[T]K {
	result := make(map[T]K, len(data))
	for _, item := range data {
		result[item] = val
	}
	return result
}

func HasInSlice[T comparable](needle T, data []T) bool {
	for i := range data {
		if data[i] == needle {
			return true
		}
	}
	return false
}

func UniqSlice[T comparable](data []T) []T {
	uniqSlice := make([]T, 0, len(data))
	uniqVal := make(map[T]struct{}, len(data))
	for _, item := range data {
		if _, ok := uniqVal[item]; !ok {
			uniqVal[item] = struct{}{}
			uniqSlice = append(uniqSlice, item)
		}
	}
	return uniqSlice
}

func BatchSlice[T interface{}](data []T, count int) [][]T {
	batch := make([][]T, 0, 100)
	tempBatch := make([]T, 0, count)
	for _, item := range data {
		if len(tempBatch) >= count {
			batch = append(batch, tempBatch)
			tempBatch = make([]T, 0, count)
		}
		tempBatch = append(tempBatch, item)
	}
	batch = append(batch, tempBatch)

	return batch
}

func ClearSlice[T interface{}](data []T, fn func(item T) bool) []T {
	result := make([]T, 0, len(data))
	for _, item := range data {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}
