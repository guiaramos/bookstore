package users

import (
	"fmt"

	"github.com/guiaramos/bookstore/users/utils/errors"
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
	c := usersDB[u.Id]
	if c != nil {
		if c.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("u %d already exists", u.Id))
	}

	usersDB[u.Id] = u

	return nil
}
