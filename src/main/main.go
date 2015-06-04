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
	http.HandleFunc("/hello", server.HandlerHW)
    fmt.Println("Serving on http://192.168.1.4:8888/hello")


    http.HandleFunc("/access", server.HandlerGetAccess)
    fmt.Println("Serving on http://192.168.1.4:8888/access")

    http.HandleFunc("/access/new", server.HandlerSaveAccess)
    fmt.Println("Serving on http://localhost:8888/access/new")


    log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}