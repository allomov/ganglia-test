package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/ghost/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


type Server struct {
	http *http.Server
}

func (s *Server) Start() {
	fmt.Println("Server started.")
	log.Panic(s.http.ListenAndServe())
}

func NewServer(ivpump *InfusionPump) (*Server) {
    fmt.Println("Let's Start Server!")

	r := mux.NewRouter()

	r.PathPrefix("/scripts/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/styles/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/images/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/fonts/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/ico/").Handler(http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/favicon.ico").Handler(http.FileServer(http.Dir("./static/")))
	
	f := func(w http.ResponseWriter, r *http.Request) {
    	http.ServeFile(w, r, "/Users/allomov/work/altoros/node-red/iv-simulator/src/github.com/allomov/iv-simulator/static/index.html")
	}

	r.HandleFunc("/", f)
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	port := flag.String("port", "8080", "port number, default: 8080")
	flag.Parse()

	log.Printf("Ifconfig server started. :%v", *port)
	log.Printf("---------------------------")

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: handlers.PanicHandler(handlers.LogHandler(r, handlers.NewLogOptions(log.Printf, "_default_")), nil),
	}

	return &Server{s}
}
