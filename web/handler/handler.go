package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

//HandlerRoute - This interface provides the object with the responsibility to handle requests.
type HandlerRoute interface {
	Props() (path string, method string)
	Handler(w http.ResponseWriter, r *http.Request)
}

// ResponseError - This method provides to return errors in JSON format.
func ResponseError(w http.ResponseWriter, httpStatus int, message string) {

	data := map[string]string{}
	data["error"] = message

	ResponseJSON(w, httpStatus, data)
}

// ResponseJSON - This method provides to return the results in JSON format.
func ResponseJSON(w http.ResponseWriter, httpStatus int, payload interface{}) {

	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(data)
}
