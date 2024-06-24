package home

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rogeriopiatek/goLweet/backend/db"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	conn := db.GetDB()

	rows, err := conn.Query("SELECT NOW()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var currentTime string
	for rows.Next() {
		if err := rows.Scan(&currentTime); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintf(w, "Current time: %s", currentTime)
}
