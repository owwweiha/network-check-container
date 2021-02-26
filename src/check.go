package main

import (
	"time"
	"net"
	"net/http"
	"log"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
)


type Connection struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Protocol string `json:"protocol"`
}

func main() {
    r := mux.NewRouter()

    api := r.PathPrefix("/api/v1").Subrouter()
    api.HandleFunc("/connect", checkConnection).Methods(http.MethodGet)
    api.HandleFunc("/", notFound)
    log.Fatal(http.ListenAndServe(":8080", r))
}

func checkConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var connection Connection
	err := json.NewDecoder(r.Body).Decode(&connection)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status": "error", "message": "bad request"}`))
		fmt.Println("Bad request")
		return
	}

	if connection.Host == "" {
	        badRequest(w, r)
		return
	}

	if connection.Port == "" {
		badRequest(w, r)
		return
	}

	if connection.Protocol == "" {
		badRequest(w, r)
		return
	}

	fmt.Println("Received request: " + connection.Host + ":" + connection.Port + "/" + connection.Protocol)

	timeout := time.Second
	conn, err := net.DialTimeout(connection.Protocol, net.JoinHostPort(connection.Host, connection.Port), timeout)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"status": "error", "host": %q, "port": %q, "protocol": %q, "message": %q}`, connection.Host, connection.Port, connection.Protocol, err)))
		fmt.Println("Connecting failed:", err)
	}

	if conn != nil {
		defer conn.Close()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"status": "success", "host": %q, "port": %q, "protocol": %q, "message": "connection successful"}`, connection.Host, connection.Port, connection.Protocol)))
		fmt.Println("Connection succesful:", net.JoinHostPort(connection.Host, connection.Port))
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(`{"status": "error", "message": "bad request"}`))
        fmt.Println("Bad request")
}
