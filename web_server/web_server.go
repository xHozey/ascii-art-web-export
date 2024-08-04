package ascii_art_web

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	art "ascii_art_web/ascii_art"
)

var Generated string

// Index handles requests to the root URL ("/").
// It renders the main index template.
func Index(w http.ResponseWriter, r *http.Request) {
	if NotGetMethod(r.Method) {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	if NotTheRootDirectory(r.URL.Path) {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if ServerErrorOccured(err) {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	execErr := tmpl.Execute(w, nil)
	if ServerErrorOccured(execErr) {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

// Generate handles POST requests to generate ASCII art.
// It renders the ASCII art template based on user input.
func Generate(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("text")
	banner := r.FormValue("banner")

	if NotPostMethod(r.Method) {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/ascii-art.html")
	if ServerErrorOccured(err) || LargeInputSize(input) {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if !validBannerName(banner) {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	input = filterInput(input)
	bannerPath := "ascii_art/banners/" + banner + ".txt"
	Generated, err = art.ArgProcessor(input, bannerPath)
	if ServerErrorOccured(err) {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	execErr := tmpl.Execute(w, Generated)
	if ServerErrorOccured(execErr) {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func Download(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Content-Disposition", "attachment; filename=Generated.txt")
	w.Header().Add("Content-Length", fmt.Sprintln(len(Generated)))
	w.Write([]byte(Generated))

}

// NotFoundHandler checks if the request path ends with "/" and handles accordingly.
func NotFoundHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			ErrorHandler(w, http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
