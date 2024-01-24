package jsonb

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

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
