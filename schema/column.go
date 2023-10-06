package schema

type ColumnDef struct {
	Name    ColumnName   `json:"column_name"`
	Type    ColumnType   `json:"column_type"`
	SQLDef  SQLColumnDef `json:"sql_def"`
	Comment string       `json:"comment"`
}

type SQLColumnDef struct {
	Type   byte `json:"type"`
	Flag   uint `json:"flag"`
	Length int  `json:"length"`
}

type ColumnName struct {
	Origin  string `json:"origin_name"`
	LowCase string `json:"low_case_name"`
}

type ColumnType uint8

const (
	InvalidColumn ColumnType = iota
	IntColumn
	UintColumn
	FloatColumn
	BoolColumn
	StringColumn
	BytesColumn
)

type ColumnValue interface{}

var Null = ColumnValue(nil)
