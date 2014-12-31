package main

import (
 //    "net"
	// "net/http"
	// "log"
	// "net/url"
	// // "text/template"
	// "time"
	// "fmt"
)

type Notificator struct {
	port int
}

func NewNotificator() (notificator Notificator) {
	notificator = Notificator{8080}
	return
	// rootUrl := "localhost:1880"
	// // url := "ws://localhost:1880/api/infusion"
	// var url, err = url.Parse(url)

	// if err != nil {
	// 	log.Fatal(err)
 //    }

	// conn, err := net.Dial("tcp", rootUrl)
	// handle(err)

	// if err != nil {
	// 	log.Fatal(err)
 //    }

	// wsConn, _, err := websocket.NewClient(conn, url, http.Header{}, 1024, 1024)
	// if err != nil {
	// 	log.Fatal(err)
 //    }

 //    message := "{\"dropped\": 10}"

 //    c := time.Tick(1 * time.Second)
 //    for now := range c {
 //        err = websocket.WriteJSON(wsConn, message)
 //        if err != nil {
 //            log.Fatal(err)
 //        }    	
 //        fmt.Printf("[%v] data is sent.\n", now)
 //    }	

}



