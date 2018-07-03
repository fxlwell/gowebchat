package mysql

import (
	"fmt"
	"strings"
)

const (
	MYSQL_SELECT_FIELD   = "field"
	MYSQL_SELECT_CONDS   = "conds"
	MYSQL_SELECT_LIMIT   = "limit"
	MYSQL_SELECT_OFFSET  = "offset"
	MYSQL_SELECT_ORDERBY = "orderby"
	MYSQL_SELECT_GROUPBY = "groupby"
	MYSQL_SELECT_HAVING  = "having"
)

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
func (om *OrderMap) Maps() map[string]interface{} {
	return om._map
}

/* interface{} */
type SqlExpr interface {
	Set(key string, value interface{})
	Prepare() (string, error)
	ExecVal() []interface{}
}

/*condition*/
type Condition struct {
	val *OrderMap
}

func NewCondition() *Condition {
	cond := &Condition{
		val: NewOrderMap(),
	}
	return cond
}

func (c *Condition) Set(cond string, value interface{}) {
	c.val.Set(cond, value)
}
func (c *Condition) Prepare() (string, error) {
	keys := c.val.Keys()
	if len(keys) == 0 {
		return "", nil
	}
	var conds []string
	for _, v := range keys {
		conds = append(conds, fmt.Sprintf("%s ?", v))
	}

	sql := fmt.Sprintf("WHERE %s ", strings.Join(conds, " and "))

	return sql, nil
}
func (c *Condition) ExecVal() []interface{} {
	return c.val.Values()
}

/* field */
type Field struct {
	val *OrderMap
}

func NewField() *Field {
	f := &Field{
		val: NewOrderMap(),
	}
	return f
}
func (f *Field) Set(cond string, value interface{}) {
	f.val.Set(cond, value)
}
func (f *Field) Prepare() (string, error) {
	keys := f.val.Keys()
	if len(keys) == 0 {
		return "*", nil
	}
	sql := fmt.Sprintf(" %s ", strings.Join(keys, ","))
	return sql, nil
}
func (f *Field) PrepareSet() (string, error) {
	keys := f.val.Keys()
	if len(keys) == 0 {
		return "", fmt.Errorf("have no set value")
	}
	for k, v := range keys {
		keys[k] = fmt.Sprintf("%s = ?", v)
	}

	sql := fmt.Sprintf("SET %s ", strings.Join(keys, ","))

	return sql, nil
}

func (f *Field) ExecVal() []interface{} {
	return f.val.Values()
}

/* limit offse*/
type LimitOffset struct {
	val *OrderMap
}

func NewLimitOffset() *LimitOffset {
	l := &LimitOffset{
		val: NewOrderMap(),
	}
	return l
}
func (l *LimitOffset) Set(cond string, value interface{}) {
	l.val.Set(cond, value)
}
func (l *LimitOffset) Prepare() (string, error) {
	maps := l.val.Maps()
	if len(maps) == 0 {
		return "", nil
	}
	_, ok_offset := maps[MYSQL_SELECT_OFFSET]
	_, ok_limit := maps[MYSQL_SELECT_LIMIT]

	if ok_offset && !ok_limit {
		return "", fmt.Errorf("offset is set but limit not set")
	}

	if ok_offset {
		return "LIMIT ?,? ", nil
	}

	return "LIMIT ? ", nil
}
func (l *LimitOffset) ExecVal() []interface{} {
	return l.val.Values()
}

/* groupby having */
type GroupByHaving struct {
	groupby string
	val     *OrderMap
}

func NewGroupByHaving() *GroupByHaving {
	g := &GroupByHaving{
		groupby: "",
		val:     NewOrderMap(),
	}
	return g
}
func (g *GroupByHaving) Set(cond string, value interface{}) {
	if cond == MYSQL_SELECT_GROUPBY {
		g.groupby = value.(string)
	} else {
		g.val.Set(cond, value)
	}
}
func (g *GroupByHaving) Prepare() (string, error) {
	if g.groupby == "" {
		return "", nil
	}
	keys := g.val.Keys()
	var having []string
	for _, v := range keys {
		having = append(having, fmt.Sprintf("%s ?", v))
	}

	if len(keys) > 0 {
		return fmt.Sprintf("GROUP BY %s HAVING %s ", g.groupby, strings.Join(having, " and ")), nil
	}

	return fmt.Sprintf("GROUP BY %s ", g.groupby), nil
}
func (g *GroupByHaving) ExecVal() []interface{} {
	var r []interface{}
	if g.groupby == "" {
		return r
	}
	return g.val.Values()
}

/*
type SelectMap struct {
	allkeys *OrderMap
	field   *OrderMap
	conds   *OrderMap
	limit   int64
	offset  int64
	orderby string
	groupby string
	having  *OrderMap
}

const (
	MYSQL_SELECT_FIELD   = "field"
	MYSQL_SELECT_CONDS   = "conds"
	MYSQL_SELECT_LIMIT   = "limit"
	MYSQL_SELECT_OFFSET  = "offset"
	MYSQL_SELECT_ORDERBY = "orderby"
	MYSQL_SELECT_GROUPBY = "groupby"
	MYSQL_SELECT_HAVING  = "having"
)

func NewSelectMap() *SelectMap {
	sm := &SelectMap{
		allkeys: NewOrderMap(),
		field:   NewOrderMap(),
		conds:   NewOrderMap(),
		having:  NewOrderMap(),
	}
	return sm
}
func (sm *SelectMap) SetField(field string) *SelectMap {
	sm.field.Set(field, field)
	sm.allkeys.Set(MYSQL_SELECT_FIELD, MYSQL_SELECT_FIELD)
	return sm
}
func (sm *SelectMap) SetCondition(cond string, value interface{}) *SelectMap {
	sm.conds.Set(cond, value)
	sm.allkeys.Set(MYSQL_SELECT_CONDS, MYSQL_SELECT_CONDS)
	return sm
}
func (sm *SelectMap) SetLimit(limit int64) *SelectMap {
	sm.limit = limit
	sm.allkeys.Set(MYSQL_SELECT_LIMIT, MYSQL_SELECT_LIMIT)
	return sm
}
func (sm *SelectMap) SetOffset(offset int64) *SelectMap {
	sm.offset = offset
	sm.allkeys.Set(MYSQL_SELECT_OFFSET, MYSQL_SELECT_OFFSET)
	return sm
}
func (sm *SelectMap) SetOrderBy(orderby string) *SelectMap {
	sm.orderby = orderby
	sm.allkeys.Set(MYSQL_SELECT_ORDERBY, MYSQL_SELECT_ORDERBY)
	return sm
}
func (sm *SelectMap) SetGroupBy(groupby string) *SelectMap {
	sm.groupby = groupby
	sm.allkeys.Set(MYSQL_SELECT_GROUPBY, MYSQL_SELECT_GROUPBY)
	return sm
}
func (sm *SelectMap) SetHaving(cond string, value interface{}) *SelectMap {
	sm.having.Set(cond, value)
	sm.allkeys.Set(MYSQL_SELECT_HAVING, MYSQL_SELECT_HAVING)
	return sm
}

func (sm *SelectMap) GetPrepareSql(table string) string {
	key_maps := sm.allkeys.Maps()

	key := MYSQL_SELECT_FIELD
	var f []string
	var target_f string
	if _, ok := key_maps[key]; ok {
		f = sm.field.Keys()
	} else {
		f = []string{"*"}
	}
	target_f = strings.Join(f, ',')

	key := "conds"
	var c []string
	var target_c string
	if _, ok := key_maps[key]; ok {
		c = sm.conds.Keys()
		target_f = strings.Join(f, ',')
	}

	key := "groupby"
	var g string
	if _, ok := key_maps[key]; ok {
		g = sm.groupby
	}

	key := "offset"
	var off int64
	if _, ok := key_maps[key]; ok {
		off = sm.offset
	}

	key := "limit"
	var lim int64
	if _, ok := key_maps[key]; ok {
		lim = sm.limit
	}
	sql := fmt.Sprintf("SELECT %s FROM %s %s %s %s")
	return sql
}
func (sm *SelectMap) GetExecValueSql() []string {
}
*/
