package analyzer

import (
	"github.com/glitternetwork/chain-dep/parser/gast"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"
)

func GetStatementType(stmt ast.StmtNode) (gast.StmtType, error) {
	switch v := stmt.(type) {
	case *ast.SelectStmt:
		if isSelectFunctionStmt(v) {
			return gast.SelectFunction, nil
		}
		return gast.Select, nil
	case *ast.InsertStmt:
		return gast.Insert, nil
	case *ast.DeleteStmt:
		return gast.Delete, nil
	case *ast.UpdateStmt:
		return gast.Update, nil
	case *ast.CreateTableStmt:
		return gast.CreateTable, nil
	case *ast.DropTableStmt:
		return gast.DropTable, nil
	case *ast.CreateDatabaseStmt:
		return gast.CreateDatabase, nil
	case *ast.DropDatabaseStmt:
		return gast.DropDatabase, nil
	case *ast.ShowStmt:
		if v.Tp == ast.ShowTables {
			return gast.ShowTables, nil
		}
		if v.Tp == ast.ShowDatabases {
			return gast.ShowDatabases, nil
		}
		if v.Tp == ast.ShowColumns {
			return gast.ShowColumns, nil
		}
		if v.Tp == ast.ShowCreateTable {
			return gast.ShowCreateTable, nil
		}
		return gast.Invalid, errors.New("statement not support")
	case *ast.AlterTableStmt:
		return gast.AlterTable, nil
	default:
		return gast.Invalid, errors.New("statement not support")
	}
}
