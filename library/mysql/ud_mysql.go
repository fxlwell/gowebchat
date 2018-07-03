package mysql

import (
	"fmt"
)

type MysqlConf struct {
	Username string
	Password string
	Database string
	Addr     string
}

var (
	_mysqlnodes map[string]*Mysql = make(map[string]*Mysql)
)

func Init(confs map[string]MysqlConf) error {
	for k, v := range confs {
		if dbTmp, err := NewMysql(v.Username, v.Password, v.Database, v.Addr); err != nil {
			return fmt.Errorf("node:%s,%s", k, err)
		} else {
			_mysqlnodes[k] = dbTmp
		}
	}
	return nil
}

func Node(node string) *Mysql {
	if mysql, ok := _mysqlnodes[node]; ok {
		return mysql
	}
	return nil
}
