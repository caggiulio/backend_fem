package main




import (
		"fmt"
    	"net/http"
    	"server"
        "utils"
        "os")


var mConf utils.Configuration

func main(){
	fmt.Println("Hello World!");


    mConf = utils.LoadConfiguration()
    if (mConf.Check()) {
        startListening()
    }else{
        utils.Log(utils.ERROR, "ProgettoFEM Service", "Error during Configuration load...")
        os.Exit(-1)
    }

}



func startListening(){



	http.HandleFunc("/api/hello", server.HandlerHW)
    fmt.Println("Serving on http://localhost:8008/hello")


    http.HandleFunc("/api/access", server.HandlerGetAccess)
    fmt.Println("Serving on http://localhost:8008/access")

    http.HandleFunc("/api/access/new", server.HandlerSaveAccess)
    fmt.Println("Serving on http://localhost:8008/access/new")


    err := http.ListenAndServe(mConf.Address+":"+mConf.Port, nil) 
        if (err != nil) {
            utils.Log(utils.ERROR, "ProgettoFEM Service", "Error on ListenAndServe: "+err.Error())
        }
}