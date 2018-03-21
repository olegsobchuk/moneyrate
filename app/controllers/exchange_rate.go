package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/olegsobchuk/moneyrate/app/lib/custom_converter"
	"github.com/olegsobchuk/moneyrate/app/models/exchanger"
	"github.com/olegsobchuk/moneyrate/app/services/bank_rate"
)

// ExchangeIndex index page with form
func ExchangeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Route: ExchangeIndex, Method: ", r.Method)
	tmpl := template.Must(template.ParseFiles("app/templates/welcome/index.html", "app/templates/welcome/_header.html"))
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
	decoder.RegisterConverter(time.Time{}, customConverter.StringToDate)
	if err := decoder.Decode(&exchanger, r.Form); err != nil {
		fmt.Println(err)
	}
	rates := bankRate.GetRate("usd", exchanger.Date)
	respondWithJSON(w, r, http.StatusCreated, rates)
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
