package bankRate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// BankRate struct
type BankRate struct {
	Currency string    `json:"cc"`
	Rate     float32   `json:"rate"`
	Date     time.Time `json:"exchangedate"`
}

// UnmarshalJSON unmarshal Date
func (rate *BankRate) UnmarshalJSON(data []byte) error {
	type RateAlias BankRate
	aux := &struct {
		Date string `json:"exchangedate"`
		*RateAlias
	}{
		RateAlias: (*RateAlias)(rate),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rate.Date, _ = time.Parse("2.1.2006", aux.Date)
	return nil
}

// MarshalJSON marshal data
func (rate *BankRate) MarshalJSON() ([]byte, error) {
	type dupBankRate BankRate
	basicRate := struct {
		Date     string `json:"date"`
		BankRate *dupBankRate
	}{
		BankRate: (*dupBankRate)(rate),
		Date:     rate.Date.Format("02/01/2006"),
	}
	jsn, err := json.Marshal(basicRate)
	if err != nil {
		return []byte{}, err
	}
	return jsn, nil
}

// GetRate - send API request
func GetRate(currency string, date time.Time) ([]BankRate, error) {
	dateString := date.Format("20060102")
	url := fmt.Sprintf(
		"https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?valcode=%s&date=%s&json",
		currency,
		dateString,
	)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Response has been failed")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	rates := make([]BankRate, 0)
	if err = json.Unmarshal(body, &rates); err != nil {
		log.Println(err)
		fmt.Println(err)
		return nil, err
	}

	return rates, nil
}
