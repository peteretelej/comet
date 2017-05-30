package ice

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Verbose speficies if logging should be noisy
var Verbose bool

// Serve launches a comet ice server
func Serve(listenAddr, dir string) error {
	var svr *http.Server
	switch dir {
	case "":
		svr = Server(listenAddr)
		dir = "default comet app"
	default:
		index := filepath.Join(dir, "index.html")
		if _, err := os.Stat(index); err != nil {
			log.Fatalf("cannot launch http fileserver, no index file: %s", index)
		}
		svr = DirServe(listenAddr, dir)
	}
	if Verbose {
		log.Printf("launching comet server at %s, serving '%s'", listenAddr, dir)
	}

	return svr.ListenAndServe()
}

// Server launches the comet http server
func Server(listenAddr string) *http.Server {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", "")
	})
	svr := &http.Server{
		Addr:           listenAddr,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return svr
}

// DirServe launches an http file server at the listenAddr serving the directory specified
func DirServe(listenAddr, dir string) *http.Server {
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	svr := &http.Server{
		Addr:           listenAddr,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return svr
}

func renderTemplate(w http.ResponseWriter, page string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.Header().Add("X-Frame-Options", "SAMEORIGIN")
	w.Header().Add("X-UA-Compatible", "IE=edge")

	err := defaultTmpl.ExecuteTemplate(w, page, data)
	if err != nil {
		log.Print(err.Error())
	}
}

var defaultTmpl = template.Must(template.New("tmpl").Parse(tmplHTML))

const tmplHTML = `{{define "head"}}<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Comet</title>
<meta name="viewport" content="width=device-width,initial-scale=1">
<style>
body{
	display:block; width:100%; overflow:hidden;
	background-color: #efefef;
}
</style>
</head>
<body>
{{end}}
{{define "foot"}}
</body>
</html>
{{end}}
{{define "index"}}{{template "head" .}}
<h1>Hello from Comet</h1>
<h4>Build Desktop Apps in Electron, Go, Bootstrap and Vuejs</h4>
<p style="text-align:center; margin:30px; ">
<iframe style="display:inline-block;" width="560" height="315" 
src="https://www.youtube.com/embed/nKIu9yen5nc?list=PLuKvd2GQmvCBF1YOhgkdGnbFQKnvuXHSK" 
frameborder="0" allowfullscreen></iframe>
</p>
{{template "foot" .}}
{{end}}
`
