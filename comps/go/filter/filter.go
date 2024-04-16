package filter

import (
	"github.com/linksoft-dev/single/comps/go/obj"
	"github.com/linksoft-dev/single/comps/go/str"
	"google.golang.org/protobuf/types/known/anypb"
	"strings"
)

func NewFilter(mainFilter string) (q Filter) {
	q.MainFilter = str.UpperNoSpaceNoAccent(mainFilter)
	q.Limit = 50
	return
}

// FullSearch is a special function to search by many fields at once using OR logical Operator
func (q *Filter) FullSearch(value any, fields ...string) *Filter {
	v := obj.ToString("", value)
	if v == "" {
		return q
	}
	if q.Orconditions == nil {
		q.Orconditions = []*Condition{}
	}
	for _, field := range fields {
		q.Orconditions = append(q.Orconditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: v, FilterOperator: "OR"})
	}
	return q
}

func (q *Filter) Eq(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Equals, Value: interfaceToString(value)})
	return q
}

func (q *Filter) Ne(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Not: true, Operator: Operator_Equals, Value: interfaceToString(value)})
	return q
}

//func (q *Filter) AddWhere(where string) *Filter {
//	q.where = append(q.where, where)
//	return q
//}

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
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_In, Value: obj.ToStringAsArray(value, ",", true)})
	return q
}

func (q *Filter) NotIn(field string, value ...*anypb.Any) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Not: true, Operator: Operator_In, Value: interfaceToString(value)})
	return q
}

func (q *Filter) Gt(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Gt, Value: interfaceToString(value)})
	return q
}

func (q *Filter) Gte(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Gte, Value: interfaceToString(value)})
	return q
}

func (q *Filter) Lt(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Lt, Value: interfaceToString(value)})
	return q
}

func (q *Filter) Lte(field string, value interface{}) *Filter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Lte, Value: interfaceToString(value)})
	return q
}

func interfaceToString(i interface{}) string {
	return obj.ToString("", i)
}

func (q *Filter) Select(field ...string) *Filter {
	q.SelectFields = append(q.SelectFields, field...)
	return q
}

func (q *Filter) From(tableName string) *Filter {
	//q.TableName = tableName
	return q
}

func (q *Filter) AddCondition(condition *Condition) *Filter {
	q.Conditions = append(q.Conditions, condition)
	return q
}

func (q *Filter) OrderByAsc(field string) *Filter {
	q.AddOrderBy(field, Direction_ASC)
	return q
}

func (q *Filter) OrderByDesc(field string) *Filter {
	q.AddOrderBy(field, Direction_DESC)
	return q
}
func (q *Filter) AddOrderBy(field string, direction Direction) *Filter {
	if q.OrderBy == nil {
		q.OrderBy = []*OrderBy{}
	}
	q.OrderBy = append(q.OrderBy, &OrderBy{FieldName: field, Direction: direction})
	return q
}
