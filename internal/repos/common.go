package repos

import (
	"database/sql"
)

func rowCount(res sql.Result) int64 {
	if res == nil {
		return 0
	}

	n, _ := res.RowsAffected()

	return n
}
