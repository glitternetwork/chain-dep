package analyzer

import (
	"github.com/glitternetwork/chain-dep/parser/gast"
	"github.com/pingcap/tidb/parser/ast"
)

func isSelectFunctionStmt(stmt *ast.SelectStmt) bool {
	isSelectFunction :=
		stmt.From == nil &&
			stmt.Where == nil &&
			stmt.Fields != nil &&
			len(stmt.Fields.Fields) == 1
	if !isSelectFunction {
		return false
	}
	firstField := stmt.Fields.Fields[0]
	_, ok := firstField.Expr.(*ast.FuncCallExpr)
	return ok
}

func MustParseSelectFunction(n *gast.AST) *ast.FuncCallExpr {
	stmt := n.SQLStmt.(*ast.SelectStmt)
	firstField := stmt.Fields.Fields[0]
	return firstField.Expr.(*ast.FuncCallExpr)
}
