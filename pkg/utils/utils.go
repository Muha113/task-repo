package utils

import "strings"

func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if strings.Compare(v, value) == 0 {
			return true
		}
	}
	return false
}
