package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const _httpPort = 8000

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	log.Printf("local netdata http server start and prot is %d \r\n", _httpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", _httpPort), mux)
	if err != nil {
		log.Fatalf("local netdata http server fail %v \r\n", err)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	cmd := exec.Command("/bin/bash", "/root/traffic.sh")
	out, err := cmd.CombinedOutput()
	log.Printf("%v", err)
	fmt.Fprintln(writer, string(out))
}
