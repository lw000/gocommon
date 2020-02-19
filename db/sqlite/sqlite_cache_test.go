package tysqlite

import (
	"database/sql"
	"testing"
)

func TestSQLite_Open(t *testing.T) {
	slte := NewSQLliteCache()
	if err := slte.Open("db.db"); err != nil {
		t.Errorf("SQLite.Open() error = %v", err)
	}
	defer func() {
		err := slte.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	{
		var sqlTablle = `
					CREATE TABLE IF NOT EXISTS "userInfo"(
						"username" VARCHAR(64) NULL,
						"departname" VARCHAR(64) NULL,
						"created" TIMESTAMP DEFAULT (datetime('now', 'localtime'))
					)`
		result, err := slte.DB().Exec(sqlTablle)
		if err != nil {
			t.Error(err)
			return
		}
		var n int64
		n, err = result.RowsAffected()
		if err != nil {
			t.Error(err)
			return
		}
		if n > 0 {

		}
	}

	{
		var (
			err  error
			stmt *sql.Stmt
		)
		stmt, err = slte.DB().Prepare("INSERT INTO userinfo(username, departname)  values(?, ?)")
		if err != nil {
			t.Error(err)
			return
		}
		defer stmt.Close()

		for i := 0; i < 1000; i++ {
			var res sql.Result
			res, err = stmt.Exec("astaxie", "研发部门")
			if err != nil {
				t.Error(err)
				return
			}
			var n int64
			n, err = res.LastInsertId()
			if err != nil {
				t.Error(err)
				return
			}

			if n > 0 {

			}
		}

	}
}
