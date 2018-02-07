package utils

import (
	"strconv"
)

func StringToUint(s string) (uint, error){
	pid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0, err
	}
	id := uint(pid)
	return id, nil
}