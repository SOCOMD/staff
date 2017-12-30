package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/SOCOMD/staff/go/db/user"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening on localhost:8080")
	user.GetAll()
	user.GetSingle(1)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./website/dist/"))
	if strings.Contains(r.URL.String(), ".") == false {
		r.URL.Path = "/"
	}
	fmt.Println(r.URL.String())
	fs.ServeHTTP(w, r)
}
