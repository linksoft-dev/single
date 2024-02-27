package filter

import (
	"github.com/linksoft-dev/single/comps/go/obj"
	"github.com/linksoft-dev/single/comps/go/str"
	"google.golang.org/protobuf/types/known/anypb"
	"strings"
)

type GoFilter struct {
	Filter
	OrCondition   map[int][]Condition
	hasFullSearch bool
}

func NewFilter(mainFilter string) (q GoFilter) {
	q.MainFilter = str.UpperNoSpaceNoAccent(mainFilter)
	q.Limit = 50
	return
}

// FullSearch is a special function to search by many fields at once using OR logical Operator
func (q *GoFilter) FullSearch(value any, fields ...string) *GoFilter {
	v := obj.ToString("", value)
	if v == "" {
		return q
	}
	if q.OrCondition == nil {
		q.OrCondition = map[int][]Condition{}
	}
	idx := len(q.OrCondition)
	for _, field := range fields {
		q.hasFullSearch = true
		q.OrCondition[idx] = append(q.OrCondition[idx], Condition{FieldName: field, Operator: Operator_Starts, Value: v, FilterOperator: "OR"})
	}
	return q
}

func (q *GoFilter) HasFullSearch() bool {
	return q.hasFullSearch
}

func (q *GoFilter) Eq(field string, value interface{}) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Equals, Value: interfaceToString(value)})
	return q
}

func (q *GoFilter) Ne(field string, value interface{}) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Not: true, Operator: Operator_Equals, Value: interfaceToString(value)})
	return q
}

func (q *GoFilter) AddWhere(where string) *GoFilter {
	q.AdditionalConditions = append(q.AdditionalConditions, where)
	return q
}

func (q *GoFilter) Contains(field, value string) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Contains, Value: value})
	return q
}

func (q *GoFilter) StartsWith(field, value string) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Starts, Value: value})
	return q
}

// StartsOrContain funcao para consultar comecando com ou contendo o valor, se houver * no valor,
// ele consulta por contém, caso contrário consulta por start
func (q *GoFilter) StartsOrContain(field, value string) *GoFilter {
	if strings.Contains(value, "*") {
		q.Contains(field, strings.ReplaceAll(value, "*", ""))
	} else {
		q.StartsWith(field, value)
	}
	return q
}

func (q *GoFilter) In(field string, value ...interface{}) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_In, Value: obj.ToStringAsArray(value, ",", true)})
	return q
}

func (q *GoFilter) NotIn(field string, value ...*anypb.Any) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Not: true, Operator: Operator_In, Value: interfaceToString(value)})
	return q
}

func (q *GoFilter) Gt(field string, value interface{}) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Gt, Value: interfaceToString(value)})
	return q
}

func (q *GoFilter) Gte(field string, value interface{}) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Gte, Value: interfaceToString(value)})
	return q
}

func (q *GoFilter) Lt(field string, value interface{}) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Lt, Value: interfaceToString(value)})
	return q
}

func (q *GoFilter) Lte(field string, value interface{}) *GoFilter {
	q.Conditions = append(q.Conditions, &Condition{FieldName: field, Operator: Operator_Lte, Value: interfaceToString(value)})
	return q
}

func interfaceToString(i interface{}) string {
	return ""
}

func (q *GoFilter) Select(field ...string) *GoFilter {
	q.SelectFields = append(q.SelectFields, field...)
	return q
}

func (q *GoFilter) From(tableName string) *GoFilter {
	//q.TableName = tableName
	return q
}

func (q *GoFilter) AddCondition(condition Condition) *GoFilter {
	q.Conditions = append(q.Conditions, &condition)
	return q
}

func (q *GoFilter) OrderByAsc(field string) *GoFilter {
	q.AddOrderBy(field, Direction_ASC)
	return q
}

func (q *GoFilter) OrderByDesc(field string) *GoFilter {
	q.AddOrderBy(field, Direction_DESC)
	return q
}
func (q *GoFilter) AddOrderBy(field string, direction Direction) *GoFilter {
	if q.OrderBy == nil {
		q.OrderBy = []*OrderBy{}
	}
	q.OrderBy = append(q.OrderBy, &OrderBy{FieldName: field, Direction: direction})
	return q
}
