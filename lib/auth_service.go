package lib

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/jmcvetta/randutil"
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

func (as *AuthService) GeneratePassword(password string) (string, string) {
	salt, err := randutil.AlphaString(32)

	if err != nil {
		log.Fatal(err)
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(salt+password), 10)

	if err != nil {
		log.Fatal(err)
	}

	return string(pass), salt
}

func (as *AuthService) IsPasswordEqualToHashedOne(salt string, password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(salt+password))

	return err == nil
}

func (as *AuthService) IsValid(username string, password string) bool {

	user, err := as.userMapper.GetByUsername(username)
	if err != nil {
		log.Fatal(err)
	}

	if user == nil {
		return false
	}

	if as.IsPasswordEqualToHashedOne(user.Salt, password, user.Password) {
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
