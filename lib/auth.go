package lib

import (
	"github.com/jmcvetta/randutil"
	dry "github.com/ungerik/go-dry"
	"time"
)

type AuthService struct {
	tokens *dry.SyncMap
}

type Token struct {
	Val      string
	ExpireAt time.Time
	Username string
}

func NewToken(username string) *Token {
	tkn, _ := randutil.AlphaString(32)

	t := new(Token)
	t.Val = tkn
	t.ExpireAt = time.Now().Add(24 * time.Hour)
	t.Username = username

	return t
}

func (as *AuthService) IsValid(username string, password string) bool {
	return true
}

func (as *AuthService) GetToken(token string) *Token {
	t := as.tokens.Get(token)
	tok, ok := t.(Token)

	if !ok {
		return nil
	}

	return &tok
}

func (as *AuthService) GenerateToken(username string) Token {

	token := *NewToken(username)
	as.tokens.Add(token.Val, token)

	return token
}

func NewAuthService() *AuthService {
	auth := new(AuthService)
	auth.tokens = dry.NewSyncMap()

	return auth
}
