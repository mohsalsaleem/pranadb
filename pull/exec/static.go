package exec

import (
	"github.com/squareup/pranadb/common"
)

// Empty executor.
var Empty = &StaticRows{rows: common.NewRows(nil, 0)}

type StaticRows struct {
	pullExecutorBase
	rows *common.Rows
}

var _ PullExecutor = &StaticRows{}

func NewSingleValueBigIntRow(val int64, colName string) *StaticRows {
	colTypes := []common.ColumnType{common.BigIntColumnType}
	rf := common.NewRowsFactory(colTypes)
	rows := rf.NewRows(1)
	rows.AppendInt64ToColumn(0, val)
	colNames := []string{colName}
	return &StaticRows{
		pullExecutorBase: pullExecutorBase{
			colTypes:       rows.ColumnTypes(),
			colNames:       colNames,
			simpleColNames: colNames,
		},
		rows: rows,
	}
}

func ShowTableRows(rows *common.Rows) *StaticRows {
	colTypes := []common.ColumnType{common.VarcharColumnType}
	rf := common.NewRowsFactory(colTypes)
	rfRows := rf.NewRows(rows.RowCount())
	for i := 0; i < rows.RowCount(); i++ {
		val := rows.GetRow(i)
		rfRows.AppendStringToColumn(0, val.GetString(3))
	}
	return &StaticRows{
		pullExecutorBase: pullExecutorBase{
			colTypes:       rfRows.ColumnTypes(),
			colNames:       []string{"tables"},
			simpleColNames: []string{"tables"},
		},
		rows: rfRows,
	}
}

func (s *StaticRows) GetRows(limit int) (rows *common.Rows, err error) {
	return s.rows, nil
}
