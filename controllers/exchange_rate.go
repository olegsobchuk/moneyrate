package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/olegsobchuk/moneyrate/models/exchanger"
)

// ExchangeIndex index page with form
func ExchangeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route: ExchangeIndex, Method: ", r.Method)
	tmpl := template.Must(template.ParseFiles("templates/welcome/index.html", "templates/welcome/_header.html"))
	data := map[string]interface{}{
		"currentDate": time.Now().Format("02/01/2006"),
		"Exchanger":   exchanger.New(),
	}
	tmpl.Execute(w, data)
}

// ExchangeFindRate send request for fetch rate for specific day
func ExchangeFindRate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route: ExchangeFindRate, Method: ", r.Method)

	exchanger := exchanger.New()
	r.ParseForm()
	defer r.Body.Close()

	decoder := schema.NewDecoder()
	if err := decoder.Decode(&exchanger, r.Form); err != nil {
		fmt.Println(err)
	}
	respondWithJSON(w, r, http.StatusCreated, exchanger)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
