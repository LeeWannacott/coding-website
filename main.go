package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func codeChallenge(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "GET":
		fmt.Println("GET!!!")
		http.ServeFile(w, request, "index.html")

	case "POST":
		// Call ParseForm() to parse the raw query and update request.PostForm and request.Form.
		fmt.Printf("POST!: ")
		if err := request.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(w, "Post from website! request.PostFrom = %v\n", request.PostForm)
		code := strings.Join(request.Form["code"], "\n")
		fmt.Printf("code: %v\n", code)

		// Create a new Cmd to run the Node.js process
		cmd := exec.Command("node", "-e", code)

		// Run the Node.js process and capture its output
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("output: ", string(output))

		http.ServeFile(w, request, "index.html")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

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
