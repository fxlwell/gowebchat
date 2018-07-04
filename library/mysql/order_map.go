package mysql

import (
	"fmt"
	"regexp"
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

type Table struct {
	Name     string
	FieldNum int
}

func NewTable(name string, flen int) *Table {
	t := &Table{
		name,
		flen,
	}
	return t
}

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

	sql := fmt.Sprintf("WHERE %s", strings.Join(conds, " and "))

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
	/*
		var fields []string
		for _, v := range keys {
			fields = append(fields, fmt.Sprintf("`%s`", v))
		}
	*/
	sql := fmt.Sprintf("%s", strings.Join(keys, ","))
	return sql, nil
}
func (f *Field) PrepareSet() (string, error) {
	keys := f.val.Keys()
	if len(keys) == 0 {
		return "", fmt.Errorf("have no set value")
	}
	var sets []string
	for _, v := range keys {
		sets = append(sets, fmt.Sprintf("%s = ?", v))
	}

	sql := fmt.Sprintf("SET %s ", strings.Join(sets, ","))

	return sql, nil
}

func (f *Field) ExecValSet() []interface{} {
	return f.val.Values()
}

func (f *Field) ExecVal() []interface{} {
	var r []interface{}
	return r
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
	var r []interface{}
	maps := l.val.Maps()
	if len(maps) == 0 {
		return r
	}
	v_offset, ok_offset := maps[MYSQL_SELECT_OFFSET]
	v_limit, ok_limit := maps[MYSQL_SELECT_LIMIT]
	if ok_limit && ok_offset {
		r = append(r, v_offset, v_limit)
	} else if ok_limit && !ok_offset {
		r = append(r, v_limit)
	}

	return r
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

type SelectMap struct {
	Field         *Field
	Conds         SqlExpr
	LimitOffset   SqlExpr
	GroupbyHaving SqlExpr
	Orderby       string
}

func NewSelectMap() *SelectMap {
	sm := &SelectMap{
		Field:         NewField(),
		Conds:         NewCondition(),
		LimitOffset:   NewLimitOffset(),
		GroupbyHaving: NewGroupByHaving(),
	}
	return sm
}
func (sm *SelectMap) SetField(fields ...string) *SelectMap {
	for _, v := range fields {
		sm.Field.Set(v, v)
	}
	return sm
}
func (sm *SelectMap) SetCondition(cond string, value interface{}) *SelectMap {
	sm.Conds.Set(cond, value)
	return sm
}
func (sm *SelectMap) SetLimit(limit int64) *SelectMap {
	sm.LimitOffset.Set(MYSQL_SELECT_LIMIT, limit)
	return sm
}
func (sm *SelectMap) SetOffset(offset int64) *SelectMap {
	sm.LimitOffset.Set(MYSQL_SELECT_OFFSET, offset)
	return sm
}
func (sm *SelectMap) SetOrderBy(orderby string) *SelectMap {
	sm.Orderby = orderby
	return sm
}
func (sm *SelectMap) SetGroupBy(groupby string) *SelectMap {
	sm.GroupbyHaving.Set(MYSQL_SELECT_GROUPBY, groupby)
	return sm
}
func (sm *SelectMap) SetHaving(cond string, value interface{}) *SelectMap {
	sm.GroupbyHaving.Set(cond, value)
	return sm
}

func (sm *SelectMap) GetPrepareSql(table string) (string, error) {
	var (
		err error
		s   string
	)
	s, err = sm.Field.Prepare()
	if err != nil {
		return "", err
	}
	s_field := s

	s, err = sm.Conds.Prepare()
	if err != nil {
		return "", err
	}
	s_cond := s

	s, err = sm.GroupbyHaving.Prepare()
	if err != nil {
		return "", err
	}
	s_grouphaving := s

	s, err = sm.LimitOffset.Prepare()
	if err != nil {
		return "", err
	}
	s_limitoffset := s

	s_orderby := ""
	if sm.Orderby != "" {
		s_orderby = fmt.Sprintf("ORDER BY %s", sm.Orderby)
	}

	re, _ := regexp.Compile(`\s{2,}`)
	sql := fmt.Sprintf("SELECT %s FROM %s %s %s %s %s", s_field, table, s_cond, s_grouphaving, s_orderby, s_limitoffset)
	sql = re.ReplaceAllString(sql, " ")
	sql = strings.TrimRight(sql, " ")
	return sql, nil
}
func (sm *SelectMap) ExecVal() []interface{} {
	var ret []interface{}
	for _, v := range sm.Field.ExecVal() {
		ret = append(ret, v)
	}
	for _, v := range sm.Conds.ExecVal() {
		ret = append(ret, v)
	}
	for _, v := range sm.GroupbyHaving.ExecVal() {
		ret = append(ret, v)
	}
	for _, v := range sm.LimitOffset.ExecVal() {
		ret = append(ret, v)
	}
	return ret
}
