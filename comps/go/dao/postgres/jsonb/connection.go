package jsonb

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbs map[string]*gorm.DB

func init() {
	dbs = map[string]*gorm.DB{}
}

// Connect establishes a connection to the database
func Connect(host, port, user, password, dbName string, ssl bool) error {
	if dbs[dbName] == nil {
		db, err := getDbConnection(host, port, user, password, dbName, ssl)
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			if isMissingDatabase(err, dbName) {
				if err = createDatabase(host, port, user, password, dbName, ssl); err == nil {
					db, err = getDbConnection(host, port, user, password, dbName, ssl)
					if err != nil {
						return err
					}
				}
			}
			if err != nil {
				log.WithError(err).Errorf("failed when try to connect to database")
			}
		}
		dbs[dbName] = db
	}
	return nil
}

func getDbConnection(host, port, user, password, dbName string, ssl bool) (*gorm.DB, error) {
	dsn := getStringConnection(host, port, user, password, dbName, ssl)
	//db, err := sqlx.Connect("postgres", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
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
