package model

type User struct {
	Id             int
	OrganizationId int
	Name           string
	Salt           string
	Password       string
}
