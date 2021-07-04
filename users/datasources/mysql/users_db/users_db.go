package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	//mysql_users_username = "mysql_users_username"
	//mysql_users_password = "mysql_users_password"
	//mysql_users_host     = "mysql_users_host"
	//mysql_users_schema   = "mysql_users_schema"
	username = "root"
	password = "password"
	host     = "localhost:3306"
	schema   = "users_db"
)

var (
	Client *sql.DB

	//username = os.Getenv(mysql_users_username)
	//password = os.Getenv(mysql_users_password)
	//host     = os.Getenv(mysql_users_host)
	//schema   = os.Getenv(mysql_users_schema)
)

func init() {
	var err error

	dbName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	// Open up database connection
	Client, err = sql.Open("mysql", dbName)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	//check if the connection is working
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connected")
}
