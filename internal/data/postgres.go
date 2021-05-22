package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
)

func getConnection() (*sql.DB, error) {
	//uri := os.Getenv("DATABASE_URI")
	return sql.Open("postgres", "postgres://lucas_piegas:1234@127.0.0.1:5432/go_store?sslmode=disable")
}

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
