package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Mysql struct {
	dsn string
	db  *sql.DB
}

func NewMysql(username, password, databases, addr string) (*Mysql, error) {
	m := &Mysql{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, addr, databases)
	m.dsn = dsn

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

func (m *Mysql) Insert(table string, value *OrderMap) (int64, error) {
	var (
		filed []string
		noval []string
	)
	for _, v := range value.Keys() {
		filed = append(filed, fmt.Sprintf("`%s`", v))
		noval = append(noval, "?")
	}
	stmt, err := m.db.Prepare(fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", table, strings.Join(filed, ","), strings.Join(noval, ",")))
	defer stmt.Close()
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(value.Values()...)
	if err != nil {
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -2, err
	}
	return id, nil
}
