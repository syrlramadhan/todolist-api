package util

import "strconv"

func GetLimitValue(limitStr string) int {
	limit := 5
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}
	return limit
}