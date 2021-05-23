package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_POSTGRES_URI")
	driverName := os.Getenv("DATABASE_DRIVER")
	return sql.Open(driverName, uri)
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
