package main

import (
	"encoding/json"
	"fmt"
	"github.com/goji/httpauth"
	"net/http"
	"os"
	"os/exec"
)

type SessionMessage struct {
	Token string `json:"token"`
}

type DeploymentMessage struct {
	ServiceName string `json:"service-name"`
	ImagePath   string `json:"image-path"`
}

func createToken(sm *SessionMessage) string {
	// app := "./create-token.sh"
	app := "/opt/deployer/create-token.sh"

	arg1 := sm.Token

	cmd := exec.Command(app, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		return string(err.Error())
	}

	return string(stdout)
}

func deploy(dm *DeploymentMessage) string {
	// app := "./restart-service.sh"
	app := "/opt/deployer/restart-service.sh"

	arg1 := dm.ServiceName
	arg2 := dm.ImagePath

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

func SessionHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var sm SessionMessage
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&sm)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, createToken(&sm))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DeploymentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var dm DeploymentMessage
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&dm)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, deploy(&dm))
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.Handle("/__deployer/health", 		httpauth.SimpleBasicAuth(os.Getenv("DEPLOYER_USER"), os.Getenv("DEPLOYER_PASSWORD"))(http.HandlerFunc(HealthHandler)))
	http.Handle("/__deployer/session", 		httpauth.SimpleBasicAuth(os.Getenv("DEPLOYER_USER"), os.Getenv("DEPLOYER_PASSWORD"))(http.HandlerFunc(SessionHandler)))
	http.Handle("/__deployer/deployment", httpauth.SimpleBasicAuth(os.Getenv("DEPLOYER_USER"), os.Getenv("DEPLOYER_PASSWORD"))(http.HandlerFunc(DeploymentHandler)))
	http.ListenAndServe(":7000", nil)
}
