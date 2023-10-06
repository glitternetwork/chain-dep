package gast

import (
	"github.com/pingcap/tidb/parser/ast"
)

type AST struct {
	Type    StmtType
	SQLStmt ast.StmtNode
}

type StmtType uint8

const (
	Invalid StmtType = iota
	Select
	Insert
	Delete
	CreateTable
	DropTable
	ShowTables
	Update
	CreateDatabase
	DropDatabase
	ShowDatabases
	AlterTable
	ShowColumns
	ShowCreateTable
	SelectFunction
)
