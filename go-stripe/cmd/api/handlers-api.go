package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) GetPyamentIntent(w http.ResponseWriter, r *http.Request) {
	var payload stripePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	j := jsonResponse{
		OK: true,
	}
	out, err := json.MarshalIndent(j, "", " ")
	if err != nil {
		app.errorLog.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
