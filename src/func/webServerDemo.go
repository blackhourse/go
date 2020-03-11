package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w,"11111")

	bytes, _ := ioutil.ReadFile("./src/func/index.html")
	w.Write(bytes)

}
