package handlers

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendJson", sendJson)
}

func sendJson(w http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardanstudios.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&u)
}
