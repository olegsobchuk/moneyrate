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
	type Alias BankRate
	aux := &struct {
		Date string `json:"exchangedate"`
		*Alias
	}{
		Alias: (*Alias)(rate),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rate.Date, _ = time.Parse("02.01.2006", aux.Date)
	return nil
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
