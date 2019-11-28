package web

import "net/http"

func ShortenHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("gago"))
}
