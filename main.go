package main

import (
	"net/http"
	"proxy/router"
)

func main() {
	http.ListenAndServe(":80", router.NewMux())
}
