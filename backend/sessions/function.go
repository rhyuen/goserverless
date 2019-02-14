package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("gobuypizzarightnow")
	store = sessions.NewCookieStore(key)
)

func Handler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "Pizza")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(w, "sessions, guise. %s", session)
}
