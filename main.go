package main

// Importing the libraries
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ibilalkayy/Portfolio/routes"
)

// Execute calls the routes and sets the port
func Execute() error {
	routes.Route()
	fmt.Println("Starting the server at port: 4000")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}
	err := http.ListenAndServe(":"+port, nil)
	return err
}

// main sets the error message for the Execute() function
func main() {
	if err := Execute(); err != nil {
		log.Fatal(err)
	}
}
