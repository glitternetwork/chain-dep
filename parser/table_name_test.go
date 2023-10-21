package parser

import (
	"context"
	"fmt"
	"github.com/glitternetwork/chain-dep/parser/analyzer"
	"testing"
)

func Test_ABC(t *testing.T) {
	sql := "select * from otherpay.profile"
	ctx := context.Background()
	ast, e1 := ParseOneSQL(ctx, sql)
	if e1 != nil {
		fmt.Println(e1)
		return
	}
	s, e2 := analyzer.FindDBTablID(ast)
	fmt.Println(e2)
	fmt.Println(s.Database)
	fmt.Println(s.Table)

}
