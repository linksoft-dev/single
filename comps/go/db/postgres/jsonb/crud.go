package jsonb

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Doc represent a record structured in the database
type Doc struct {
	Id         string
	Collection string       `db:"collection"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	Doc        string
}

//
//// Add é uma funcao para inserir um modelo de documento
//func (d *Database) Create(collection string, obj interface{}) (err error) {
//
//	b, err := json.Marshal(obj)
//	if err != nil {
//		return
//	}
//
//	id := str.ExtractValue(string(b), "id")
//	if id == "" {
//		return errors.New("id is empty")
//	}
//	fields := map[string]interface{}{
//		"id":         id,
//		"collection": collection,
//		"doc":        string(b),
//	}
//	sqlQuery, params := NewRecord(fields).GetSQLInsert(d.TenantId)
//	_, err = d.ExecSQL(sqlQuery, params...)
//
//	if err != nil {
//		if d.createTableIfDoesntExists(err) {
//			return d.Create(collection, obj)
//		}
//	}
//	return
//}

type ObjI interface {
	GetId() string
}

//
//// Change funcao especifica para alterar um modelo de documento
//func (d *Database) Update(id, tableName string, obj interface{}) (affected int, err error) {
//	b, err := json.Marshal(obj)
//	if err != nil {
//		return
//	}
//
//	q := fmt.Sprintf("UPDATE %s set doc=? where id=? and collection=?", d.TenantId)
//	res, err := d.ExecSQL(q, b, id, tableName)
//	if err != nil {
//		return
//	}
//	if err != nil {
//		if d.createTableIfDoesntExists(err) {
//			return d.Update(id, tableName, obj)
//		}
//	}
//	affected = int(res.RowsAffected)
//	return
//}

//func (d *Database) Delete(id string, table string) (affected int, err error) {
//	q := fmt.Sprintf("UPDATE %s set deleted_at=? where id=? and collection=?", d.TenantId)
//	res, err := d.ExecSQL(q, time.Now(), id, table)
//	if err != nil {
//		return
//	}
//
//	affected = int(res.RowsAffected)
//	return
//}

func (d *Database[T]) DeleteRaw(id string, table string) (affected int, err error) {
	//	res, err := d.ExecSQL(fmt.Sprintf("DELETE FROM %s where id=?", table), id)
	//	if err != nil {
	//		return
	//	}
	//	affected = int(res.RowsAffected)
	return
}

// Find executa uma Query livremente, em obj,
// e pra passar uma estutura que seja compatível com o resultado esperado da Query
// Exemplo: se a Query retorna um array dois campos, resultado de uma agregacao, então
// passe como obj um array com os mesmos nomes de campos que é retornado na Query
//func (d *Database) Find(query string, obj interface{}) (err error) {
//	//err = d.Select(obj, Query)
//	return
//}

func (d *Database[T]) FindById(id string) (t T, err error) {
	//_ = d.Query.Eq("id", id).FindRaw(obj)
	return
}

func (d *Database[T]) createTableIfDoesntExists(err error) bool {
	//if d.isMissingTable(err) {
	//	log.Warnf("creating table... '%v'", err)
	//	_, err = d.ExecSQL(d.GetOrgTable(d.TenantId))
	//	return true
	//}
	return false
}

// ExecSQL funcao para executar insert, update, delete e atualizar schema do banco, essa funcao NÃO FAZ QUERY
//func (d *Database[T]) ExecSQL(query string, args ...interface{}) (result *gorm.DB, err error) {
//d.ResetQuery()
//if d.tx != nil {
//	result = d.tx.Exec(query, args...)
//} else {
//	result = d.db.Exec(query, args...)
//}
//err = result.Error
//return
//}
