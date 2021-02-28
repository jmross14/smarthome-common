package smarthome_common

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

// RespondWithJSON creates the JSON response that will be sent back in the response.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		RespondWithJSONError(w, http.StatusInternalServerError, "Unable to Marshal JSON data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

// RespondWithJson creates a JSON response error that follows error : errorMessage.
func RespondWithJSONError(w http.ResponseWriter, code int, errorMessage string) {
	message := map[string]string{"error": errorMessage}

	RespondWithJSON(w, code, message)
}

func GetIPAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
