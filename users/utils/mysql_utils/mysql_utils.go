package mysql_utils

import (
	"strings"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
	"github.com/guiaramos/bookstore/users/utils/errors"
)

const (
	errorNoRow = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError("not record found on database")
		}
		return errors.NewInterServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case mysqlerr.ER_DUP_ENTRY:
		return errors.NewInterServerError("invalid data")
	}

	return errors.NewInterServerError("error processing request")
}
