package util

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		SendPanicIfError(errRollBack)
		panic(err)
	} else {
		errCommit := tx.Commit()
		SendPanicIfError(errCommit)
	}
}
