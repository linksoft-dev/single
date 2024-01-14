package filter

import (
	"github.com/linksoft-dev/single/comps/go/str"
	"strings"
)

func NewQuery(mainFilter string) (q Filter) {
	q.MainFilter = str.UpperNoSpaceNoAccent(mainFilter)
	q.Limit = 50
	q.Page = 1
	return
}

//
//type Filter struct {
//	MainFilter         string
//	TableName          string
//	result             []interface{}
//	records            []map[string]interface{}
//	Fields             []string
//	Ids                []string
//	sql                string
//	Sort               map[string]string
//	RawQuery           string
//	rawParams          []interface{}
//	Limit              int
//	Page               int
//	Last               int
//	First              int
//	Conditions         []Condition
//	orgId              string
//	whereCondition     []string
//	fixedWhere         string
//	CreatedAtGte       time.Time
//	CreatedAtLte       time.Time
//	IncludeSoftDeleted bool
//}

//type Condition struct {
//	Field    string
//	Operator operator
//	Value    interface{}
//}

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

func (q *Filter) Eq(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, Condition{field, OperatorEquals, value})
	return q
}

func (q *Filter) Ne(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, Condition{field, OperatorNotEquals, value})
	return q
}

func (q *Filter) AddWhere(where string) *Filter {
	q.whereCondition = append(q.whereCondition, where)
	return q
}

func (q *Filter) Contains(field, value string) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Contains, Value: value})
	return q
}

func (q *Filter) StartsWith(field, value string) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: value})
	return q
}

// StartsOrContain funcao para consultar comecando com ou contendo o valor, se houver * no valor,
// ele consulta por contém, caso contrário consulta por start
func (q *Filter) StartsOrContain(field, value string) *Filter {
	if strings.Contains(value, "*") {
		q.Contains(field, strings.ReplaceAll(value, "*", ""))
	} else {
		q.StartsWith(field, value)
	}
	return q
}

func (q *Filter) In(field string, value ...interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_In, Value: value})
	return q
}

func (q *Filter) NotIn(field string, value ...interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Not: true, Operator: Operator_In, Value: value})
	return q
}

func (q *Filter) InString(field string, value ...string) *Filter {
	if len(value) > 0 {
		q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: value})
		q.Conditions = append(q.Conditions, Condition{field, OperatorIn, value})
	}
	return q
}

func (q *Filter) Gt(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: value})
	q.Conditions = append(q.Conditions, Condition{field, OperatorGt, value})
	return q
}

func (q *Filter) Gte(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: value})
	q.Conditions = append(q.Conditions, Condition{field, OperatorGte, value})
	return q
}

func (q *Filter) Lt(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: value})
	q.Conditions = append(q.Conditions, Condition{field, OperatorLt, value})
	return q
}

func (q *Filter) Lte(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: value})
	q.Conditions = append(q.Conditions, Condition{field, OperatorLte, value})
	return q
}

func (q *Filter) Select(field ...string) *Filter {
	q.Fields = append(q.Fields, field...)
	return q
}

func (q *Filter) From(tableName string) *Filter {
	q.TableName = tableName
	return q
}

func (q *Filter) OrderByAsc(field string) *Filter {
	if q.Sort == nil {
		q.Sort = make(map[string]string)
	}
	q.Sort[field] = "asc"
	return q
}

func (q *Filter) OrderByDesc(field string) *Filter {
	if q.Sort == nil {
		q.Sort = make(map[string]string)
	}
	q.Sort[field] = "desc"
	return q
}
