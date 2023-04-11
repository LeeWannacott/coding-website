package main

import (
	"fmt"
	"log"
	"net/http"
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
		code := request.FormValue("code")
		validCode := "test"
		if code == validCode {
			fmt.Println("Valid Code")
		}
		fmt.Println(code)
		// fmt.Fprintf(w, "Name = %s\n", code)
		// fmt.Println("Code: ", request)

		http.ServeFile(w, request, "index.html")
		// http.PostForm("http://localhost:8080", request.Form)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func test(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/test" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch request.Method {
	case "GET":
		fmt.Println("noob")
		http.ServeFile(w, request, "index.html")

	case "POST":
		// Call ParseForm() to parse the raw query and update request.PostForm and request.Form.
		fmt.Printf("tube!: ")
		if err := request.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		// fmt.Fprintf(w, "Post from website! request.PostFrom = %v\n", request.PostForm)
		code := request.FormValue("code")
		validCode := "test"
		if code == validCode {
			fmt.Println("Valid Code")
		}
		fmt.Println(code)
		// fmt.Fprintf(w, "Name = %s\n", code)
		// fmt.Println("Code: ", request)

		// http.ServeFile(w, request, "index.html")
		// http.Redirect(w, request, "http://localhost:8080", http.StatusSeeOther)
		// http.PostForm("http://localhost:8080", request.Form)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", codeChallenge)
	http.HandleFunc("/test", test)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
