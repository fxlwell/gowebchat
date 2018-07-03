package mysql

type OrderMap struct {
	keys []string
	_map map[string]interface{}
}

func NewOrderMap() *OrderMap {
	om := &OrderMap{
		_map: make(map[string]interface{}),
	}
	return om
}

func (om *OrderMap) Set(key string, value interface{}) *OrderMap {
	if _, ok := om._map[key]; !ok {
		om.keys = append(om.keys, key)
	}
	om._map[key] = value
	return om
}

func (om *OrderMap) Keys() []string {
	return om.keys
}

func (om *OrderMap) Values() []interface{} {
	var value []interface{}
	for _, v := range om.keys {
		value = append(value, om._map[v])
	}
	return value
}
