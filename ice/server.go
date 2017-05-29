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
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css" integrity="sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ" crossorigin="anonymous">
<script src="https://code.jquery.com/jquery-3.1.1.slim.min.js" integrity="sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js" integrity="sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb" crossorigin="anonymous"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js" integrity="sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn" crossorigin="anonymous"></script>
</head>
<body>
{{end}}
{{define "foot"}}
</body>
</html>
{{end}}
{{define "index"}}{{template "head" .}}
<div class="jumbotron jumbotron-fluid" style="height:100%">
<div class="container">
<h3>Hello from Comet</h3>
<p class="lead">Build Desktop Apps in Electron, Go, Bootstrap and Vuejs</p>
<p class="text-center">
<iframe width="560" height="315" src="https://www.youtube.com/embed/nKIu9yen5nc?list=PLuKvd2GQmvCBF1YOhgkdGnbFQKnvuXHSK" frameborder="0" allowfullscreen></iframe>
</p>

</div><!--/.container-->
</div><!--/.jumbotron-->
{{template "foot" .}}
{{end}}
`
