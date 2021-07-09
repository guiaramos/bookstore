package users

import (
	"github.com/guiaramos/bookstore/users/datasources/mysql/users_db"
	"github.com/guiaramos/bookstore/users/utils/date_utils"
	"github.com/guiaramos/bookstore/users/utils/errors"
	"github.com/guiaramos/bookstore/users/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRow       = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
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
	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
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

	r, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := r.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	u.Id = userId

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInterServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}
