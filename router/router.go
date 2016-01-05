package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proxy/db"

	"github.com/gorilla/mux"
)

func NewMux() http.Handler {
	Rmux := mux.NewRouter()
	Rmux.HandleFunc("/api/v1/addurl", addurl).Methods("POST")
	Rmux.HandleFunc("/api/v1/allurl", allurl).Methods("GET")
	return Rmux
}

func addurl(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "ADD New URL")
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
