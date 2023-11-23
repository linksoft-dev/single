package dao

import (
	"github.com/linksoft-dev/single/comps/go/obj"
	"strings"
	"time"
)

func NewQuery() (q Query) {
	q.Limit = 50
	q.Page = 1
	return
}

type Query struct {
	TableName          string
	result             []interface{}
	records            []map[string]interface{}
	Fields             []string
	Ids                []string
	sql                string
	Sort               map[string]string
	RawQuery           string
	rawParams          []interface{}
	Limit              int
	Page               int
	Last               int
	First              int
	Conditions         []Condition
	OrCondition        map[int][]Condition
	orgId              string
	whereCondition     []string
	fixedWhere         string
	CreatedAtGte       time.Time
	CreatedAtLte       time.Time
	IncludeSoftDeleted bool
	hasFullSearch      bool
}

type Condition struct {
	Field          string
	Operator       Operator
	Value          interface{}
	FilterOperator string
}

type Operator string

const (
	OperatorEquals    Operator = "="
	OperatorNotEquals Operator = "!="
	OperatorContains  Operator = "contains"
	OperatorStarts    Operator = "starts"
	OperatorIn        Operator = "in"
	OperatorNotIn     Operator = "notIn"
	OperatorGt        Operator = ">"
	OperatorGte       Operator = ">="
	OperatorLt        Operator = "<"
	OperatorLte       Operator = "<="
)

// FullSearch is a special function to search by many fields at once using OR logical Operator
func (q *Query) FullSearch(value any, fields ...string) *Query {
	if obj.ToString("", value) == "" {
		return q
	}
	if q.OrCondition == nil {
		q.OrCondition = map[int][]Condition{}
	}
	idx := len(q.OrCondition)
	for _, field := range fields {
		q.hasFullSearch = true
		q.OrCondition[idx] = append(q.OrCondition[idx], Condition{field, OperatorStarts, value, "OR"})
	}
	return q
}

func (q *Query) HasFullSearch() bool {
	return q.hasFullSearch
}

func (q *Query) Eq(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorEquals, value, ""})
	return q
}

func (q *Query) Ne(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorNotEquals, value, ""})
	return q
}

func (q *Query) AddWhere(where string) *Query {
	q.whereCondition = append(q.whereCondition, where)
	return q
}

func (q *Query) Contains(field, value string) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorContains, value, ""})
	return q
}

func (q *Query) StartsWith(field, value string) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorStarts, value, ""})
	return q
}

// StartsOrContain funcao para consultar comecando com ou contendo o valor, se houver * no valor,
// ele consulta por contém, caso contrário consulta por start
func (q *Query) StartsOrContain(field, value string) *Query {
	if strings.Contains(value, "*") {
		q.Contains(field, strings.ReplaceAll(value, "*", ""))
	} else {
		q.StartsWith(field, value)
	}
	return q
}

func (q *Query) In(field string, value ...string) *Query {
	if len(value) == 0 {
		return q
	}
	q.Conditions = append(q.Conditions, Condition{field, OperatorIn, value, ""})
	return q
}

func (q *Query) NotIn(field string, value ...interface{}) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorNotIn, value, ""})
	return q
}

func (q *Query) InString(field string, value ...string) *Query {
	if len(value) > 0 {
		q.Conditions = append(q.Conditions, Condition{field, OperatorIn, value, ""})
	}
	return q
}

func (q *Query) Gt(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorGt, value, ""})
	return q
}

func (q *Query) Gte(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorGte, value, ""})
	return q
}

func (q *Query) Lt(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorLt, value, ""})
	return q
}

func (q *Query) Lte(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, Condition{field, OperatorLte, value, ""})
	return q
}

func (q *Query) Select(field ...string) *Query {
	q.Fields = append(q.Fields, field...)
	return q
}

func (q *Query) From(tableName string) *Query {
	q.TableName = tableName
	return q
}

func (q *Query) OrderByAsc(field string) *Query {
	if q.Sort == nil {
		q.Sort = make(map[string]string)
	}
	q.Sort[field] = "asc"
	return q
}

func (q *Query) OrderByDesc(field string) *Query {
	if q.Sort == nil {
		q.Sort = make(map[string]string)
	}
	q.Sort[field] = "desc"
	return q
}

// Reset coloca os campos de pesquisa ao estado inicial
func (q *Query) Reset() {
	q.Sort = nil
	q.Limit = 50
	q.Conditions = nil
}
