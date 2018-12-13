package SQLHelper

import (
	. "../Config"
	. "../G"
	. "database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func Connection() *DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		Username, Password, Address, Dbname, Options)
	conn, err := Open("mysql", dataSourceName)
	CheckErr(err)
	return conn
}
