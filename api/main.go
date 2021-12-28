package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID     string `json:"id"`
	Today  bool   `json:"today"`
	Streak int    `json:"streak"`
}

var users = []User{{
	ID:     "CleyL",
	Today:  true,
	Streak: 50,
}, {
	ID:     "tourist",
	Today:  false,
	Streak: 0,
}}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&users); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	http.HandleFunc("/user", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
