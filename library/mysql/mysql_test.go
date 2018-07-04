package mysql

import (
	"fmt"
	"testing"
)

func Test_Insert(t *testing.T) {
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

func Test_GetRows(t *testing.T) {
	db, err := NewMysql("nice", "Cb84eZaa229ddnm", "test", "10.10.200.12:3306")
	if err != nil {
		panic(err)
	}

	c := NewSelectMap()
	c.SetLimit(10)

	var ret []map[string]string
	err = db.GetRows(NewTable("go_test", 4), c, &ret)
	fmt.Println(ret, err)

}

func Benchmark_insert(b *testing.B) {
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

	for i := 0; i < b.N; i++ { //use b.N for looping
		db.Insert("go_test", om)
	}
}
