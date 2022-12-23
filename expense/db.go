package expense

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type Expense struct {
	Id     int
	Title  string
	Amount float32
	Note   string
	Tags   []string
}

type DBQuery interface {
	QueryRow(query string, args ...any) Row
}

type Row interface {
	Err() error
	Scan(dest ...any) error
}

func InitDB() *sql.DB {
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createTb := `CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	return db
}
