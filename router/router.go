package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"proxy/db"
	"proxy/model"

	"github.com/gorilla/mux"
)

func NewMux() http.Handler {
	addu := http.HandlerFunc(addurl)
	getu := http.HandlerFunc(allurl)
	Rmux := mux.NewRouter()
	Rmux.Handle("/api/v1/addurl", loggerMiddleware(addu)).Methods("POST")
	Rmux.Handle("/api/v1/allurl", loggerMiddleware(getu)).Methods("GET")
	Rmux.PathPrefix("/").Handler(loggerMiddleware(http.FileServer(http.Dir("/root/go/src/proxy/html/"))))
	return Rmux
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s [%s] %s [%s] [%s]", r.RemoteAddr, r.Method, r.Header.Get("User-Agent"), r.URL, r.Host)
		next.ServeHTTP(w, r)
	})
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
