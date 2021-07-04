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
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	r := usersDB[u.Id]
	if r == nil {
		return errors.NewBadRequestError(fmt.Sprintf(" u %d not found", u.Id))
	}

	u.Id = r.Id
	u.FirstName = r.FirstName
	u.LastName = r.LastName
	u.Email = r.Email
	u.DateCreated = r.DateCreated

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
