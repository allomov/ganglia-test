package main

import (
	// "os"
	// "strconv"
	// "log"
	"fmt"
)

func main() {
	// var port int64
	// var err error
	// portStr := os.Getenv("PORT")
	// if portStr != "" {
	// 	port, err = strconv.ParseInt(portStr, 10, 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	port = 8080
	// }

	fmt.Println("Creating IV pump.")
    ivpump := NewInfusionPump()
	fmt.Println("Creating Server.")
    server := NewServer(&ivpump)
    
    // n := notificator.NewNotificator(ivpump, "localhost:1880")
    fmt.Println("Starting Server.")
    server.Start()
    // go n.Start()
}
