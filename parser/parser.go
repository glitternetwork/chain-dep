package parser

import (
	"context"

	"github.com/glitternetwork/chain-dep/parser/analyzer"
	"github.com/glitternetwork/chain-dep/parser/gast"
	"github.com/glitternetwork/chain-dep/parser/mysql"
	"github.com/glitternetwork/chain-dep/parser/visitor"
	"github.com/pkg/errors"
)

var supportStmtType = map[gast.StmtType]bool{
	gast.Select:          true,
	gast.Insert:          true,
	gast.Delete:          true,
	gast.Update:          true,
	gast.CreateTable:     true,
	gast.CreateDatabase:  true,
	gast.ShowTables:      true,
	gast.ShowDatabases:   true,
	gast.AlterTable:      true,
	gast.DropTable:       true,
	gast.DropDatabase:    true,
	gast.ShowCreateTable: true,
}

func ParseOneSQL(ctx context.Context, sql string, params ...interface{}) (*gast.AST, error) {
	p := mysql.DefaultParser()
	stmts, _, err := p.ParseSQL(sql)

	if err != nil {
		return nil, err
	}
	if len(stmts) == 0 {
		return nil, errors.New("empty sql statment")
	}
	if len(stmts) > 1 {
		return nil, errors.New("too many sql statments")
	}
	if len(params) > 0 {
		replacedStmt, err := visitor.ReplaceParams(stmts[0], params...)
		if err != nil {
			return nil, err
		}
		stmts[0] = replacedStmt
	}
	stmtType, err := analyzer.GetStatementType(stmts[0])
	if err != nil {
		return nil, err
	}
	if _, ok := supportStmtType[stmtType]; !ok {
		return nil, errors.New("statement not support now")
	}
	return &gast.AST{
		Type:    stmtType,
		SQLStmt: stmts[0],
	}, nil
}
