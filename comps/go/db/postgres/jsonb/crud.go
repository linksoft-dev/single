package jsonb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"
	db2 "github.com/linksoft-dev/single/comps/go/db"
	"github.com/linksoft-dev/single/comps/go/obj"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/jackc/pgconn"
)

const (
	duplicatedFieldCode = "23505"
)

// NewDataBase factory method para criar uma isntancia da struct Database
func NewDataBase[T db2.ObjI[T]](dbName, tenantId, resourceName string) (*Database[T], error) {
	db := &Database[T]{db: dbs[dbName], TenantId: tenantId, resourceName: resourceName}
	db.Query = &db2.Query{}
	return db, nil
}

type Doc struct {
	Id         string
	Collection string       `db:"collection"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	Doc        string
}

type Database[T db2.ObjI[T]] struct {
	TenantId     string
	resourceName string
	Query        *db2.Query
	updateField  db2.UpdateField
	//db           *sqlx.DB
	db *gorm.DB
	tx *gorm.DB // guardar conexao quando for uma transacao
}

func (d *Database[T]) Create(obj T) (T, error) {
	list, err := d.Save(obj)
	if err != nil {
		return obj, err
	}
	if len(list) > 0 {
		return list[0], nil
	}

	return obj, nil
}

func (d *Database[T]) Update(obj T, fields db2.UpdateField) error {
	d.updateField = fields
	_, err := d.Save(obj)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database[T]) Save(objs ...T) (list []T, err error) {
	if err = d.StartTransaction(); err != nil {
		return
	}
	var sql strings.Builder
	count := 0
	step := 1000
	length := len(objs)
	if length < step {
		step = length
	}
	for idx, record := range objs {
		count++
		doc, err2 := json.Marshal(record)
		if err2 != nil {
			return list, err2
		}
		docStr := string(doc)
		docStr = strings.ReplaceAll(docStr, "'", "''")
		if record.GetId() == "" {
			sql.WriteString(fmt.Sprintf(`insert into %s(id,collection,doc)values('%s','%s','%s');`, d.TenantId,
				record.GetId(), d.TenantId, docStr))
		} else {
			sql.WriteString(fmt.Sprintf(`update %s set doc='%s' where id='%s' and collection like '%s';`, d.TenantId,
				docStr, record.GetId(), d.TenantId))
		}

		// check if step was reached or if it is the last record
		if count == step || idx == length-1 {
			count = 0
			result := d.db.Exec(sql.String())
			if err2 != nil {
				return
			}
			if result.Error != nil {
				return
			}
			sql.Reset()
		}
	}
	err = d.CommitTransaction()
	return
}

func (d *Database[T]) Delete(obj T) error {
	query := fmt.Sprintf("UPDATE %s set deleted_at=? where id=? and collection=?", d.TenantId)
	result := d.db.Exec(query, time.Now(), obj.GetId(), d.resourceName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Database[T]) DeleteHard(obj T) error {
	query := fmt.Sprintf("DELETE FROM %s where id=? and collection=?", d.TenantId)
	result := d.db.Exec(query, obj.GetId(), d.resourceName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Database[T]) Find(filter db2.Query) (records []T, err error) {
	sqlSb := sqlbuilder.NewSelectBuilder()
	sqlSb.Select("*")
	sqlSb.From(d.TenantId)
	setWhere(sqlSb, filter, d.resourceName)
	setOrderBy(sqlSb, filter)
	setLimit(sqlSb, filter)
	sql, args := sqlSb.Build()

	// concat os docs em uma string para formar um array de dos em string, para então fazer o marshal para o destino
	var docs []Doc
	err = d.Select(&docs, sql, args)
	if err != nil {
		return nil, err
	}

	// do unmarshal of all Docs found, concat all docs into list of T to be unmarshal at once
	var sb strings.Builder
	sb.WriteString("[")
	for _, value := range docs {
		sb.WriteString(value.Doc)
		sb.WriteString(",")
	}
	str := sb.String()
	str = strings.TrimSuffix(str, ",")
	str += "]"
	err = json.Unmarshal([]byte(str), &records)
	return
}

func (d *Database[T]) Get(id string) (t T, err error) {
	filter := db2.Query{}
	filter.First = 1
	filter.Eq("id", id)
	records, err := d.Find(filter)
	if err != nil {
		return
	}
	if len(records) > 0 {
		return records[0], nil
	}
	return
}
func (d *Database[T]) Select(dest interface{}, query string, args ...interface{}) (err error) {
	var result *gorm.DB
	if d.tx != nil {
		result = d.tx.Raw(query, args...).Scan(dest)
	} else {
		result = d.db.Raw(query, args...).Scan(dest)
	}

	err = result.Error
	if d.createTableIfDoesntExists(err) {
		result = d.db.Raw(query, args...).Scan(dest)
		err = result.Error
	}
	return
}

// Insert inicia uma transacao
func (d *Database[T]) StartTransaction() (err error) {
	//d.tx = d.db.Begin()
	//err = d.tx.Error
	return err
}

// Insert commita uma transacao,
// o rollback eh feito automatico caso alguma operacao tenah erro dentro do scopo do banco de dados
func (d *Database[T]) CommitTransaction() (err error) {
	//if d.tx == nil {
	//	err = errors.New("Tentando fazer commit em uma transacao mas a mesma nao foi iniciada")
	//	return
	//}
	//
	//d.tx.Commit()
	//err = d.tx.Error
	//d.tx = nil
	return err
}

// RollbackTransaction rollback é realizado automático dentro do socopo das operacoes de banco,
// se houver um erro em operacoes de banco de dados, entao é feito um rollback caso uma transacao tenha sido iniciada,
// porém há situacoes que o rollback precisa ser chamado caso tenha erro em outra camadas fora das operacoes de banco
// de dados, geralmente usando funcao defer para ter a certeza do rollback em caso de erro
func (d *Database[T]) RollbackTransaction() (err error) {
	//if d.tx == nil {
	//	err = errors.New("Tentando fazer rollback em uma transacao mas a mesma nao foi iniciada")
	//	return
	//}
	//d.tx.Rollback()
	//err = d.tx.Error
	//d.tx = nil
	return err
}

// getValidationError return the validation error translated to friendly message
func (d *Database[T]) getValidationError(err error) error {
	if pqError, ok := err.(*pgconn.PgError); ok {
		switch pqError.Code {
		case duplicatedFieldCode:
			values := strings.Split(pqError.Detail, "=")
			return ViolationUniqKey{Field: values[0], Value: values[1]}
		}
	}
	return err
}

func setWhere(sb *sqlbuilder.SelectBuilder, filter db2.Query, resourceName string) {
	// se estiver consultado com rawquery, nao processe nada, apenas faça o scan para o `dest`
	if filter.RawQuery == "" {
		if resourceName == "" {
			//return nil, errors.New("Nome da tabela nao foi passado para a Query")
		}

		sb.Where("deleted_at is null")
		// select the collection
		if strings.Contains(resourceName, "%") {
			sb.Where(sb.Like("collection", fmt.Sprintf("'%s%%'", resourceName)))
		} else {
			sb.Where(sb.E("collection", resourceName))
		}

		// caso nao tenha passado o campo doc, adicione automaticamente
		for idx, value := range filter.Conditions {
			if !strings.Contains(value.Field, "doc") {
				filter.Conditions[idx].Field = fmt.Sprintf("doc ->> '%s'", value.Field)
			}

			for _, c := range filter.Conditions {

				// converts some special values to the right format
				switch val := c.Value.(type) {
				case time.Time:
					if val.Hour() == 0 && c.Operator == db2.OperatorLte {
						val = time.Date(val.Year(), val.Month(), val.Day(), 23, 59, 59, 0, time.Local)
					}
					c.Value = val.Format("2006-01-02T15:04:05")
				case bool:
					c.Value = fmt.Sprintf("%t", c.Value)
				case float64:
					c.Field = fmt.Sprintf("COALESCE((%s)::float, 0) %s", c.Field, c.Operator)
					c.Value = fmt.Sprintf("%f", c.Value)
				case int:
					c.Field = fmt.Sprintf("COALESCE((%s)::integer, 0) %s", c.Field, c.Operator)
					c.Value = fmt.Sprintf("%d", c.Value)
				}

				switch c.Operator {
				case db2.OperatorEquals:
					if c.Value == nil {
						sb.Where(sb.IsNull(c.Field))
					} else {
						sb.Where(fmt.Sprintf("%s = ?", c.Field))
					}
				case db2.OperatorNotEquals:
					if c.Value == nil {
						sb.Where(sb.IsNotNull(c.Field))
					} else {
						sb.Where(sb.NE(c.Field, c.Value))
					}
				case db2.OperatorStarts:
					sb.Where(fmt.Sprintf("%s ilike %s%%", c.Field, c.Value))
				case db2.OperatorContains:
					sb.Where(fmt.Sprintf("%s ilike %%%s%%", c.Field, c.Value))
				case db2.OperatorIn:
					stringArray := obj.ToStringArray(c.Value)
					sb.Where(sb.In(c.Field, stringArray))
					continue
				case db2.OperatorNotIn:
					stringArray := obj.ToStringArray(c.Value)
					sb.Where(sb.NotIn(c.Field, stringArray))
					continue
				case db2.OperatorGte:
					sb.Where(sb.GE(c.Field, c.Value))
				case db2.OperatorGt:
					sb.Where(sb.G(c.Field, c.Value))
				case db2.OperatorLte:
					sb.Where(sb.LE(c.Field, c.Value))
				case db2.OperatorLt:
					sb.Where(sb.L(c.Field, c.Value))
				}
			}

		}
	}
}

func setOrderBy(sb *sqlbuilder.SelectBuilder, filter db2.Query) {
	//Se o Sort n ta com o cast pro doc, o mesmo deve ser adicionado
	for i, v := range filter.Sort {
		if !strings.Contains(i, "doc") {
			delete(filter.Sort, i)
			filter.Sort[fmt.Sprintf("doc -> '%s'", i)] = v
		}
	}

	// sorting
	if filter.First > 0 || filter.Last > 0 {
		sb.OrderBy("doc->>'createdAt' desc").Limit(1)
	} else {
		for key, value := range filter.Sort {
			if value == "asc" {
				sb.OrderBy(key).Asc()
				continue
			}
			sb.OrderBy(key).Desc()
		}
	}
	// caso seja uma pesquisa padrao, ou seja, passou da condicao acima, então adicione a ordenacao pelo ultimo inserido
	// caso nao tenha nenhuma instrucao de sort
	if len(filter.Sort) == 0 {
		filter.OrderByDesc("doc -> 'createdAt'")
	}
}

func setLimit(sb *sqlbuilder.SelectBuilder, filter db2.Query) {
	if filter.Limit > 0 {
		sb.Limit(filter.Limit)
	}

	if filter.Page > 1 {
		sb.Offset((filter.Page - 1) * filter.Limit)
	}
}

type ViolationUniqKey struct {
	msg   string
	Field string
	Value string
}

func (v ViolationUniqKey) Error() string {
	return fmt.Sprintf("Já existe um registro com o valor %s para o campo %s ", v.Value, v.Field)
}
