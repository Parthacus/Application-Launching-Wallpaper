package main

import (
	"net/http"
	"os"
	"os/exec"
	"strings"
  "io/ioutil"
)

var addr = "localhost:8080"

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	if a := os.Getenv("ADDR"); a != "" {
		addr = a
	}

	http.HandleFunc("/", api)
	if err := http.ListenAndServe(addr, nil); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func api(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var command string

	if r.Method == "GET" || r.Method == "POST" {
		command = r.FormValue("command")
		if command == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("?command= cannot be empty"))
		}
	}

	switch r.Method {
	case "GET":
		path, err := exec.LookPath(strings.Fields(command)[0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(path))

	case "POST":
		var (
			parts   = strings.Fields(command)
			command = parts[0]
			args    = []string{}
		)

		if command == "read" {
			dat, err := ioutil.ReadFile(parts[1])
			check(err)
			w.Write([]byte(string(dat)))
		}


		if len(parts) > 1 {
			args = parts[1:]
		}

		cmd := exec.Command(command, args...)

		if err := cmd.Start(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
		w.Write([]byte("GET/POST only"))
	}
}
