package ice

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var Verbose bool

var tmpl *template.Template

// Serve launches the comet http server
func Serve(listenAddr string) error {
	var err error
	tmpl, err = template.ParseGlob("tmpl/*.gohtml")
	if err != nil {
		return fmt.Errorf("unable to parse templates at tmpl/: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", "")
	})
	svr := &http.Server{
		Addr:           listenAddr,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if Verbose {
		log.Printf("launching http server at %s", listenAddr)
	}
	return svr.ListenAndServe()
}

func renderTemplate(w http.ResponseWriter, page string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.Header().Add("X-Frame-Options", "SAMEORIGIN")
	w.Header().Add("X-UA-Compatible", "IE=edge")

	err := tmpl.ExecuteTemplate(w, page, data)
	if err != nil {
		log.Print(err.Error())
	}
}
