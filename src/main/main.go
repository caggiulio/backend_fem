package main




import (
		"fmt"
		"log"
    	"net/http"
    	"server")

func main(){
	fmt.Println("Hello World!");

	startListening();



}



func startListening(){
	http.HandleFunc("/api/hello", server.HandlerHW)
    fmt.Println("Serving on http://localhost:8008/hello")


    http.HandleFunc("/api/access", server.HandlerGetAccess)
    fmt.Println("Serving on http://localhost:8008/access")

    http.HandleFunc("/api/access/new", server.HandlerSaveAccess)
    fmt.Println("Serving on http://localhost:8008/access/new")


    log.Fatal(http.ListenAndServe("localhost:8008", nil)) //questa porta per permettere a selinux di far passare il servizio
}