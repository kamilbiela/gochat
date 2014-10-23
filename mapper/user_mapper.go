package mapper

import (
	"database/sql"
	"github.com/kamilbiela/gochat/model"
	"log"
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

func (u *UserMapper) Save(user *model.User) {
	stmt, err := u.db.Prepare("UPDATE user SET name=?, organization_id=? password=?, salt=? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.Name, user.OrganizationId, user.Password, user.Salt)

	if err != nil {
		log.Fatal(err)
	}
}
