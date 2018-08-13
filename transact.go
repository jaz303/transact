package transact

import (
    "context"
    "database/sql"
)

type Tx interface {
    Exec(query string, args ...interface{}) (sql.Result, error)
    ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
    Prepare(query string) (*sql.Stmt, error)
    PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
    QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type TxFn func(Tx) error

func Do(db *sql.DB, fn TxFn) (err error) {
    tx, err := db.Begin()
    if err != nil {
        return
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    err = fn(tx)
    return err
}
