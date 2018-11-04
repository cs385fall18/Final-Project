package main

import (
	"fmt"
	"minibank/handlers"
	"minibank/models"
	"net/http"
	"os"
)

func main() {
	// Connect to the database

	dbDone_chan := make(chan bool)
	dbDone := false

	dbConn := fmt.Sprintf("minibank:minibank@tcp(mysql)/minibank")
	go models.InitDB(dbConn, dbDone_chan)
	defer models.Database.Close()
	/*
	http.HandleFunc()
	*/
	http.HandleFunc("/api/account/register", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			// 405 may not be the right number for a post error
			http.Error(w, http.StatusTest(405), http.StatusMethodNotAllowed)
			return
		}
		//id : r
		fmt.PrintLn(r)
		/*if dbDone {
			handlers.RegisterHandler(w, r)
		} else {
			handlers.ServerUnavailableHandler(w, r)
		}*/
	})



	/*http.HandleFunc("/api/account/register", func(w http.ResponseWriter, r *http.Request) {
		if dbDone {
			handlers.RegisterHandler(w, r)
		} else {
			handlers.ServerUnavailableHandler(w, r)
		}
	})*/
	go updateDBDone(&dbDone, dbDone_chan)
	http.ListenAndServe(port(), nil)
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
