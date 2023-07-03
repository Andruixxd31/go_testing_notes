package httpservice

import (
	"fmt"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "Welcome to the HomePage")
    fmt.Println("Endpoint hit: HomePage")
}

func Setup(){
    http.HandleFunc("/", HomePage)
}

func main(){
    Setup()
    log.Fatal(http.ListenAndServe(":10000", nil))
}

