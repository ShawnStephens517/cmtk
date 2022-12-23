package query

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Record struct {
	ID          int
	System      string `json:"system"`
	Description string `json:"description"`
	User        string `json:"user"`
	Date        string `json:"date"`
}

// TODO Must Change From Explicit to User Input UN&Pass
func Recentrecords(db *sql.DB) ([]Record, error) {
	connStr := "postgres://cmtk:cmtk@localhost/maintenance?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM records ORDER BY id DESC LIMIT 25")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records := []Record{}
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.System, &r.Description); err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	return records, nil
}

var record, _ = Recentrecords()

func RecordsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		recordsJSON, _ := json.Marshal(record)
		w.Write(recordsJSON)
	case "POST":
		var record Record
		if err := json.NewDecoder(r.Body).Decode(&record);
		//Input()

		err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		record = append(record, r)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
