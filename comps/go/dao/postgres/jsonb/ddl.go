package jsonb

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	log "github.com/sirupsen/logrus"
	"strings"
)

// createDatabase given all necessary parameters, return error if any
func createDatabase(host, port, user, password, dbName string, ssl bool) error {
	db, err := getDbConnection(host, port, user, password, dbName, ssl)
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

// createTableIfDoesntExists based on the error, check if it's needed to create a table for current tenantId
func (d *Database[T]) createTableIfDoesntExists(err error) bool {
	if isMissingTable(err) {
		log.Errorf("creating table... '%v'", d.TenantId)
		result := d.db.Exec(getDocTableDDL(d.TenantId))
		if result != nil {
			if result.Error == nil {
				return true
			}
			log.WithError(result.Error).Errorf("error while creating table '%s'", d.TenantId)
		}
	}
	return false
}

// isMissingTable function to return if the error is related to missing table
func isMissingTable(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "relation \"") &&
		strings.Contains(err.Error(), "\" does not exist")
}

// isMissingDatabase function to return if the error is related of missing database
func isMissingDatabase(err error, dbName string) bool {
	msg := fmt.Sprintf("database \"%s\" does not exist", dbName)
	return strings.Contains(err.Error(), msg)
}

// getDocTableDDL return the DDL script for DocTable
func getDocTableDDL(tableName string) string {
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
