package main

import (
	"encoding/json"
	"fmt"
	"github.com/goji/httpauth"
	"net/http"
	"os"
	"os/exec"
)

type Message struct {
	ServiceName string `json:"service-name"`
	ImagePath   string `json:"image-path"`
}

func runCmd(t *Message) string {
	// app := "./restart-service.sh"
	app := "/opt/deployer/restart-service.sh"

	arg1 := t.ServiceName
	arg2 := t.ImagePath

	cmd := exec.Command(app, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		return string(err.Error())
	}

	return string(stdout)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK!")
}

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		decoder := json.NewDecoder(req.Body)
		var t Message
		err := decoder.Decode(&t)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, runCmd(&t))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.Handle("/__deployer/health", httpauth.SimpleBasicAuth(os.Getenv("DEPLOYER_USER"), os.Getenv("DEPLOYER_PASSWORD"))(http.HandlerFunc(HealthHandler)))
	http.Handle("/__deployer/deploy", httpauth.SimpleBasicAuth(os.Getenv("DEPLOYER_USER"), os.Getenv("DEPLOYER_PASSWORD"))(http.HandlerFunc(MainHandler)))
	http.ListenAndServe(":7000", nil)
}
