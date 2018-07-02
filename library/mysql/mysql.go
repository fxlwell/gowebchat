package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	dsn string
	db  *sql.DB
}

func NewMysql(username, password, databases, addr string) (*Mysql, error) {
	m := &Mysql{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, addr, databases)
	m.dsn = dsn

	fmt.Println("dsn", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	m.db = db

	if err := m.db.Ping(); err != nil {
		return nil, err
	}

	return m, nil
}
