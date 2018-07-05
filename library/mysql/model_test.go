package mysql

import (
	"testing"
)

func Test_GetRows(t *testing.T) {
	db, err := NewMysql("nice", "Cb84eZaa229ddnm", "test", "10.10.200.12:3306")
	if err != nil {
		panic(err)
	}

	om := NewOrderMap()
	//om.Set("id", 10009)
	om.Set("type", "user_test")
	om.Set("value", "1234567898")
	om.Set("type", "user_test111")
	om.Set("type", "user_test222")

	id, err := db.Insert("go_test", om)
	if err == nil && id > 0 {

	} else {
		t.Fail()
	}
}
