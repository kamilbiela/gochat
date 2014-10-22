package lib

import (
	"github.com/kamilbiela/gochat/mapper"
	dry "github.com/ungerik/go-dry"
	"log"
)

type AuthService struct {
	tokens     *dry.SyncMap
	userMapper *mapper.UserMapper
}

func NewAuthService(userMapper *mapper.UserMapper) *AuthService {
	auth := new(AuthService)
	auth.tokens = dry.NewSyncMap()
	auth.userMapper = userMapper

	return auth
}

func (as *AuthService) IsValid(username string, password string) bool {

	user, err := as.userMapper.GetByUsername(username)
	if err != nil {
		log.Fatal(err)
	}

	if user == nil {
		return false
	}

	if user.Name == username {
		return true
	}

	return false
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
