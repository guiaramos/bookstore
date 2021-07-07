package users

import (
	"fmt"
	"strings"

	"github.com/guiaramos/bookstore/users/datasources/mysql/users_db"
	"github.com/guiaramos/bookstore/users/utils/date_utils"
	"github.com/guiaramos/bookstore/users/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRow       = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInterServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
		}
		return errors.NewInterServerError(fmt.Sprintf("error when trying to get user %d: %s", u.Id, err.Error()))
	}

	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInterServerError(err.Error())
	}
	defer stmt.Close()

	u.DateCreated = date_utils.GetNowString()

	r, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", u.Email))
		}
		return errors.NewInterServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := r.LastInsertId()
	if err != nil {
		return errors.NewInterServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	u.Id = userId

	return nil
}
