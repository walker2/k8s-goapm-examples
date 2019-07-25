package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgorilla"
)

func roll20() int {
	return rand.Intn(20) + 1
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Handling connection right now")
	if roll20() > 18 {
		ctx := req.Context()
		span, ctx := apm.StartSpan(ctx, "WAIT FOR IT", "wait")
		time.Sleep(2 * time.Second)
		span.End()

		span, ctx = apm.StartSpan(ctx, "WAIT FOR LITTLE MORE", "wait")
		time.Sleep(2 * time.Second)
		span.End()
	}
}

func main() {
	log.Println("Started hello service")
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)
	r.Use(apmgorilla.Middleware())
	log.Fatal(http.ListenAndServe(":8000", r))
}
