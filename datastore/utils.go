package datastore

import "fmt"

// DeleteAll deletes all records from a table.
func DeleteAll(table string) (err error) {
	q := fmt.Sprintf("delete from %s", table)
	_, err = Db.ExecContext(ctx, q)
	if err != nil {
		return
	}
	return
}
