package app

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"time"
)

var HolidayPersistence []Holiday

type ResponseInterface struct {
	Data   []Holiday `json:"data" xml:"data"`
	Status Status    `json:"status" xml:"status"`
}

type Status struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func OrchestationPersistence() {
	// Invocar ambassador external service
	responseHolidayPersistence, err := FetchHolidays()
	if err != nil {
		log.Fatalf("OrchestationPersistenceError :: %v", err)
	}

	log.Println("OrchestationPersistence :: ", len(responseHolidayPersistence.Data), "records found!")

	// Parsear ambassador service
	HolidayPersistence, err = TransformStr2TimeInterface(responseHolidayPersistence)
	if err != nil {
		log.Fatalf("OrchestationPersistenceError :: %v", err)
	}

	log.Println("OrchestationPersistence :: successfully load of synchronized list!")
}

func HolidayHandler(w http.ResponseWriter, r *http.Request) {

	// Leer parámetros de la solicitud
	query := r.URL.Query()
	holidayType := query.Get("type")
	startDateStr := query.Get("startDate")
	endDateStr := query.Get("endDate")
	codeExecution := 200
	messageExecution := "Proceso ejecutado sin problemas"

	var startDate time.Time
	var endDate time.Time

	// Parsear Str to Time.time
	const layout = "2006-01-02"
	if startDateStr != "" {
		startDate, _ = TransformStr2Time(startDateStr)
	}
	if endDateStr != "" {
		endDate, _ = TransformStr2Time(endDateStr)
	}

	// Filtrar los feriados según los parámetros
	var filtered []Holiday
	for _, h := range HolidayPersistence {
		if holidayType != "" && h.Type != holidayType {
			continue
		}

		if startDateStr != "" && endDateStr != "" {
			// Filtrado entre fechas
			if h.Date.Before(startDate) || h.Date.After(endDate) {
				continue
			}
		} else if startDateStr != "" {
			// Filtrado de una sola fecha de inicio
			if h.Date.Before(startDate) {
				continue
			}
		}
		filtered = append(filtered, h)
	}

	log.Println("HolidayHandler :: ", len(filtered), "records send!")

	if len(filtered) == 0 {
		codeExecution = 400
		messageExecution = "No se encontraron registros!"
	}

	// Establecer response
	responseInterface := ResponseInterface{
		Data: filtered,
		Status: Status{
			Code:    codeExecution,
			Message: messageExecution,
		},
	}

	// Establecer el tipo de contenido según la solicitud
	contentType := r.Header.Get("Accept")
	if contentType == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(responseInterface)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseInterface)
	}

	log.Println("=====================================================================")
}
