package analyzer

import (
	"github.com/glitternetwork/chain-dep/parser/gast"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"
)

func FindEngineName(n *gast.AST) (string, error) {
	switch stmt := n.SQLStmt.(type) {
	case *ast.CreateTableStmt:
		return getTableEngine(stmt.Options), nil
	default:
		return "", errors.Errorf("not support find engine from this statement: type=%s", n.Type)
	}
}

func getTableEngine(options []*ast.TableOption) string {
	for _, co := range options {
		if co.Tp == ast.TableOptionEngine {
			return co.StrValue
		}
	}
	return ""
}
