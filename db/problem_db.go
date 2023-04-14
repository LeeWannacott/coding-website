package problemdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ProblemDatabase() {
	db, err := sql.Open("sqlite3", "./coding.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	codeProblemsTableName := "code_problems"
	createCodeProblemsTable :=
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (id INTEGER PRIMARY KEY AUTOINCREMENT, %v TEXT NOT NULL, %v TEXT
		NOT NULL, %v TEXT NOT NULL, %v TEXT NOT NULL)`, codeProblemsTableName, "question", "code", "output", "language")
	_, err = db.Exec(createCodeProblemsTable)

	if err != nil {
		log.Fatal(err)
	}
	selectCodeProblems := fmt.Sprintf("SELECT * FROM %v", codeProblemsTableName)
	rows, err := db.Query(selectCodeProblems)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
	defer rows.Close()
}
