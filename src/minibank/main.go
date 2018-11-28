package main

import (
	"minibank/handlers"
	"minibank/models"
	"net/http"
	"os"
	"fmt"
)

func main() {
	// Connect to the database

	dbDoneCh := make(chan bool)
	dbDone := false
	//dbConn := fmt.Sprintf("minibank:minibank@tcp(mysql)/minibank")

	go models.InitDB(dbDoneCh)
	defer models.Database.Close()

	/*
	if models.CassandraEnabled {
		go models.InitCassandra()
		defer models.CassandraSession.Close()
	}*/
	/*http.HandleFunc("/api/account/register", func(w http.ResponseWriter, r *http.Request) {
		if dbDone {
			handlers.RegisterHandler(w, r)
		} else {
			handlers.ServerUnavailableHandler(w, r)
		}
	})*/
	http.HandleFunc("/api/account/register", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			// 405 may not be the right number for a post error
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return
		}
		//if r.Method == "POST" {

		//}
		handlers.RegisterHandler(w, r)
 		//w.Write([]byte("passes\n"))
 		//id : r
		//fmt.Println(r.Body)
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
			// what the json input looks like
			// '{"username": "john", "password": "john123456"}'
			w.Write([]byte("this is a POST url\n"))
			handlers.TokenHandler(w, r)
			//fmt.Println(r)

		}
		if r.Method == "DELETE" {
			// what the json input looks like
			// '{"token" : "23456765432345676543234676543"}'
			// how to know the token is bound to the user?
			w.Write([]byte("this is a DELETE url\n"))


			// fmt.Println(r)
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
			w.Write([]byte("this is a GET content url\n"))

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
	/*http.HandleFunc("/api/account/login", validateDBConn(handlers.LoginHandler, &dbDone))
	http.HandleFunc("/api/account/token", validateDBConn(handlers.TokenHandler, &dbDone))
	http.HandleFunc("/api/account/sessions", handlers.AuthValidationMiddleware(handlers.SessionListHandler))
	*/

	go updateDBDone(&dbDone, dbDoneCh)
	http.ListenAndServe(port(), nil)
}

func validateDBConn(next http.HandlerFunc, dbDone *bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *dbDone {
			next(w, r)
		} else {
			handlers.ServerUnavailableHandler(w, r)
		}
	})
}

func updateDBDone(dbdone *bool, dbDoneCh <-chan bool) {
	*dbdone = <-dbDoneCh
}

// port looks up service listening port
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
//func message()
