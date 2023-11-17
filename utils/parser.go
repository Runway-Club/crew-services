package utils

import (
	"strconv"
)

func ParseInt(value string) (int, error) {
	if value == "" {
		return 0, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
