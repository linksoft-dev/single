package jsonb

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/linksoft-dev/single/comps/go/str"
	"gorm.io/gorm"
)

// Doc represent a record structured in the database
type Doc struct {
	Id         string
	Collection string    `Db:"collection"`
	DeletedAt  time.Time `Db:"deleted_at"`
	Doc        string
}

// Add é uma funcao para inserir um modelo de documento
func (d *Database) Create(collection string, obj interface{}) (err error) {

	b, err := json.Marshal(obj)
	if err != nil {
		return
	}

	id := str.ExtractValue(string(b), "id")
	if id == "" {
		return errors.New("id is empty")
	}
	fields := map[string]interface{}{
		"id":         id,
		"collection": collection,
		"doc":        string(b),
	}
	sqlQuery, params := NewRecord(fields).GetSQLInsert(d.TenantId)
	_, err = d.ExecSQL(sqlQuery, params...)

	if err != nil {
		if d.createTableIfDoesntExists(err) {
			return d.Create(collection, obj)
		}
	}
	return
}

type ObjI interface {
	GetId() string
}

type Obj struct {
	Id string
}

func (o *Obj) GetId() string {
	return o.Id
}

func (d *Database) Save(objs []ObjI) (err error) {
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
			return err2
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
			result, err2 := d.ExecSQL(sql.String())
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

// Change funcao especifica para alterar um modelo de documento
func (d *Database) Update(id, tableName string, obj interface{}) (affected int, err error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return
	}

	q := fmt.Sprintf("UPDATE %s set doc=? where id=? and collection=?", d.TenantId)
	res, err := d.ExecSQL(q, b, id, tableName)
	if err != nil {
		return
	}
	if err != nil {
		if d.createTableIfDoesntExists(err) {
			return d.Update(id, tableName, obj)
		}
	}
	affected = int(res.RowsAffected)
	return
}

func (d *Database) Delete(id string, table string) (affected int, err error) {
	q := fmt.Sprintf("UPDATE %s set deleted_at=? where id=? and collection=?", d.TenantId)
	res, err := d.ExecSQL(q, time.Now(), id, table)
	if err != nil {
		return
	}

	affected = int(res.RowsAffected)
	return
}

func (d *Database) DeleteRaw(id string, table string) (affected int, err error) {
	res, err := d.ExecSQL(fmt.Sprintf("DELETE FROM %s where id=?", table), id)
	if err != nil {
		return
	}
	affected = int(res.RowsAffected)
	return
}

// Find executa uma Query livremente, em obj,
// e pra passar uma estutura que seja compatível com o resultado esperado da Query
// Exemplo: se a Query retorna um array dois campos, resultado de uma agregacao, então
// passe como obj um array com os mesmos nomes de campos que é retornado na Query
func (d *Database) Find(query string, obj interface{}) (err error) {
	//err = d.Select(obj, Query)
	return
}

func (d *Database) FindById(id string, obj interface{}) (err error) {
	_ = d.Query.Eq("id", id).FindRaw(obj)
	return
}

// Select funcao para executar uma consulta(Query), o parametro `dest` é usado para gravar o resultado da Query
func (d *Database) Select(dest interface{}, query string, args ...interface{}) (err error) {
	d.ResetQuery()
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

func (d *Database) createTableIfDoesntExists(err error) bool {
	if d.isMissingTable(err) {
		log.Warnf("creating table... '%v'", err)
		_, err = d.ExecSQL(d.GetOrgTable(d.TenantId))
		return true
	}
	return false
}

// ExecSQL funcao para executar insert, update, delete e atualizar schema do banco, essa funcao NÃO FAZ QUERY
func (d *Database) ExecSQL(query string, args ...interface{}) (result *gorm.DB, err error) {
	d.ResetQuery()
	if d.tx != nil {
		result = d.tx.Exec(query, args...)
	} else {
		result = d.db.Exec(query, args...)
	}
	err = result.Error
	return
}
