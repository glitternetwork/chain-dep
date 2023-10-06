package schema

import "fmt"

type Table struct {
	TableID
	Engine     string       `json:"engine"`
	Columns    []*ColumnDef `json:"columns"`
	Indexes    []*IndexDef  `json:"indexes"`
	CreateStmt string       `json:"create_stmt"`
	Comment    string       `json:"comment"`
}

type TableID struct {
	Table    string `json:"table_name"`
	Database string `json:"db_name"`
}

func (id TableID) String() string{
	return fmt.Sprintf("%s.%s",id.Database,id.Table)
}

func NewTableID(db, table string) TableID {
	id := TableID{
		Table:    table,
		Database: db,
	}
	if len(id.Database) == 0 {
		id.Database = DefaultStoreDatabaseName
	}
	return id
}
