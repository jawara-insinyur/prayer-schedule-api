package handler

import (
	"log"
	"net/http"
)

type RootHandler func(w http.ResponseWriter, r *http.Request) error

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err == nil {
		log.Printf("[success request] %v", r.URL)
		return
	}

	e, ok := err.(clientError)
	if !ok {
		log.Println("[server error] cause: Failed to casting error to clientError")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := e.ResponseBody()
	if err != nil {
		log.Printf("[server error] cause: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status, header := e.ResponseHeader()
	for k, v := range header {
		w.Header().Set(k, v)
	}

	log.Printf("[client error] cause: %v", e)
	w.WriteHeader(status)
	w.Write(body)
}
