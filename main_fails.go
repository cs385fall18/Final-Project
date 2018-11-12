package main

import (
	"fmt"
	//"minibank/handlers"
	"minibank/models"
	"net/http"
	"os"
)
func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "It works again!")
}

func main() {
	// Connect to the database

	dbDone_chan := make(chan bool)
	dbDone := false

	dbConn := fmt.Sprintf("minibank:minibank@tcp(mysql)/minibank")
	go models.InitDB(dbConn/*, dbDone_chan*/)
	defer models.Database.Close()
	http.HandleFunc("/", index_handler)

	/*
	http.HandleFunc()
	*/
	// is /api needed?
	http.HandleFunc("/api/account/register", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			// 405 may not be the right number for a post error
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("passes")
		//id : r
		fmt.Println(r.Body)
		/*if dbDone {
			handlers.RegisterHandler(w, r)
		} else {
			handlers.ServerUnavailableHandler(w, r)
		}*/
	})

	http.HandleFunc("/api/account/token", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" || r.Method != "DELETE" {
			// 405 may not be the right number for a post error
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return
		}
		if r.Method == "POST" {
			fmt.Println("this is a POST url")
			fmt.Println(r)

		}
		if r.Method == "DELETE" {
			fmt.Println("this is a DELETE url")
			fmt.Println(r)
		}
	})

	http.HandleFunc("/api/content", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" {
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println(r)

	})

	http.HandleFunc("/api/content/{content_id}", func(w http.ResponseWriter, r *http.Request){
		if r.Method != "GET" || r.Method != "PUT" || r.Method != "DELETE" {
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return
		}
		if r.Method == "GET" {
			fmt.Println("this is a GET content url")

			fmt.Println(r)
		}
		if r.Method == "PUT" {
			fmt.Println("this is a PUT content url")
			fmt.Println(r)
		}
		if r.Method == "DELETE" {
			fmt.Println("this is a DELETE content url")
			fmt.Println(r)
		}
	})

	http.HandleFunc("/api/subscriptions/", func(w http.ResponseWriter, r *http.Request){

		if r.Method != "POST" || r.Method != "GET" {
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return

		}
		if r.Method == "POST" {
			fmt.Println("this is a POST subscriptions url")
			fmt.Println(r)

		}
		if r.Method == "GET" {
			fmt.Println("this is a GET subscriptions url")
			fmt.Println(r)

		}

	})

	http.HandleFunc("/api/subscriptions/subscribers{username}", func(w http.ResponseWriter, r *http.Request){

		if r.Method != "GET" {
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return

		}
		fmt.Println(r)
	})
	http.HandleFunc("/api/subscriptions/subscribers{tag}", func(w http.ResponseWriter, r *http.Request){

		if r.Method != "GET" {
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return

		}
		fmt.Println(r)
	})


	http.HandleFunc("/api/feed", func(w http.ResponseWriter, r *http.Request){

		if r.Method != "GET" {
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return

		}
		fmt.Println(r)
	})

	http.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request){

		if r.Method != "POST" {
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return

		}
		fmt.Println(r)
	})

	/*http.HandleFunc("/api/account/register", func(w http.ResponseWriter, r *http.Request) {
		if dbDone {
			handlers.RegisterHandler(w, r)
		} else {
			handlers.ServerUnavailableHandler(w, r)
		}
	})*/
	go updateDBDone(&dbDone, dbDone_chan)
	http.ListenAndServe(/*port()*/":8080", nil)
}

func updateDBDone(dbdone *bool, dbDone_ch <-chan bool) {
	*dbdone = <-dbDone_ch
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
