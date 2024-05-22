package utils

import (
	"encoding/json"
)

func Contains(s []string, element string) bool {
	for _, a := range s {
		if a == element {
			return true
		}
	}
	return false
}

func ConvertInterfaceToStruct(paramsMap interface{}, paramsStruct any) error {
	jsonData, err := json.Marshal(paramsMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, paramsStruct)
}

// TrimLeftZeroes returns a substring of s without leading zeroes
func TrimLeftZeroes(s string) string {
	idx := 0
	for ; idx < len(s); idx++ {
		if s[idx] != '0' {
			break
		}
	}
	return s[idx-1:]
}
