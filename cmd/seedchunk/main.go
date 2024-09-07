package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	names := mustOpenFile()
	db := mustConnDB()

	start := time.Now()

	chunkSize := 10000
	for _, chunk := range chunk(names, chunkSize) {
		query := prepareQuery(len(chunk))
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatalf("Error preparing query: %v", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(chunk...)
		if err != nil {
			log.Fatalf("Error executing query: %v", err)
		}
	}

	fmt.Println("Seeding completed in", time.Since(start).Milliseconds(), "ms")
}

func mustOpenFile() []any {
	var cnt string
	fmt.Println("Enter the number of records you want to seed: ")
	fmt.Scanln(&cnt)

	file, err := os.Open(fmt.Sprintf("./data/names_%s.csv", cnt))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	names := make([]any, 0, len(records))
	for _, record := range records {
		names = append(names, record[0])
	}

	return names
}

func mustConnDB() *sql.DB {
	dsn := "dev:dev@tcp(localhost:3306)/dev"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db
}

func prepareQuery(length int) string {
	p := make([]string, 0, length)

	for i := 0; i < length; i++ {
		p = append(p, "(?)")
	}

	return fmt.Sprintf("INSERT INTO users (name) VALUES %s", strings.Join(p, ","))
}

func chunk[T any](slice []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}

	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
