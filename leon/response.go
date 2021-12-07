package leon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorLog(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return
}

func ResponseWriter(w http.ResponseWriter, responseCode int, msg string, data interface{}) {
	resp := Response{
		Code:    responseCode,
		Message: msg,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	err := json.NewEncoder(w).Encode(resp)
	ErrorLog(err)
}
