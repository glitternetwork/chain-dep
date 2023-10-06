package schema

type IndexDef struct {
	Name       string
	Type       IndexType
	Columns    []*ColumnDef
	PrimaryKey bool
	Parser     string
	Comment    string
}

type IndexType int

const (
	InvalidIndex = iota
	ValueIndex
	FullTextIndex
)
