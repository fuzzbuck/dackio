package comms

import (
	"flag"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")


func StartWSHub() {
	flag.Parse()
	hub := newHub()
	GlobalHub = hub // "game" hub

	go hub.run()
	http.HandleFunc("/sock", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "comms_test.html")
	})
	go func() {
		err := http.ListenAndServe(*addr, nil)
		if err != nil {
			panic(err)
		}
	}()
}