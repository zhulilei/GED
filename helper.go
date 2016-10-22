package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func writeBody(item interface{}, w http.ResponseWriter, status int) {
	if body, err := json.Marshal(item); err == nil {
		w.WriteHeader(status)
		io.WriteString(w, string(body))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func parseRequest(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	var item map[string]string
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return item, err
	}
	return item, nil
}

func DataPack(code int, data interface{}, message string) (result Result) {
	result.Data = data
	result.Code = code
	result.Message = message
	return
}

func DealTime(timestring string) string {
	t, _ := time.Parse("2006/01/02-15:04", timestring) //time.time
	//fmt.Println(t, err)
	d, _ := time.ParseDuration("-8h")
	timestamp := t.Add(d)
	//fmt.Println(timestamp, err)
	return timestamp.Format("2006/01/02-15:04") //string
}
