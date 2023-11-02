package jsonb

import (
	"fmt"
	"github.com/linksoft-dev/single/comps/go/obj"
	"strings"
)

// Record é uma strutura para ir montando o registro baseado em maps
// insere no banco na mesma ordem(bson.D) que é adicionada na variável Record
type Record struct {
	record map[string]interface{}
}

func NewRecord(fields map[string]interface{}) *Record {
	return &Record{
		record: fields,
	}
}

func (r *Record) AddField(fieldName string, value interface{}) *Record {
	r.record[fieldName] = value
	return r
}

// AddRecord Add a nested Record(not list)
func (r *Record) AddRecord(fieldName string, value Record) {
	if value.record != nil {
		r.record[fieldName] = value
	}
}

func (r *Record) AddRecordList(fieldName string, value []Record) {
	var list []map[string]interface{}
	for _, record := range value {
		list = append(list, record.record)
	}
	r.record[fieldName] = list
}

func (r *Record) GetRecord() map[string]interface{} {
	return r.record
}

func (r *Record) GetSQLInsertNoId(tableName string) (sql string, params []interface{}) {
	sql, params = r.GetSQLInsert(tableName)
	sql = strings.ReplaceAll(sql, "RETURNING id", "")
	return sql, params
}

func (r *Record) GetSQLInsert(tableName string) (sql string, params []interface{}) {
	fields := ""
	values := ""
	paramsCount := 1
	for key, value := range r.record {
		fields += key + ","
		values += fmt.Sprintf("?,")
		paramsCount++
		params = append(params, value)
	}
	fields = strings.TrimSuffix(fields, ",")
	values = strings.TrimSuffix(values, ",")
	sql = fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s) RETURNING id", tableName, fields, values)
	return sql, params
}

func (r *Record) GetSQLUpdate(tableName string, q Query) (query string, params []interface{}) {
	fields := ""
	paramsCount := 1

	for key, value := range r.record {
		if obj.IsZeroValue(value) {
			fields += key + "=null,"
		} else {
			fields += key + "=" + fmt.Sprintf("?,")
			paramsCount++
			params = append(params, value)
		}
	}
	fields = strings.TrimSuffix(fields, ",")

	where, whereParams := q.getWhere(paramsCount)
	params = append(params, whereParams...)
	query = fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, fields, where)
	return query, params
}
