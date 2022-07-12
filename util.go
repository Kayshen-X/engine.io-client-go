package engine_io_client_go

import (
	"net/url"
)

//DecodeQS Decode Query String,duplicate key will be cover
func DecodeQS(qs string) (map[string]string, error) {
	result := make(map[string]string)
	params, err := url.ParseQuery(qs)
	if err != nil {
		return nil, err
	}

	for key, value := range params {
		result[key] = value[0]
	}

	return result, nil
}

//IsContain judge string array if contains  string
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// MapClone clone a map
func MapClone[T map[string]string](tarMap T, srcMap T) {
	for key, value := range srcMap {
		tarMap[key] = value
	}
}
