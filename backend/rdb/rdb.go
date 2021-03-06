package rdb

import (
	"database/sql"
	"fmt"
	"os"

	// MySQL library, comment out to use PostgreSQL.
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER = "mysql"
)

var (
	user  = os.Getenv("CLOUDSQL_USER")
	pass  = os.Getenv("CLOUDSQL_PASSWORD")
	host  = os.Getenv("CLOUDSQL_HOST")
	db    = os.Getenv("CLOUDSQL_DB")
	dbURI = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, db)
)

type DBContext struct {
	Connection *sql.DB
}

func NewDBContext() (*DBContext, error) {
	conn, err := sql.Open(DB_DRIVER, dbURI)
	if err != nil {
		return nil, err
	}
	return &DBContext{Connection: conn}, nil
}
