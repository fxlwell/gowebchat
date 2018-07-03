package mysql

import (
	"testing"
)

func Test_Keys(t *testing.T) {
	om := NewOrderMap()
	om.Set("id", 10009)
	om.Set("type", "user_test")
	om.Set("value", "1234567898")
	om.Set("type", "user_test111")
	om.Set("type", "user_test222")
	keys := om.Keys()
	values := om.Values()
	if keys[0] == "id" && keys[1] == "type" && keys[2] == "value" && values[0] == 10009 && values[1].(string) == "user_test222" && values[2].(string) == "1234567898" {
	} else {
		t.Fail()
	}
}

func Test_Condition(t *testing.T) {
	c := NewCondition()
	c.Set("id >", 10009)
	c.Set("type =", "user_test")
	v := c.ExecVal()
	if s, _ := c.Prepare(); s != "WHERE id > ? and type = ?" {
		t.Fail()
	}
	if v[0] != 10009 || v[1].(string) == "type" {
		t.Fail()
	}
}

func Test_Field(t *testing.T) {
	c := NewField()
	c.Set("id", 10009)
	c.Set("type", "user_test")
	v := c.ExecValSet()
	if s, _ := c.Prepare(); s != "id,type" {
		t.Fail()
	}
	s, err := c.PrepareSet()
	if err != nil || s != "SET id = ?,type = ? " {
		t.Fail()
	}
	if v[0] != 10009 || v[1].(string) == "type" {
		t.Fail()
	}

	c1 := NewField()
	s, err = c1.PrepareSet()
	if err == nil {
		t.Fail()
	}
}

func Test_LimitOffset(t *testing.T) {
	var c SqlExpr
	c = NewLimitOffset()
	s, err := c.Prepare()
	v := c.ExecVal()
	if err != nil || s != "" || len(v) != 0 {
		t.Fail()
	}

	c.Set("limit", 109)
	s, err = c.Prepare()
	v = c.ExecVal()
	if err != nil || s != "LIMIT ? " || len(v) != 1 {
		t.Fail()
	}

	c.Set("offset", 2)
	s, err = c.Prepare()
	v = c.ExecVal()
	if err != nil || s != "LIMIT ?,? " || len(v) != 2 {
		t.Fail()
	}

	c = NewLimitOffset()
	c.Set("offset", 2)
	s, err = c.Prepare()
	if err == nil {
		t.Fail()
	}
}

func Test_GroupByHaving(t *testing.T) {
	var c SqlExpr
	c = NewGroupByHaving()
	s, err := c.Prepare()
	v := c.ExecVal()
	if err != nil || s != "" || len(v) != 0 {
		t.Fail()
	}

	c.Set("groupby", "id desc")
	s, err = c.Prepare()
	v = c.ExecVal()
	if err != nil || s != "GROUP BY id desc " || len(v) != 0 {
		t.Fail()
	}

	c.Set("cnt >", 12)
	s, err = c.Prepare()
	v = c.ExecVal()
	if err != nil || s != "GROUP BY id desc HAVING cnt > ? " || len(v) != 1 {
		t.Fail()
	}

	c.Set("mame =", "boy")
	s, err = c.Prepare()
	v = c.ExecVal()
	if err != nil || s != "GROUP BY id desc HAVING cnt > ? and mame = ? " || len(v) != 2 {
		t.Fail()
	}

	c = NewGroupByHaving()
	c.Set("mame =", "boy")
	s, err = c.Prepare()
	v = c.ExecVal()
	if err != nil || s != "" || len(v) != 0 {
		t.Fail()
	}
}

func Test_SelectMap(t *testing.T) {
	c := NewSelectMap()
	sql, err := c.GetPrepareSql("go_test")
	v := c.ExecVal()
	if sql != "SELECT * FROM go_test" || err != nil || len(v) != 0 {
		t.Fail()
	}

	c.SetField("type as t", "value")
	sql, err = c.GetPrepareSql("go_test")
	v = c.ExecVal()
	if sql != "SELECT type as t,value FROM go_test" || err != nil || len(v) != 0 {
		t.Fail()
	}

	c.SetCondition("id>=", "20001")
	c.SetCondition("type=", "user")
	c.SetCondition("value=", "1234567")
	sql, err = c.GetPrepareSql("go_test")
	v = c.ExecVal()

	if sql != "SELECT type as t,value FROM go_test WHERE id>= ? and type= ? and value= ?" || err != nil || len(v) != 3 {
		t.Fail()
	}

	c.SetLimit(1000)
	sql, err = c.GetPrepareSql("go_test")
	v = c.ExecVal()
	if sql != "SELECT type as t,value FROM go_test WHERE id>= ? and type= ? and value= ? LIMIT ?" || err != nil || len(v) != 4 {
		t.Fail()
	}

	c.SetOffset(10)
	sql, err = c.GetPrepareSql("go_test")
	v = c.ExecVal()
	if sql != "SELECT type as t,value FROM go_test WHERE id>= ? and type= ? and value= ? LIMIT ?,?" || err != nil || len(v) != 5 || v[4].(int64) != 1000 {
		t.Fail()
	}

	c.SetOrderBy("id desc")
	sql, err = c.GetPrepareSql("go_test")
	v = c.ExecVal()
	if sql != "SELECT type as t,value FROM go_test WHERE id>= ? and type= ? and value= ? ORDER BY id desc LIMIT ?,?" || err != nil || len(v) != 5 || v[4].(int64) != 1000 {
		t.Fail()
	}

	c.SetGroupBy("cnt")
	sql, err = c.GetPrepareSql("go_test")
	v = c.ExecVal()
	if sql != "SELECT type as t,value FROM go_test WHERE id>= ? and type= ? and value= ? GROUP BY cnt ORDER BY id desc LIMIT ?,?" || err != nil || len(v) != 5 || v[4].(int64) != 1000 {
		t.Fail()
	}
	c.SetHaving("cnt>", 199)
	sql, err = c.GetPrepareSql("go_test")
	v = c.ExecVal()
	if sql != "SELECT type as t,value FROM go_test WHERE id>= ? and type= ? and value= ? GROUP BY cnt HAVING cnt> ? ORDER BY id desc LIMIT ?,?" || err != nil || len(v) != 6 || v[3].(int) != 199 {
		t.Fail()
	}
}
