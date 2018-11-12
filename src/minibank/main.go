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

	go models.InitDB(dbDoneCh)
	defer models.Database.Close()

	if models.CassandraEnabled {
		go models.InitCassandra()
		defer models.CassandraSession.Close()
	}

	http.HandleFunc("/api/account/register", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			// 405 may not be the right number for a post error
			http.Error(w, "405", http.StatusMethodNotAllowed)
			return
		}
 		w.Write([]byte("passes\n"))
 		//id : r
		fmt.Println(r.Body)
		/*if dbDone {
			handlers.RegisterHandler(w, r)
		} else {
			handlers.ServerUnavailableHandler(w, r)
		}*/
	})
	http.HandleFunc("/api/account/login", validateDBConn(handlers.LoginHandler, &dbDone))
	http.HandleFunc("/api/account/token", validateDBConn(handlers.TokenHandler, &dbDone))
	http.HandleFunc("/api/account/sessions", handlers.AuthValidationMiddleware(handlers.SessionListHandler))

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
