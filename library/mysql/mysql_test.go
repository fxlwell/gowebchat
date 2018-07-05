package mysql

import (
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
	err = db.GetRows("go_test", c, &ret)
	if err != nil || len(ret) != 10 {
		t.Fail()
	}

}

func Test_GetRowsIn(t *testing.T) {
	db, err := NewMysql("nice", "Cb84eZaa229ddnm", "test", "10.10.200.12:3306")
	if err != nil {
		panic(err)
	}

	c := NewSelectMap()
	in := []string{"20118", "20119", "20120"}
	c.SetCondIn("id", in)
	c.SetLimit(10)

	var ret []map[string]string
	err = db.GetRows("go_test", c, &ret)
	if err != nil || len(ret) != 3 {
		t.Fail()
	}
}

func Test_GetRowsNotIn(t *testing.T) {
	db, err := NewMysql("nice", "Cb84eZaa229ddnm", "test", "10.10.200.12:3306")
	if err != nil {
		panic(err)
	}

	c := NewSelectMap()
	notin := []string{"20118", "20119", "20120"}
	c.SetCondNotIn("id", notin)
	c.SetLimit(10)

	var ret []map[string]string
	err = db.GetRows("go_test", c, &ret)
	if err != nil || len(ret) != 10 {
		t.Fail()
	}
}

func Test_Like(t *testing.T) {
	db, err := NewMysql("nice", "Cb84eZaa229ddnm", "test", "10.10.200.12:3306")
	if err != nil {
		panic(err)
	}

	c := NewSelectMap()
	c.SetCondition("type LIKE", "%user%")
	c.SetLimit(10)

	var ret []map[string]string
	err = db.GetRows("go_test", c, &ret)
	if err != nil || len(ret) != 10 {
		t.Fail()
	}

}

func Test_Update(t *testing.T) {
	db, err := NewMysql("nice", "Cb84eZaa229ddnm", "test", "10.10.200.12:3306")
	if err != nil {
		panic(err)
	}

	c := NewSelectMap()
	c.SetLimit(10)
	c.SetFieldUp("value", "888888")
	c.SetFieldUp("type", "999999")
	c.SetCondition("type !=", "999999")

	var num int64
	num, err = db.Update("go_test", c)
	if err != nil || num != 10 {
		t.Fail()
	}

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
func Benchmark_getRows(b *testing.B) {
	db, err := NewMysql("nice", "Cb84eZaa229ddnm", "test", "10.10.200.12:3306")
	if err != nil {
		panic(err)
	}

	c := NewSelectMap()
	c.SetLimit(1)

	for i := 0; i < b.N; i++ { //use b.N for looping
		var ret []map[string]string
		db.GetRows("go_test", c, &ret)
	}
}
