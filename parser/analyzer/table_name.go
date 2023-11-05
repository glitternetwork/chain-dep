package analyzer

import (
	"github.com/glitternetwork/chain-dep/parser/gast"
	"github.com/glitternetwork/chain-dep/schema"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"
)

func FindDBTablID(n *gast.AST) (schema.TableID, error) {
	var refs *ast.Join
	switch stmt := n.SQLStmt.(type) {
	case *ast.SelectStmt:
		if stmt.From == nil {
			return schema.TableID{}, syntaxError(n)
		}
		refs = stmt.From.TableRefs
	case *ast.InsertStmt:
		if stmt.Table == nil {
			return schema.TableID{}, syntaxError(n)
		}
		refs = stmt.Table.TableRefs
	case *ast.DeleteStmt:
		if stmt.TableRefs == nil {
			return schema.TableID{}, syntaxError(n)
		}
		refs = stmt.TableRefs.TableRefs
	case *ast.UpdateStmt:
		if stmt.TableRefs == nil {
			return schema.TableID{}, syntaxError(n)
		}
		refs = stmt.TableRefs.TableRefs
	case *ast.CreateTableStmt:
		return schema.NewTableID(stmt.Table.Schema.O, stmt.Table.Name.O), nil
	default:
		return schema.TableID{}, errors.Errorf("not support find table from this statement: type=%d", n.Type)
	}
	return getTableName(refs)
}

func getTableName(refs *ast.Join) (schema.TableID, error) {
	tableSrc, isTableSrc := refs.Left.(*ast.TableSource)
	if !allOK(
		isTableSrc,
		!refs.ExplicitParens,
		!refs.NaturalJoin,
		!refs.StraightJoin,
		refs.Using == nil,
		refs.On == nil,
		refs.Tp == 0,
		refs.Right == nil,
	) {
		return schema.TableID{}, errors.New("not support join statment")
	}
	n, ok := tableSrc.Source.(*ast.TableName)
	if !ok {
		return schema.TableID{}, errors.New("not support select from a complex table name")
	}
	return schema.NewTableID(n.Schema.O, n.Name.O), nil
}

func syntaxError(n *gast.AST) error {
	return errors.Errorf("syntax error: %s", n.SQLStmt.Text())
}

func allOK(cond ...bool) bool {
	for _, ok := range cond {
		if !ok {
			return false
		}
	}
	return true
}
