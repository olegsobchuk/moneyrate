package bankRate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BankRate struct
type BankRate struct {
	Currency string  `json:"cc"`
	Rate     float32 `json:"rate"`
	Date     time.Time
}

// GetRate - send API request
func GetRate(currency string, date time.Time) string {
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
	fmt.Println(string(body))
	return string(body)
}
