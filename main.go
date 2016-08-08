package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var port int

func init() {
	flag.IntVar(&port, "port", 4321, "Port to listen on")
}

func main() {
	done := make(chan bool, 1)

	flag.Parse()

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK\n"))
		})

		http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
		done <- true
	}()

	args := flag.Args()

	if len(args) > 0 {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		<-done
	}
}
