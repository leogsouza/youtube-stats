package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/leogsouza/youtube-stats/websocket"
)

// homePage will be a simple "hello World" style page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)
}

// setupRoutes handles setting up our servers
// routes and matching them to their respective
// functions
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	// here we kick off our server on localhost:8080
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// our main function
func main() {
	fmt.Println("YouTube Subscriber Monitor")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// calls setup routes
	setupRoutes()
}
