# transact

Copied and adapted slightly from [David Muto's blog post](https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/).

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