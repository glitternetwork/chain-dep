package visitor

import (
	"fmt"
	"reflect"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/charset"
	driver "github.com/pingcap/tidb/types/parser_driver"
)

func ReplaceParams(stmt ast.StmtNode, params ...interface{}) (ast.StmtNode, error) {
	visitor := &replaceVisitor{params: params, paramMarkerIndex: -1}
	replacedStmt, _ := stmt.Accept(visitor)
	return replacedStmt.(ast.StmtNode), visitor.err
}

type replaceVisitor struct {
	params           []interface{}
	paramMarkerIndex int
	err              error
}

func (visitor *replaceVisitor) nextParamMarkerIndex() int {
	visitor.paramMarkerIndex++
	return visitor.paramMarkerIndex
}

func (visitor *replaceVisitor) Enter(n ast.Node) (ast.Node, bool) {
	if visitor.err != nil {
		return n, true
	}
	switch n := n.(type) {
	case *driver.ParamMarkerExpr:
		index := visitor.nextParamMarkerIndex()
		if index >= len(visitor.params) {
			visitor.err = fmt.Errorf("ParamMarkerExpr index out of bounds [%d,%d]", index, len(visitor.params))
			return n, true
		}

		value := visitor.params[index]
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n.Datum.SetInt64(rv.Int())
			return n, true
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			n.Datum.SetUint64(rv.Uint())
			return n, true
		case reflect.Float32, reflect.Float64:
			n.Datum.SetFloat64(rv.Float())
			return n, true
		case reflect.Bool:
			if rv.Bool() {
				n.Datum.SetInt64(1)
			} else {
				n.Datum.SetInt64(0)
			}
			return n, true
		case reflect.String:
			n.Datum.SetString(rv.String(), charset.CharsetUTF8MB4)
			return n, true
		case reflect.Slice:
			if rv.Elem().Type().Kind() == reflect.Uint8 {
				n.Datum.SetBytes(rv.Bytes())
				return n, true
			}
		}
		visitor.err = fmt.Errorf("invalid params type: %T on index %d", value, index)
		return n, true
	default:
		return n, false
	}
}

func (visitor *replaceVisitor) Leave(n ast.Node) (ast.Node, bool) {
	if visitor.err != nil {
		return n, false
	}
	switch n := n.(type) {
	case *driver.ParamMarkerExpr:
		return &n.ValueExpr, true
	default:
		return n, true
	}
}
