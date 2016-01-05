package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"proxy/db"
	"proxy/model"

	"github.com/gorilla/mux"
)

func NewMux() http.Handler {
	Rmux := mux.NewRouter()
	Rmux.HandleFunc("/api/v1/addurl", addurl).Methods("POST")
	Rmux.HandleFunc("/api/v1/allurl", allurl).Methods("GET")
	return Rmux
}

func addurl(w http.ResponseWriter, r *http.Request) {
	urls := []model.Redirect{}
	json.NewDecoder(r.Body).Decode(&urls)
	var err error
	for _, url := range urls {
		err = db.AddRedirect(url)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, r)
			return
		}
	}
	if err == nil {
		fmt.Fprint(w, errors.New("No erros, all urls inserted"))
	}
}
func allurl(w http.ResponseWriter, r *http.Request) {
	res, err := db.GetAllUrls()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
	jso, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
	w.Write(jso)
}
