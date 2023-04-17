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

type problemsList struct {
	Problems []codeProblem
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

func SelectCodeProblems(db *sql.DB, tableName string) problemsList {
	selectCodeProblems := fmt.Sprintf("SELECT * FROM %v", tableName)
	rows, err := db.Query(selectCodeProblems)
	fmt.Println(rows)
	checkError(err)
	defer rows.Close()

	var codeProblemsList problemsList
	for rows.Next() {
		var problem_id int
		var question string
		var code string
		var output string
		var language string

		// Scan the values in the current row into variables

		err := rows.Scan(&problem_id, &question, &code, &output, &language)
		codeSolution := codeProblem{
			ProblemID:    problem_id,
			Question:     question,
			CodeFilePath: code,
			Output:       output,
			Language:     language,
		}
		checkError(err)
		codeProblemsList.Problems = append(codeProblemsList.Problems, codeSolution)
	}

	// Access the values as needed
	return codeProblemsList
}

func GetTableName(tableName string) string {
	return tableName
}

func InitProblemsDatabase() problemsList {
	db := openDB("./problems.db")
	defer db.Close()
	tableName := createTable(db)
	GetTableName(tableName)

	problemID := 1
	codingLanguage := "javascript"
	codeProblem1 := codeProblem{
		ProblemID:    problemID,
		Question:     "Refactor this code so that you don't need to declare the empty variable. Hint use: <a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/map' target='_blank'>Array.prototype.map()<a>",
		CodeFilePath: fmt.Sprintf("./code_problems/%v/problem_%d.js", codingLanguage, problemID),
		Output:       "[ 'it was a good book', 'average book', 'worst book I have ever read' ]",
		Language:     codingLanguage,
	}
	insertCodeProblem(db, codeProblem1)

	problemID = 2
	codingLanguage = "javascript"
	codeProblem2 := codeProblem{
		ProblemID:    problemID,
		Question:     "Refactor this code to use <a href='https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduce'>.reducer()</a> to sum the stars from each audiobook.",
		CodeFilePath: fmt.Sprintf("./code_problems/%v/problem_%d.js", codingLanguage, problemID),
		Output:       "9",
		Language:     codingLanguage,
	}
	insertCodeProblem(db, codeProblem2)

	fmt.Println("Inserted code problem into database")
	problemsList := SelectCodeProblems(db, tableName)

	return problemsList

}
