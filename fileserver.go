package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", cors(fs))

	PORT, err := getPortFromArgs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	message := fmt.Sprintf("server is listening on port %s", PORT)
	fmt.Println(message)

	http.ListenAndServe(PORT, nil)
}

func getPortFromArgs() (string, error) {
	var err error
	args := os.Args
	if len(args) < 2 {
		err = errors.New("please pass a port like so `--port=4200`")
		return "", err
	}

	firstArg := args[1:2]
	portArgument := strings.Split(firstArg[0], "=")
	if portArgument[0] != "--port" || len(portArgument) < 2 {
		err = errors.New("please pass a port like so `--port=4200`")
		return "", err
	}

	port := portArgument[1]
	if len(port) != 4 {
		err = errors.New("please pass a 4 character port like `4200`")
		return "", err
	}

	return ":" + port, err
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		next.ServeHTTP(w, r)
	})
}
