/**
		GDG GoLang Backend Daemon
        Copyright (C) 2014+  Gabriele Baldoni
 */

package main //nome package

import ( //package importati
	"fmt"    
    "os"
)

func start() { //avvia un nuovo processo
   var procAttr os.ProcAttr 
   procAttr.Files = []*os.File{nil, nil, nil}
   _, err := os.StartProcess("/usr/bin/fembackend", nil, &procAttr)
   if err != nil {
       fmt.Printf("%v", err)
   }
}

func main() { //main dell'app

	start()

}
