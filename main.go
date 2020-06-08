package main

import (
	"fmt"
	"ges"
	"net/http"
)

func main() {
	r := ges.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.RUN(":9999")
}
