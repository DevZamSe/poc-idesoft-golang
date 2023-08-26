package app

import (
	"log"
	"time"
)

type Holiday struct {
	Date        time.Time `json:"date" xml:"date"`
	Title       string    `json:"title" xml:"title"`
	Type        string    `json:"type" xml:"type"`
	Inalienable bool      `json:"inalienable" xml:"inalienable"`
	Extra       string    `json:"extra" xml:"extra"`
}

func TransformStr2TimeInterface(inputInterface ResponseCore) (responseHoliday []Holiday, err error) {

	for _, h := range inputInterface.Data {

		var transformResponseCore Holiday
		transformResponseCore.Date, err = TransformStr2Time(h.Date)
		transformResponseCore.Extra = h.Extra
		transformResponseCore.Inalienable = h.Inalienable
		transformResponseCore.Title = h.Title
		transformResponseCore.Type = h.Type

		responseHoliday = append(responseHoliday, transformResponseCore)
	}

	return
}

func TransformStr2Time(inputDate string) (date time.Time, err error) {
	date, err = time.Parse("2006-01-02", inputDate)
	if err != nil {
		log.Fatalf("TransformStr2Time :: %v", err)
	}
	return
}
