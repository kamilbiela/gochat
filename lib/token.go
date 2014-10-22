package lib

import (
	"github.com/jmcvetta/randutil"
	"time"
)

type Token struct {
	Val      string
	ExpireAt time.Time
	Username string
}

func (t *Token) IsExpired() bool {
	return t.ExpireAt.After(time.Now())
}

func NewToken(username string) *Token {
	tkn, _ := randutil.AlphaString(32)

	t := new(Token)
	t.Val = tkn
	t.ExpireAt = time.Now().Add(24 * time.Hour)
	t.Username = username

	return t
}
