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
