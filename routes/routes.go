package routes

// Importing the libraries
import (
	"net/http"

	"github.com/ibilalkayy/Portfolio/controllers"
	"github.com/ibilalkayy/Portfolio/middleware"
)

// Route handles the paths that should be visited
// It uses middleware to check for errors
func Route() {
	http.HandleFunc("/", middleware.ErrorHandling(controllers.Template))
	// Enabling and adding CSS files
	fileServer := http.FileServer(http.Dir("./views/static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
}
