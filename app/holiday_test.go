package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFetchHolidays(t *testing.T) {
	// Crear un servidor HTTP falso
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"data": [{"date": "2023-05-01T00:00:00Z","title": "Día Nacional del Trabajo","type": "Civil","inalienable": true,"extra": "Civil e Irrenunciable"}],"status": {"code": 200,"message": "Proceso ejecutado sin problemas"}}`))
	}))
	// Cerrar el servidor cuando termine la prueba
	defer server.Close()

	// Crear un cliente HTTP con la URL del servidor falso como endpoint
	_ = &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(server.URL)
			},
		},
	}

	// Llamar a FetchHolidays
	_, err := FetchHolidays()
	if err != nil {
		t.Errorf("FetchHolidays() failed, expected %v, got %v", nil, err)
	}
}
