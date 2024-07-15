package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DedMokus/go-ml-vk-test-task/internal/document"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseProcessor interface {
	Connect() error
	Disconnect() error
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

type PostgreSQLProcessor struct {
	DB *sqlx.DB
}

func (pdb *PostgreSQLProcessor) Connect() error {
	var connStr string

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db_name := os.Getenv("POSTGRES_DB")

	connStr = fmt.Sprintf("host=db port=5432 user=%s password=%s dbname=%s sslmode=disable", username, password, db_name)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database not connected: %v", err)
	}

	pdb.DB = db
	return err
}

func (pdb *PostgreSQLProcessor) Disconnect() error {
	if pdb.DB != nil {
		return pdb.DB.Close()
	}
	return nil
}

func (pdb *PostgreSQLProcessor) Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := pdb.DB.Query(query)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	return rows, err
}

func (pdb *PostgreSQLProcessor) QueryRow(query string, args ...any) *sql.Row {
	doc := document.GenerateRandomDocument(query)

	sqlString := `
	INSERT INTO docs (url, pubdate, fetchtime, text, firstfetchtime)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	id := 0
	err := pdb.DB.QueryRow(sqlString, doc.Url, doc.PubDate, doc.FetchTime, doc.Text, doc.FirstFetchTime).Scan(&id)
	if err != nil {
		log.Fatalf("Error insert row: %v", err)
		return nil
	}
	return nil
}
