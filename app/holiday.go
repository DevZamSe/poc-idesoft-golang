package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type HolidayCore struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Inalienable bool   `json:"inalienable"`
	Extra       string `json:"extra"`
}

type ResponseCore struct {
	Status string        `json:"status"`
	Data   []HolidayCore `json:"data"`
}

// Obtener los datos desde el servicio web externo
func FetchHolidays() (responseCore ResponseCore, err error) {

	resp, err := http.Get("https://api.victorsanmartin.com/feriados/en.json")
	if err != nil {
		log.Fatalf("FetchHolidays :: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("FetchHolidays :: %v", err)
		return
	}

	err = json.Unmarshal(body, &responseCore)
	if err != nil {
		log.Fatalf("FetchHolidays :: %v", err)
		return
	}

	return
}
