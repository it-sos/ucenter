package mysql

import (
	"context"
	"database/sql"
)

// https://github.com/go-xorm/xorm

type MySQL struct {
	Conn *sql.DB
}

func ConnectMySQL(dsn string) (*MySQL, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &MySQL{
		Conn: conn,
	}, nil
}

func config() {

}

var ErrNoRows = sql.ErrNoRows

func (db *MySQL) Insert() {
}

func (db *MySQL) MultiInsert() {
}

func (db *MySQL) Delete() {
}

func (db *MySQL) Update() {
}

func (db *MySQL) Query(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	rows, err := db.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if scannable, ok := dest.(Scannable); ok {
		return scannable.Scan(rows)
	}

	if !rows.Next() {
		return ErrNoRows
	}
	return rows.Scan(dest)
}

func (db *MySQL) QueryOne() {
}

func (db *MySQL) QSql() {
}

func (db *MySQL) QSqlOne() {
}

func (db *MySQL) Count() {
}

func (db *MySQL) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.Conn.ExecContext(ctx, query, args...)
}

func (db *MySQL) Transaction() {
}

func (db *MySQL) Commit() {
}

func (db *MySQL) Rollback() {
}

