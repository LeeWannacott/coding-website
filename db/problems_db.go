package codeProblems

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type codeProblem struct {
	ProblemID    int
	Question     string
	CodeFilePath string
	Output       string
	Language     string
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func openDB(filePath string) *sql.DB {
	db, err := sql.Open("sqlite3", filePath)
	checkError(err)
	// Check if the database file exists
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		// If the file doesn't exist, create it
		file, err := os.Create(filePath)
		checkError(err)
		file.Close()
	}
	return db
}

func createTable(db *sql.DB) string {
	codeProblemsTableName := "code_problems"
	createCodeProblemsTable :=
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (%v INTEGER NOT NULL PRIMARY KEY, %v TEXT NOT NULL, %v TEXT
		NOT NULL, %v TEXT NOT NULL, %v TEXT NOT NULL)`, codeProblemsTableName, "problem_id", "question", "code", "output", "language")
	_, err := db.Exec(createCodeProblemsTable)
	checkError(err)
	return codeProblemsTableName
}

func insertCodeProblem(db *sql.DB, codeProblem codeProblem) {
	insertCodeProblem := fmt.Sprintf(`INSERT OR IGNORE INTO code_problems (problem_id, question, code, output, language) VALUES (?, ?, ?, ?, ?)`)
	_, err := db.Exec(insertCodeProblem, codeProblem.ProblemID, codeProblem.Question, codeProblem.CodeFilePath, codeProblem.Output, codeProblem.Language)
	checkError(err)
}

func selectCodeProblems(db *sql.DB, tableName string) {
	selectCodeProblems := fmt.Sprintf("SELECT * FROM %v", tableName)
	rows, err := db.Query(selectCodeProblems)
	checkError(err)
	fmt.Println(rows)
	defer rows.Close()
}

func InitProblemsDatabase() {
	db := openDB("./problems.db")
	defer db.Close()
	tableName := createTable(db)

	problemID := 1
	codingLanguage := "javascript"
	codeProblem := codeProblem{
		ProblemID:    problemID,
		Question:     "Refactor this code so that you don't need to declare the empty variable. Hint use: <a>Array.prototype.map()<a>",
		CodeFilePath: fmt.Sprintf("../code_problems/%v/problem_%d.js", codingLanguage, problemID),
		Output:       "[ 'it was a good book', 'average book', 'worst book I have ever read' ]",
		Language:     codingLanguage,
	}
	insertCodeProblem(db, codeProblem)
	fmt.Println("Inserted code problem into database")
	selectCodeProblems(db, tableName)

}
