package jsonb

import (
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbs map[string]*gorm.DB

func init() {
	dbs = map[string]*gorm.DB{}
}

const (
	duplicatedFieldCode = "23505"
)

// Connect efetua uma conexao com o banco de dados
func Connect(host, port, user, password, dbName string, ssl bool) error {
	if dbs[dbName] == nil {
		dsn := getStringConnection(host, port, user, password, dbName, ssl)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			if isMissingDatabase(err, dbName) {
				if err = createDatabase(host, port, user, password, dbName, ssl); err == nil {
					db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
					if err == nil {
						//cria a tabela para a autenticacao, pois é a partir dela que outras tabelas serao criadas
						a := Database{}
						sql := a.GetOrgTable("org_auth")
						if res := db.Exec(sql); res != nil {
							err = res.Error
						}
					}
				}
			}
			if err != nil {
				err = errors.New(err.Error())
				log.Error(fmt.Sprintf("Nao foi possivel conectar ao banco de dados: msg %v dsn: %s", err, dsn))
			}
		}
		dbs[dbName] = db
	}
	return nil
}

type Database struct {
	TenantId string
	Query    *Query
	db       *gorm.DB
	tx       *gorm.DB // guardar conexao quando for uma transacao
}

// NewDataBase factory method para criar uma isntancia da struct Database
func NewDataBase(dbName, tenantId string) (*Database, error) {
	db := &Database{db: dbs[dbName], TenantId: tenantId}
	db.Query = &Query{db: db}
	return db, nil
}

// Insert inicia uma transacao
func (d *Database) StartTransaction() (err error) {
	d.tx = d.db.Begin()
	err = d.tx.Error
	return err
}

// Insert commita uma transacao,
// o rollback eh feito automatico caso alguma operacao tenah erro dentro do scopo do banco de dados
func (d *Database) CommitTransaction() (err error) {
	if d.tx == nil {
		err = errors.New("Tentando fazer commit em uma transacao mas a mesma nao foi iniciada")
		return
	}

	d.tx.Commit()
	err = d.tx.Error
	d.tx = nil
	return err
}

// RollbackTransaction rollback é realizado automático dentro do socopo das operacoes de banco,
// se houver um erro em operacoes de banco de dados, entao é feito um rollback caso uma transacao tenha sido iniciada,
// porém há situacoes que o rollback precisa ser chamado caso tenha erro em outra camadas fora das operacoes de banco
// de dados, geralmente usando funcao defer para ter a certeza do rollback em caso de erro
func (d *Database) RollbackTransaction() (err error) {
	if d.tx == nil {
		err = errors.New("Tentando fazer rollback em uma transacao mas a mesma nao foi iniciada")
		return
	}
	d.tx.Rollback()
	err = d.tx.Error
	d.tx = nil
	return err
}

// GetOrgTable estrutura para criar a tabela para a organizacao
func (d *Database) GetOrgTable(tableName string) string {
	table := sqlbuilder.NewCreateTableBuilder()
	table.CreateTable(tableName).IfNotExists()
	table.Define("id", "VARCHAR(36)", "NOT NULL", "primary key")
	table.Define("collection", "VARCHAR(30)", "NOT NULL")
	table.Define("deleted_at", "timestamp with time zone")
	table.Define("doc", "jsonb", "not null")
	sql := table.String()
	sql += fmt.Sprintf(";create index if not exists index_default_filter on %s (collection, "+
		"deleted_at) where (deleted_at IS NULL)", tableName)
	return sql
}

// getStringConnection retorna a string de conexao para o banco dado
func getStringConnection(host, port, user, password, dbName string, ssl bool) string {
	sslMode := "disable"
	if ssl {
		sslMode = "enable"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=%s connect_timeout=5", host, port,
		user, password, dbName, sslMode)
}

// createDatabase função para criar um database, utilize o banco padrão para conectar ao servidor de banco de dados
// e então executar o comando sql para criar outro banco
func createDatabase(host, port, user, password, dbName string, ssl bool) error {
	dsn := getStringConnection(host, port, user, password, "postgres", ssl)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\";", dbName))
	return nil
}

// getValidationError return the validation error translated to friendly message
func (d *Database) getValidationError(err error) error {
	if pqError, ok := err.(*pgconn.PgError); ok {
		switch pqError.Code {
		case duplicatedFieldCode:
			values := strings.Split(pqError.Detail, "=")
			return ViolationUniqKey{Field: values[0], Value: values[1]}
		}
	}
	return err
}

func (d *Database) ResetQuery() {
	d.Query = &Query{db: d}
}

type ViolationUniqKey struct {
	msg   string
	Field string
	Value string
}

func (v ViolationUniqKey) Error() string {
	return fmt.Sprintf("Já existe um registro com o valor %s para o campo %s ", v.Value, v.Field)
}
