package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var dbs map[string]*gorm.DB

func init() {
	dbs = map[string]*gorm.DB{}
}

// GetConnection establishes a connection to the database and return the connection instance
func GetConnection(host, port, user, password, dbName string, ssl bool) (*gorm.DB, error) {
	if dbs[dbName] == nil {
		db, err := getDbConnection(host, port, user, password, dbName, ssl)

		if err != nil {
			if isMissingDatabase(err, dbName) {
				if err = createDatabase(host, port, user, password, dbName, ssl); err == nil {
					db, err = getDbConnection(host, port, user, password, dbName, ssl)
					if err != nil {
						return nil, err
					}
				}
			}
			if err != nil {
				log.WithError(err).Errorf("failed when try to connect to database")
			}
		}
		dbs[dbName] = db
	}
	return dbs[dbName], nil
}

// createDatabase given all necessary parameters, return error if any
func createDatabase(host, port, user, password, dbName string, ssl bool) error {
	db, err := getDbConnection(host, port, user, password, "postgres", ssl)
	if err != nil {
		return err
	}
	tx := db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\";", dbName))
	if tx != nil {
		if tx.Error == nil {
			return nil
		}
		return tx.Error
	}
	return fmt.Errorf("failed when try to create database %s, no retunred error message", dbName)
}

// IsMissingTableError function to return if the error is related to missing table
func IsMissingTableError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "relation \"") &&
		strings.Contains(err.Error(), "\" does not exist")
}

// CreateTableIfDoesntExists based on the error, check if it's needed to create a table for current tenantId
func CreateTableIfDoesntExists(db *gorm.DB, tableName, ddlScript string) error {
	log.Warnf("creating table... '%v'", tableName)
	result := db.Exec(ddlScript)
	if result != nil && result.Error != nil {
		log.WithError(result.Error).Errorf("error while creating table '%s'", tableName)
		return result.Error
	}
	return nil
}

// isMissingDatabase function to return if the error is related of missing database
func isMissingDatabase(err error, dbName string) bool {
	msg := fmt.Sprintf("database \"%s\" does not exist", dbName)
	return strings.Contains(err.Error(), msg)
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
