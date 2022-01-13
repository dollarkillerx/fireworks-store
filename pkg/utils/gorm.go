package utils

import "strings"

func NotFound(err error) bool {
	if strings.Contains(err.Error(), "not found") {
		return true
	}

	return false
}
