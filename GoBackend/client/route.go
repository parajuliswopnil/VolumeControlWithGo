package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
	"path"
)

type Volume struct{
	Volume string
}

type VolumeData struct {
	VolumeInPercentage string
}

func NoTrailingSlash(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

var Serve = NoTrailingSlash(serve)

func serve(w http.ResponseWriter, r *http.Request){
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	switch head {
	case "":
		serveHome(w, r)
	case "changeVolume":
		fmt.Println("reached change volume")
		id, err := strconv.Atoi(head)
		if err != nil{
			http.NotFound(w, r)
		}

		Volume{string(id)}.serveHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request){
	fmt.Println("Welcome home")
}

func (vol Volume) serveHTTP(w http.ResponseWriter, r *http.Request){
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)

	var reply VolumeData
	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("error in connection")
	}

	volume := VolumeData{VolumeInPercentage: head}
	client.Call("API.SetVolume", volume, &reply)
}

func main() {

	router := http.HandlerFunc(Serve)
	log.Printf("listening on port %s", ":8000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8000"), router))
}

