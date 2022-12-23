package query

import (
	"database/sql"
	"time"
)

func Insert() {
	now := time.Now()
	formatnow := now.Format("12/24/2022 15:42:00")
	// TODO Must Change From Explicit to User Input UN&Pass
	connStr := "postgres://cmtk:cmtk@localhost/maintenance?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	r := Record{System: "SR-Test", Description: "First Entry", User: "Shawn Stephens", Date: formatnow}
	res, err := db.Exec("INSERT INTO records (system, description, user, date) VALUES ($1, $2)", r.System, r.Description)
	if err != nil {
		return err
	}
	return
}
