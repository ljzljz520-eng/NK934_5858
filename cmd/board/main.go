package main

import (
	"flag"
	"log"
	"net/http"
	"waveboard/internal/api"
	"waveboard/internal/fixture"
	"waveboard/internal/service"
	"waveboard/internal/store"
)

func main() {
	path := flag.String("db", "waveboard.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	s, e := store.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	m := service.NewManager(s)
	if _, e = s.Wave("W-100"); e != nil {
		if e = fixture.Seed(m); e != nil {
			log.Fatal(e)
		}
	}
	log.Printf("warehouse board listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, api.New(m).Routes()))
}
