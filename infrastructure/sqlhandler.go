package infrastructure

import (
	"database/sql"
	"os"

	"github.com/daisuke0925m/sample-saas/interfaces/database"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "root:@tcp("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

func (handler *SqlHandler) Begin() (database.Tx, error) {
	tx, err := handler.Conn.Begin()
	if err != nil {
		return nil, err
	}
	newTx := new(Tx)
	newTx.Tx = tx
	return newTx, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}

type Tx struct {
	Tx *sql.Tx
}

func (t Tx) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := t.Tx.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (t Tx) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := t.Tx.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

func (t Tx) Begin() (database.Tx, error) {
	return t, nil
}

func (t Tx) RollBack() error {
	return t.Tx.Rollback()
}
