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

// Connect efetua uma conexao com o banco de dados
func Connect(host, port, user, password, dbName string, ssl bool) error {
	if dbs[dbName] == nil {
		dsn := getStringConnection(host, port, user, password, dbName, ssl)
		//db, err := sqlx.Connect("postgres", dsn)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		if err != nil {
			if isMissingDatabase(err, dbName) {
				if err = createDatabase(host, port, user, password, dbName, ssl); err == nil {
					//db, err = sqlx.Connect("postgres", dsn)
					db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
					if err != nil {
						return err
					}
				}
			}
			if err != nil {
				log.Error(fmt.Sprintf("Nao foi possivel conectar ao banco de dados: msg %v dsn: %s", err, dsn))
			}
		}
		dbs[dbName] = db
	}
	return nil
}
