package mapper

import (
	"database/sql"
	"github.com/kamilbiela/gochat/model"
)

type UserMapper struct {
	db *sql.DB
}

func NewUserMapper(db *sql.DB) *UserMapper {
	um := new(UserMapper)
	um.db = db
	return um
}

func (u *UserMapper) GetByUsername(username string) (*model.User, error) {
	user := model.User{}

	err := u.db.QueryRow("SELECT id, name, organization_id, password, salt FROM user WHERE name = ?", username).
		Scan(&user.Id, &user.Name, &user.OrganizationId, &user.Password, &user.Salt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
