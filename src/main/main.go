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


    var Backend server.FEMbackend

    Backend.Initialize()


	http.HandleFunc("/api/hello", Backend.HandlerHW)
    utils.Log(utils.ASSERT, "ProgettoFEM Service", "Serving on "+mConf.Address+":"+mConf.Port+"/api/hello")


    http.HandleFunc("/api/access", Backend.HandlerGetAccess)
    utils.Log(utils.ASSERT, "ProgettoFEM Service", "Serving on "+mConf.Address+":"+mConf.Port+"/api/access")

    http.HandleFunc("/api/access/new", Backend.HandlerSaveAccess)
    utils.Log(utils.ASSERT, "ProgettoFEM Service", "Serving on "+mConf.Address+":"+mConf.Port+"/api/access/new")

    http.HandleFunc("/api/alarm/new", Backend.HandlerSaveAlarm)
    utils.Log(utils.ASSERT, "ProgettoFEM Service", "Serving on "+mConf.Address+":"+mConf.Port+"/api/alarm/new")

    http.HandleFunc("/api/exit/new", Backend.HandlerSaveExit)
    utils.Log(utils.ASSERT, "ProgettoFEM Service", "Serving on "+mConf.Address+":"+mConf.Port+"/api/exit/new")


    http.HandleFunc("/api/user/new", Backend.HandlerSaveUser)
    utils.Log(utils.ASSERT, "ProgettoFEM Service", "Serving on "+mConf.Address+":"+mConf.Port+"/api/user/new")


    err := http.ListenAndServe(mConf.Address+":"+mConf.Port, nil) 
        if (err != nil) {
            utils.Log(utils.ERROR, "ProgettoFEM Service", "Error on ListenAndServe: "+err.Error())
        }
}