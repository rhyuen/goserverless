package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Arbitrary struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("pg_port")
	convPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		fmt.Print(w, "env var issue ")
	}
	host := os.Getenv("pg_host")
	user := os.Getenv("pg_user")
	password := os.Getenv("pg_password")
	dbname := os.Getenv("pg_dbname")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=enable", host, convPort, user, password, dbname)

	var dbError error
	conn, dbError := sql.Open("postgres", psqlInfo)
	if dbError != nil {
		fmt.Print(w, "DB Conn error: %s", dbError)
	}

	log.Println(conn)
	log.Println(dbError)
	rows, queryErr := conn.Query("SELECT NAME, CONTENT FROM ARBITRARY")
	if queryErr != nil {
		fmt.Print(w, "Query Error: %s", queryErr)
	}
	log.Println(rows)
	log.Println(queryErr)
	defer rows.Close()
	allArbitrary := []Arbitrary{}

	for rows.Next() {
		var currArb Arbitrary
		if err := rows.Scan(&currArb.Name, &currArb.Content); err != nil {
			fmt.Print(w, "An Error occurred: %s", err)
		}
		allArbitrary = append(allArbitrary, currArb)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allArbitrary)
}
