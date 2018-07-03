package mysql

import ()

type MysqlModel struct {
	node  string
	table string
}

func NewMysqlModel(node, table string) (*MysqlModel, error) {
	m := &MysqlModel{node, table}
	return m, nil
}
