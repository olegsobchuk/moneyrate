package customConverter

import (
	"fmt"
	"reflect"
	"time"
)

// StringToDate convert string to date in format dd/mm/yyyy
func StringToDate(strDate string) reflect.Value {
	date, err := time.Parse("2/1/2006", strDate)
	if err != nil {
		fmt.Println("Failed decode string to date! Error: ", err)
		return reflect.Value{}
	}
	return reflect.ValueOf(date)
}
