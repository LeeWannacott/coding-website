package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/LeeWannacott/coding-website/problemdb"
)

type FormData struct {
	Code   string
	Output string
}

func runCodeAsChildProcess(code string) string {
	fmt.Printf("code: %v\n", code)
	// Create a new Cmd to run the Node.js process
	cmd := exec.Command("node", "-e", code)
	// Run the Node.js process and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("output: ", string(output))
	return string(output)
}

func parseTemplate(w http.ResponseWriter, formData FormData) {
	// Parse the template file
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with the data
	err = tmpl.Execute(w, formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func codeChallenge(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {

	case "GET":

		fmt.Println("GET!!!")

		// Create a FormData struct with the data to be sent to the template
		problem, err := os.ReadFile("./js_problems/problem_1.js")
		if err != nil {
			panic(err)
		}
		fmt.Print(string(problem))
		Code := string(problem)
		formData := FormData{
			Code: Code,
		}
		parseTemplate(w, formData)

	case "POST":
		// Call ParseForm() to parse the raw query and update request.PostForm and request.Form.
		fmt.Printf("POST!: ")
		if err := request.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		code := strings.Join(request.Form["code"], "\n")
		codeOutput := runCodeAsChildProcess(code)
		fmt.Print("code:", codeOutput)
		request.Form["code-output"] = strings.Split(codeOutput, " ")

		formData := FormData{
			Code:   request.Form["code"][0],
			Output: codeOutput,
		}

		parseTemplate(w, formData)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	problemdb.ProblemDatabase()

	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("tmp"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("node_modules"))))
	http.HandleFunc("/", codeChallenge)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
