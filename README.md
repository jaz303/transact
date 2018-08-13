# transact

Copied and adapted slightly from [David Muto's blog post](https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/).

Wraps a function in an SQL transaction. If the function panics or returns an error, the transaction is rolled back and the error/panic is propagated. The callback function receives a restricted interface view of `sql.Tx` that hides `Begin()`, `Commit()` and `Rollback()`, as these are the wrapper's responsibility.

## Usage

```go
if err := transact.Do(db, func(tx transact.Tx) error {
    if _, err := tx.Exec("..."); err != nil {
        return err
    }
    if _, err := tx.Exec("..."); err != nil {
        return err
    }
    return nil
}); err != nil {
    writeEmptyJSONResponse(w, http.StatusInternalServerError)
    return
}

```