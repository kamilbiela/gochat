package main

import (
	"strconv"
)

func GetUserIdForToken(token string) int {
	v, _ := strconv.ParseInt(token, 10, 32)
	return int(v)
}
