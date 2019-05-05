// +build ignore

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	wd, _ := os.Getwd()
	l := "0.0.0.0:8844"
	log.Printf("Starting HTTP Server at %s\n", l)
	h := simplehttp.New(wd, true)
	log.Fatal(http.ListenAndServe(l, h))
}
