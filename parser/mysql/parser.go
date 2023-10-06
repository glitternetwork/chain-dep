package mysql

import (
	"github.com/pingcap/tidb/parser"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

var (
	_DefaultParser *parser.Parser = parser.New()
)

func DefaultParser() *parser.Parser {
	return _DefaultParser
}
