package infrastructure

import (
	"database/sql"
	"github.com/ArtemYeremeev/go-repository-pattern/interfaces"
)

type DBHandler struct {
	Conn *sql.DB
}

func (handler *DBHandler) Execute(st string) {
	handler.Conn.Exec(st)
}

func (handler *DBHandler) Query(st string) (interfaces.IRow, error) {
	rows, err := handler.Conn.Query(st)

	if err != nil {
		return new(DBRow), err
	}
	row := new(DBRow)
	row.Rows = rows

	return row, nil
}

type DBRow struct {
	Rows *sql.Rows
}

func (r DBRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r DBRow) Next() bool {
	return r.Rows.Next()
}
