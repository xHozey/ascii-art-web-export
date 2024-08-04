package ascii_art_web

import (
	"net/http"
	"text/template"
)

// NotTheRootDirectory verifies if the request URL is for the home page.
func NotTheRootDirectory(path string) bool {
	return path != "/"
}

// NotPostMeth checks if the method of the https request
// isn't a POST method.
func NotPostMethod(method string) bool {
	return method != "POST"
}

// NotPostMeth checks if the method of the https request
// isn't a GET method.
func NotGetMethod(method string) bool {
	return method != "GET"
}

// ServerErrorOccured checks for server errors.
func ServerErrorOccured(err error) bool {
	return err != nil
}

// LargeInputSize verifies the length of the input string.
func LargeInputSize(input string) bool {
	return len(input) >= 2000
}

// ErrorHandler renders an error page with HTTP status details.
func ErrorHandler(w http.ResponseWriter, status int) {
	t, err := template.ParseFiles("templates/error.html")
	if ServerErrorOccured(err) {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	error := struct {
		Statustext string
		Statuscode int
		Message    string
	}{
		Statustext: http.StatusText(status),
		Statuscode: status,
	}
	w.WriteHeader(status)
	err = t.Execute(w, error)
	if ServerErrorOccured(err) {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
