package jsonb

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"
	"github.com/linksoft-dev/single/comps/go/dao"
	"github.com/linksoft-dev/single/comps/go/obj"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/jackc/pgconn"
)

const (
	duplicatedFieldCode = "23505"
)

// NewDataBase factory method para criar uma isntancia da struct Database
func NewDataBase[T dao.ObjI[T]](dbName, tenantId, resourceName string) (*Database[T], error) {
	db := &Database[T]{db: dbs[dbName], TenantId: tenantId, resourceName: resourceName}
	return db, nil
}

type Doc struct {
	Id         string
	Collection string       `db:"collection"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	Doc        string
}

type Database[T dao.ObjI[T]] struct {
	TenantId     string
	resourceName string
	updateField  dao.UpdateField
	//db           *sqlx.DB
	db *gorm.DB
	tx *gorm.DB // guardar conexao quando for uma transacao
}

func (d *Database[T]) Create(obj T) (T, error) {
	list, err := d.Save(true, obj)
	if err != nil {
		return obj, err
	}
	if len(list) > 0 {
		return list[0], nil
	}

	return obj, nil
}

func (d *Database[T]) Update(obj T, fields dao.UpdateField) error {
	d.updateField = fields
	_, err := d.Save(false, obj)
	if err != nil {
		return err
	}
	return nil
}

// Save objects whatever is insert or update, based on id the save method decide which operation is
func (d *Database[T]) Save(insert bool, objs ...T) (list []T, err error) {
	if err = d.StartTransaction(); err != nil {
		return
	}
	var sb strings.Builder
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
		if record.GetId() == "" {
			insert = true
		}
		docStr := string(doc)
		docStr = strings.ReplaceAll(docStr, "'", "''")
		if insert {
			sb.WriteString(fmt.Sprintf(`INSERT INTO %s(id,collection,doc) VALUES ('%s','%s','%s');`,
				d.TenantId,
				record.GetId(),
				d.resourceName,
				docStr),
			)
		} else {
			sb.WriteString(fmt.Sprintf(`UPDATE %s SET doc='%s' WHERE id='%s' AND collection = '%s';`,
				d.TenantId,
				docStr,
				record.GetId(),
				d.resourceName),
			)
		}

		// check if step was reached or if it is the last record
		if count == step || idx == length-1 {
			count = 0
			result := d.db.Exec(sb.String())
			if err2 != nil {
				return
			}
			if result.Error != nil {
				return
			}
			sb.Reset()
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

func (d *Database[T]) Find(filter dao.Query) (records []T, err error) {
	ctx, span := trace.StartSpan(context.Background(), "dao/postgres/jsonb/Find")
	defer span.End()

	// create sql statement
	sqlSb := sqlbuilder.NewSelectBuilder()
	sqlSb.Select("*")
	sqlSb.From(d.TenantId)
	setWhere(sqlSb, filter, d.resourceName)
	setOrderBy(sqlSb, filter)
	setLimit(sqlSb, filter)
	sqlStatement, args := sqlSb.Build()

	// get docs from database
	var docs []Doc
	err = d.Select(&docs, sqlStatement, args)
	if err != nil {
		return nil, err
	}

	ctx, spanParent := trace.StartSpanWithRemoteParent(ctx, "dao/postgres/jsonb/Find/unmarshalDocs", span.SpanContext())
	err = unmarshalDocs(docs, &records)
	spanParent.End()
	return
}

// unmarshalDocs given the docs, return list of T records, this function concat all docs into list of strings
// then perform unmarshal at once to a list of T structs
func unmarshalDocs[T any](docs []Doc, records *[]T) error {
	var sb strings.Builder
	sb.WriteString("[")
	for _, value := range docs {
		sb.WriteString(value.Doc)
		sb.WriteString(",")
	}
	str := sb.String()
	str = strings.TrimSuffix(str, ",")
	str += "]"
	return json.Unmarshal([]byte(str), records)
}

func (d *Database[T]) Get(id string) (t T, err error) {
	filter := dao.Query{}
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
	ctx, span := trace.StartSpan(context.Background(), "dao/postgres/jsonb/Select")
	defer span.End()
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
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed while execute query")
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

func setWhere(sb *sqlbuilder.SelectBuilder, filter dao.Query, resourceName string) {
	// se estiver consultado com rawquery, nao processe nada, apenas faça o scan para o `dest`
	if filter.RawQuery == "" {
		if resourceName == "" {
			//return nil, errors.New("Nome da tabela nao foi passado para a Query")
		}

		// if not include soft deleted, it means need to add filter to make sure bring records were
		// not psychically deleted
		if filter.IncludeSoftDeleted == false {
			sb.Where("deleted_at is null")
		}

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
					if val.Hour() == 0 && c.Operator == dao.OperatorLte {
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
				case dao.OperatorEquals:
					if c.Value == nil {
						sb.Where(sb.IsNull(c.Field))
					} else {
						sb.Where(fmt.Sprintf("%s = ?", c.Field))
					}
				case dao.OperatorNotEquals:
					if c.Value == nil {
						sb.Where(sb.IsNotNull(c.Field))
					} else {
						sb.Where(sb.NE(c.Field, c.Value))
					}
				case dao.OperatorStarts:
					sb.Where(fmt.Sprintf("%s ilike %s%%", c.Field, c.Value))
				case dao.OperatorContains:
					sb.Where(fmt.Sprintf("%s ilike %%%s%%", c.Field, c.Value))
				case dao.OperatorIn:
					stringArray := obj.ToStringArray(c.Value)
					sb.Where(sb.In(c.Field, stringArray))
					continue
				case dao.OperatorNotIn:
					stringArray := obj.ToStringArray(c.Value)
					sb.Where(sb.NotIn(c.Field, stringArray))
					continue
				case dao.OperatorGte:
					sb.Where(sb.GE(c.Field, c.Value))
				case dao.OperatorGt:
					sb.Where(sb.G(c.Field, c.Value))
				case dao.OperatorLte:
					sb.Where(sb.LE(c.Field, c.Value))
				case dao.OperatorLt:
					sb.Where(sb.L(c.Field, c.Value))
				}
			}

		}
	}
}

func setOrderBy(sb *sqlbuilder.SelectBuilder, filter dao.Query) {
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

func setLimit(sb *sqlbuilder.SelectBuilder, filter dao.Query) {
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
