package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//type HelloResponse struct {
//	name string `json:"name"`
//}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		return
	}
	str := "name" + name
	json.Marshal([]byte(str))
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(str))
	w.WriteHeader(http.StatusOK)
	log.Println("&s &S", time.Now().Format("2006/01/02 15:04:05"), fmt.Sprintf(`{"name":"%s"}`, name))
}
