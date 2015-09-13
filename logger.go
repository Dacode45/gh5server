package main

import (
	"log"
	"net/http"
	"time"
	"errors"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Add("Access-Control-Allow-Origin", "http://ayeke.me:3000")
		var err error
			defer func() {
					r := recover()
					if r != nil {
							switch t := r.(type) {
							case string:
									err = errors.New(t)
							case error:
									err = t
							default:
									err = errors.New("Unknown error")
							}
							http.Error(w, err.Error(), http.StatusInternalServerError)
					}
			}()
		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
