package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func goroutineID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("../templates/index.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/registrar", func(w http.ResponseWriter, r *http.Request) {
		tipo := r.FormValue("tipo")
		timestamp := time.Now().Format(time.RFC3339)
		gid := goroutineID()

		switch tipo {
		case "Info":
			color.Cyan("[INFO][%s][Goroutine %s] - Log registrado", timestamp, gid)
		case "Error":
			color.Red("[ERROR][%s][Goroutine %s] - Log registrado", timestamp, gid)
		case "Warning":
			color.Yellow("[WARNING][%s][Goroutine %s] - Log registrado", timestamp, gid)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	color.Cyan("Iniciando el servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
