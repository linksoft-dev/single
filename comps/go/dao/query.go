package dao

import (
	"github.com/linksoft-dev/single/comps/go/str"
	"strings"
	"time"
)

func NewQuery(mainFilter string) (q Query) {
	q.MainFilter = str.UpperNoSpaceNoAccent(mainFilter)
	q.Limit = 50
	q.Page = 1
	return
}

type Query struct {
	MainFilter     string
	TableName      string
	result         []interface{}
	records        []map[string]interface{}
	Fields         []string
	Ids            []string
	sql            string
	Sort           map[string]string
	RawQuery       string
	rawParams      []interface{}
	Limit          int
	Page           int
	Last           int
	First          int
	Conditions     []condition
	orgId          string
	whereCondition []string
	fixedWhere     string
	CreatedAtGte   time.Time
	CreatedAtLte   time.Time
}

type condition struct {
	Field    string
	Operator operator
	Value    interface{}
}

type operator string

const (
	OperatorEquals    operator = "="
	OperatorNotEquals operator = "!="
	OperatorContains  operator = "contains"
	OperatorStarts    operator = "starts"
	OperatorIn        operator = "in"
	OperatorNotIn     operator = "notIn"
	OperatorGt        operator = ">"
	OperatorGte       operator = ">="
	OperatorLt        operator = "<"
	OperatorLte       operator = "<="
)

func (q *Query) Eq(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorEquals, value})
	return q
}

func (q *Query) Ne(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorNotEquals, value})
	return q
}

func (q *Query) AddWhere(where string) *Query {
	q.whereCondition = append(q.whereCondition, where)
	return q
}

func (q *Query) Contains(field, value string) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorContains, value})
	return q
}

func (q *Query) StartsWith(field, value string) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorStarts, value})
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

func (q *Query) In(field string, value ...interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorIn, value})
	return q
}

func (q *Query) NotIn(field string, value ...interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorNotIn, value})
	return q
}

func (q *Query) InString(field string, value ...string) *Query {
	if len(value) > 0 {
		q.Conditions = append(q.Conditions, condition{field, OperatorIn, value})
	}
	return q
}

func (q *Query) Gt(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorGt, value})
	return q
}

func (q *Query) Gte(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorGte, value})
	return q
}

func (q *Query) Lt(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorLt, value})
	return q
}

func (q *Query) Lte(field string, value interface{}) *Query {
	q.Conditions = append(q.Conditions, condition{field, OperatorLte, value})
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
