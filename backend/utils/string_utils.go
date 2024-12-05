package utils

import "strings"

func LowerSlice(slice []string) []string {
	result := make([]string, len(slice))
	for i, s := range slice {
		result[i] = strings.ToLower(s)
	}
	return result
}
