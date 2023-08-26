package main

import (
	"log"
	"net/http"
	"poc/idesoft/app"
)

func main() {
	// Inicializar la lista de feriados
	app.OrchestationPersistence()

	// Definir el manejador del endpoint
	http.HandleFunc("/holidays", app.HolidayHandler)

	// Iniciar el servidor
	log.Println("CMD Server :: Listen on port 8080")
	http.ListenAndServe(":8080", nil)
}
