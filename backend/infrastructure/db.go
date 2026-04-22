package infrastructure

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)

var DB *sql.DB

func InitDB(cfg DBConfig) {
	query := url.Values{}
	query.Set("database", cfg.Name)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic("cannot open database: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("cannot connect to database: " + err.Error())
	}

	DB = db
}
