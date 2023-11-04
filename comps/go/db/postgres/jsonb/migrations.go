package jsonb

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	dbMigrationTableName = "migrations"
)

// MigrationModel mode used to manage migrationsSingleton
type MigrationModel struct {
	Id        int
	CreatedAt time.Time `Db:"created_at"`
	Module    string    `Db:"module"`
	Name      string    `Db:"name"`
	SQL       string
}

// TableName is used by the Gorm to define table name
func (m *MigrationModel) TableName() string {
	return dbMigrationTableName
}

// NewDbMigration factory method to create a new instance of MigrationModel
func NewMigration(dbName string) *MigrationModel {
	return &MigrationModel{}
}

// GetMigrationNumber get the current migration number
func (m *MigrationModel) GetMigrationNumber() int64 {
	if m.Name != "" {
		migrationNumber, err := strconv.ParseInt(m.Name[0:3], 10, 32)
		if err != nil {
			panic(err)
		}
		return migrationNumber
	}
	return 0
}

//
//// ApplyMigration aplique as migracoes caso as mesmas tenah sido setadas no objeto Database
//func (d *Database) ApplyMigration() error {
//	if d.migrations != nil {
//		for _, migration := range *d.migrations {
//			currentMigration := d.getCurrentMigration(migration.Module)
//			if migration.GetMigrationNumber() > currentMigration.GetMigrationNumber() {
//				d.ExecSQL(migration.SQL)
//				if _, err := d.InsertWithoutOrg(map[string]interface{}{
//					"created_at": time.Now(),
//					"module":     migration.Module,
//					"name":       migration.Name,
//				}, dbMigrationTableName); err != nil {
//					return err
//				}
//			}
//		}
//	}
//	return nil
//}

// getCurrentMigration return current/latest migration applied to the database
//func (d *Database) getCurrentMigration(module string) (currentMigration MigrationModel) {
//	//migrations := []MigrationModel{}
//	err := d.Query.From("migrations").Eq("module", module).FindRaw(&currentMigration)
//	if err != nil {
//		panic(err)
//	}
//	return
//}

// AutoMigrateOnError function used to identify the error and decide what to do, the goal of this function is update the
// database structure based on the error, so apply the getMigrations if needed
//func (d *Database) AutoMigrateOnError(err error) error {
//	if err != nil {
//		if d.isMissingTable(err) || d.isMissingColumn(err) {
//			return d.ApplyMigration()
//		}
//	}
//	return nil
//}

// isMissingDatabase function to return if the error is related to database
func isMissingDatabase(err error, dbName string) bool {
	msg := fmt.Sprintf("database \"%s\" does not exist", dbName)
	return strings.Contains(err.Error(), msg)
}

//
//// isMissingTable function to return if the error is related to missing table
//func (d *Database) isMissingTable(err error) bool {
//	if err == nil {
//		return false
//	}
//	return strings.Contains(err.Error(), "relation \"") &&
//		strings.Contains(err.Error(), "\" does not exist")
//}
//
//// createTable cria uma tabela baseada na mensagem de erro
//func (d *Database) createTable(e error) bool {
//	re, err := regexp.Compile(`ERROR: relation "(?P<org>.*?)"`)
//	if err != nil {
//		fmt.Print(err)
//	}
//	matches := re.FindAllStringSubmatch(e.Error(), 1)
//	for _, match := range matches {
//		d.ExecSQL(d.GetOrgTable(match[1]))
//		return true
//	}
//	return false
//}
//
//// isMissingColumn function to return if the error is related to missing column
//func (d *Database) isMissingColumn(err error) bool {
//	return strings.Contains(err.Error(), "column \"") &&
//		strings.Contains(err.Error(), "\" does not exist")
//}
//
//// isAuthenticationError function to return if the error is related Authentication for the database
//func (d *Database) isAuthenticationError(err error) bool {
//	return strings.Contains(err.Error(), "authentication failed")
//}
//
//// isDuplicatedFieldError function to return if the error is related duplicate Field
//func (d *Database) isDuplicatedFieldError(err error) bool {
//	return strings.Contains(err.Error(), "duplicate key Value violates unique constraint")
//}
