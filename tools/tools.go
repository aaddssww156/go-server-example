package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server-example/services"

	"github.com/joho/godotenv"
)

type Envelope map[string]interface{}

type Message struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

var MessageLogs = &Message{
	InfoLog:  infoLog,
	ErrorLog: errorLog,
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxByte := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return fmt.Errorf("error reading json data from request: %s", err.Error())
	}

	if err := dec.Decode(&struct{}{}); err != nil {
		return fmt.Errorf("json body must have only one object")
	}

	return nil
}

func WriteJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshaling json: %w", err)
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return fmt.Errorf("error writing json data: %w", err)
	}
	return nil
}

func ErrorJson(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload services.JsonResponse

	payload.Error = true
	payload.Message = err.Error()
	WriteJson(w, statusCode, payload)

}

func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %s", err.Error())
	}
	return nil
}
