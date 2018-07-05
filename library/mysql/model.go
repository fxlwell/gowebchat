package mysql

import (
	"fmt"
)

const (
	MYSQL_NODE_CLUSER_SLAVE  = "slave"
	MYSQL_NODE_CLUSER_MASTER = "master"
)

type MysqlModel struct {
	Node   string
	Table  string
	Cluser string
}

func NewMysqlModel(node, table string) *MysqlModel {
	mo := &MysqlModel{node, table, MYSQL_NODE_CLUSER_SLAVE}
	return mo
}

func (md *MysqlModel) GetRows(elem *SqlExpr, result *[]map[string]string) error {
	sn := fmt.Sprintf("%s-%s,", md.Node, md.Cluser)
	mysql := Node(sn)
	if mysql == nil {
		return fmt.Errorf("there is no mysql sn : %s", sn)
	}
	err := mysql.GetRows(md.Table, elem, result)
	return err
}
