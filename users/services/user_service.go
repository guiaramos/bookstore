package services

import (
	"github.com/guiaramos/bookstore/users/domain/users"
	"github.com/guiaramos/bookstore/users/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	r := &users.User{Id: userId}
	if err := r.Get(); err != nil {
		return nil, err
	}

	return r, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}
