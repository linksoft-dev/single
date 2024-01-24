package jsonb

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"
	"github.com/linksoft-dev/single/comps/go/dao"
	"github.com/linksoft-dev/single/comps/go/db"
	"github.com/linksoft-dev/single/comps/go/filter"
	"github.com/linksoft-dev/single/comps/go/obj"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/jackc/pgconn"
)

const (
	duplicatedFieldNameCode = "23505"
)

var tracer trace.Tracer

func init() {
	tracer = otel.GetTracerProvider().Tracer("single/dao/postgres/jsonb")
}

// NewDataBase factory method para criar uma isntancia da struct Database
func NewDataBase[T dao.ObjI[T]](dbConnection *gorm.DB) *Database[T] {
	return &Database[T]{
		db: dbConnection,
	}
}

type Doc struct {
	Id         string
	Collection string       `db:"collection"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	Doc        string
}

type Database[T dao.ObjI[T]] struct {
	updateFieldName dao.UpdateField
	db              *gorm.DB
	tx              *gorm.DB // guardar conexao quando for uma transacao
}

type crudData struct {
	tableName      string
	collectionName string
}

// getTenantInfoFromContext this function returns the crud information needed to perform crud
// operations
func getTenantInfoFromContext(ctx context.Context) (r crudData) {
	r.tableName, _ = ctx.Value("tenantId").(string)
	r.collectionName, _ = ctx.Value("collectionName").(string)
	return
}

func (d *Database[T]) Create(ctx context.Context, obj T) (T, error) {
	list, err := d.Save(ctx, true, obj)
	if err != nil {
		return obj, err
	}
	if len(list) > 0 {
		return list[0], nil
	}

	return obj, nil
}

func (d *Database[T]) Update(ctx context.Context, obj T, FieldNames dao.UpdateField) error {
	d.updateFieldName = FieldNames
	_, err := d.Save(ctx, false, obj)
	if err != nil {
		return err
	}
	return nil
}

// Save objects whatever is insert or update, based on id the save method decide which operation is
func (d *Database[T]) Save(ctx context.Context, insert bool, objs ...T) (list []T, err error) {
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
	r := getTenantInfoFromContext(ctx)
	for idx, record := range objs {
		count++
		doc, err2 := json.Marshal(record)
		if err2 != nil {
			return list, err2
		}
		if record.GetId() == "" {
			err = fmt.Errorf("FieldName id cannot be blank")
			return nil, err
		}

		docStr := string(doc)
		docStr = strings.ReplaceAll(docStr, "'", "''")
		if insert {
			sb.WriteString(fmt.Sprintf(`INSERT INTO %s(id,collection,doc) VALUES ('%s','%s','%s');`,
				r.tableName,
				record.GetId(),
				r.collectionName,
				docStr),
			)
		} else {
			sb.WriteString(fmt.Sprintf(`UPDATE %s SET doc='%s' WHERE id='%s' AND collection = '%s';`,
				r.tableName,
				docStr,
				record.GetId(),
				r.collectionName,
			),
			)
		}

		// check if step was reached or if it is the last record
		if count == step || idx == length-1 {
			count = 0
			result := d.db.Exec(sb.String())
			err := createTableIfDoesntExists(result.Error, d.db, r.tableName, getDocTableDDL(r.tableName))
			if err != nil {
				return list, err
			}

			result = d.db.Exec(sb.String())
			if result != nil && result.Error != nil {
				return list, result.Error
			}

			sb.Reset()
		}
	}
	err = d.CommitTransaction()
	list = objs
	return
}

func (d *Database[T]) Delete(ctx context.Context, obj T) error {
	r := getTenantInfoFromContext(ctx)
	query := fmt.Sprintf("UPDATE %s set deleted_at=? where id=? and collection=?", r.tableName)
	result := d.db.Exec(query, time.Now(), obj.GetId(), r.collectionName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Database[T]) DeleteHard(ctx context.Context, obj T) error {
	r := getTenantInfoFromContext(ctx)
	query := fmt.Sprintf("DELETE FROM %s where id=? and collection=?", r.tableName)
	result := d.db.Exec(query, obj.GetId(), r.collectionName)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *Database[T]) List(ctx context.Context, f filter.Filter) (records []T, err error) {
	ctx, span := tracer.Start(ctx, "dao/postgres/jsonb/Find")
	defer span.End()
	r := getTenantInfoFromContext(ctx)

	// create sql statement
	sqlSb := sqlbuilder.NewSelectBuilder()
	sqlSb.Select("*")
	sqlSb.From(r.tableName)
	setWhere(sqlSb, f, r.collectionName)
	setOrderBy(sqlSb, f)
	setLimit(sqlSb, f)
	sqlStatement, args := sqlSb.Build()

	// get docs from database
	var docs []Doc
	err = d.Select(ctx, &docs, sqlStatement, args)
	if err != nil {
		return nil, err
	}

	ctx, spanParent := tracer.Start(ctx, "dao/postgres/jsonb/Find/unmarshalDocs")
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

func (d *Database[T]) Get(ctx context.Context, id string) (t T, err error) {
	f := filter.Filter{}
	f.First = 1
	f.Eq("id", id)
	records, err := d.List(ctx, f)
	if err != nil {
		return
	}
	if len(records) > 0 {
		return records[0], nil
	}
	return
}
func (d *Database[T]) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) (err error) {
	ctx, span := tracer.Start(ctx, "dao/postgres/jsonb/Select")
	defer span.End()
	var result *gorm.DB
	if d.tx != nil {
		result = d.tx.Raw(query, args...).Scan(dest)
	} else {
		result = d.db.Raw(query, args...).Scan(dest)
	}

	// if select statement has error, check if it's related to missing table, create if it's missing
	err = result.Error
	if err != nil {
		r := getTenantInfoFromContext(ctx)
		err = createTableIfDoesntExists(err, result, r.tableName, r.collectionName)
		if err != nil {
			result = d.db.Raw(query, args...).Scan(dest)
			err = result.Error
		}
		if err != nil {
			log.WithContext(ctx).WithError(err).Error("failed while execute query")
		}
	}

	return
}

func createTableIfDoesntExists(err error, conn *gorm.DB, tableName, ddlScript string) error {
	if db.IsMissingTableError(err) {
		return db.CreateTableIfDoesntExists(conn, tableName, ddlScript)
	}
	return nil
}

func (d *Database[T]) StartTransaction() (err error) {
	d.tx = d.db.Begin()
	err = d.tx.Error
	return err
}

func (d *Database[T]) CommitTransaction() (err error) {
	if d.tx == nil {
		err = errors.New("trying to commit a non started transaction")
		return
	}

	d.tx.Commit()
	err = d.tx.Error
	d.tx = nil
	return err
}

func (d *Database[T]) RollbackTransaction() (err error) {
	if d.tx == nil {
		err = errors.New("trying to commit a non started transaction")
		return
	}
	d.tx.Rollback()
	err = d.tx.Error
	d.tx = nil
	return err
}

// getValidationError return the validation error translated to friendly message
func (d *Database[T]) getValidationError(err error) error {
	if pqError, ok := err.(*pgconn.PgError); ok {
		switch pqError.Code {
		case duplicatedFieldNameCode:
			values := strings.Split(pqError.Detail, "=")
			return ViolationUniqKey{FieldName: values[0], Value: values[1]}
		}
	}
	return err
}

func setWhere(sb *sqlbuilder.SelectBuilder, f filter.Filter, resourceName string) {
	// se econnect 177.54.145.68:27019stiver consultado com rawquery, nao processe nada, apenas faça o scan para o `dest`
	if f.RawFilter == "" {
		if resourceName == "" {
			//return nil, errors.New("Nome da tabela nao foi passado para a Query")
		}

		// if not include soft deleted, it means need to add f to make sure bring records were
		// not psychically deleted
		if f.IgnoreSoftDelete == false {
			sb.Where("deleted_at is null")
		}

		// select the collection
		if strings.Contains(resourceName, "%") {
			sb.Where(sb.Like("collection", fmt.Sprintf("'%s%%'", resourceName)))
		} else {
			sb.Where(sb.E("collection", resourceName))
		}

		// caso nao tenha passado o campo doc, adicione automaticamente
		for idx, value := range f.Conditions {
			if !strings.Contains(value.FieldName, "doc") {
				f.Conditions[idx].FieldName = fmt.Sprintf("doc ->> '%s'", value.FieldName)
			}
			var value interface{}
			for _, c := range f.Conditions {
				value = c.Value
				// converts some special values to the right format
				switch val := value.(type) {
				case time.Time:
					if val.Hour() == 0 && c.Operator == filter.Operator_Lte {
						val = time.Date(val.Year(), val.Month(), val.Day(), 23, 59, 59, 0, time.Local)
					}
					c.Value = val.Format("2006-01-02T15:04:05")
				case bool:
					c.Value = fmt.Sprintf("%t", c.Value)
				case float64:
					c.FieldName = fmt.Sprintf("COALESCE((%s)::float, 0) %s", c.FieldName, c.Operator)
					c.Value = fmt.Sprintf("%f", c.Value)
				case int:
					c.FieldName = fmt.Sprintf("COALESCE((%s)::integer, 0) %s", c.FieldName, c.Operator)
					c.Value = fmt.Sprintf("%d", c.Value)
				}

				switch c.Operator {
				case filter.Operator_Equals:
					if c.Not {
						if c.Value == "" {
							sb.Where(sb.IsNotNull(c.FieldName))
						} else {
							sb.Where(sb.NE(c.FieldName, c.Value))
						}
						continue
					}

					if c.Value == "" {
						sb.Where(sb.IsNull(c.FieldName))
					} else {
						sb.Where(fmt.Sprintf("%s = ?", c.FieldName))
					}
				case filter.Operator_Starts:
					sb.Where(fmt.Sprintf("%s ilike %s%%", c.FieldName, c.Value))
				case filter.Operator_Contains:
					sb.Where(fmt.Sprintf("%s ilike %%%s%%", c.FieldName, c.Value))
				case filter.Operator_In:
					stringArray := obj.ToStringArray(c.Value)
					if c.Not {
						sb.Where(sb.NotIn(c.FieldName, stringArray))
						continue
					}
					sb.Where(sb.In(c.FieldName, stringArray))
				case filter.Operator_Gte:
					sb.Where(sb.GE(c.FieldName, c.Value))
				case filter.Operator_Gt:
					sb.Where(sb.G(c.FieldName, c.Value))
				case filter.Operator_Lte:
					sb.Where(sb.LE(c.FieldName, c.Value))
				case filter.Operator_Lt:
					sb.Where(sb.L(c.FieldName, c.Value))
				}
			}

		}
	}
}

func setOrderBy(sb *sqlbuilder.SelectBuilder, f filter.Filter) {

	//if the order by has no cast to doc field, it means it has to be added
	for _, v := range f.OrderBy {
		if v == nil {
			continue
		}
		if !strings.Contains(v.FieldName, "doc") {
			v.FieldName = fmt.Sprintf("doc -> '%s'", v.FieldName)
		}
	}

	// sorting
	if f.First > 0 || f.Last > 0 {
		sb.OrderBy("doc->>'createdAt' desc").Limit(1)
	} else {
		for _, order := range f.OrderBy {
			if order.Direction == filter.Direction_ASC {
				sb.OrderBy(order.FieldName).Asc()
				continue
			}
			sb.OrderBy(order.FieldName).Desc()
		}
	}
	// if it is a standard search, that is, it passed the above condition, then add the ordering by last entered
	// if there is no sort instruction
	if len(f.OrderBy) == 0 {
		f.OrderByDesc("doc -> 'createdAt'")
	}
}

func setLimit(sb *sqlbuilder.SelectBuilder, f filter.Filter) {
	if f.Limit > 0 {
		sb.Limit(int(f.Limit))
	}

	//if f.Page > 1 {
	//	sb.Offset((f.Page - 1) * f.Limit)
	//}
}

type ViolationUniqKey struct {
	msg       string
	FieldName string
	Value     string
}

func (v ViolationUniqKey) Error() string {
	return fmt.Sprintf("Já existe um registro com o valor %s para o campo %s ", v.Value, v.FieldName)
}
