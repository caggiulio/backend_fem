package server

import(
"fmt"
"strconv"
"net/http"
"access"
"time"
) 


func HandlerHW(w http.ResponseWriter, r *http.Request) { //esportata inizia per maiuscola

	printResult(w,"Hello World!")

}



func HandlerGetAccess(w http.ResponseWriter, r *http.Request) { //esportata inizia per maiuscola

	var n access.Access


	if (r.Method == "GET"){
		fmt.Println("HTTP Method is correct")

	n.Create(3,"Garage Door","Salvatore",time.Now())



    s,err:=n.GetJSON() //ignoro un ritorno con _
    if err!=nil{
    	fmt.Println("error %s",err)
    }


    printResult(w,s)
    }else{
    	fmt.Println("HTTP Method is wrong")
    }
}

func HandlerSaveAccess(w http.ResponseWriter, r *http.Request) { //esportata inizia per maiuscola

	//var n access.Access
	var err error


	if (r.Method == "POST"){
		fmt.Println("HTTP Method is correct")


		//{"title":"Go is stunning!!","sub":"Go http package features","content":"Great http services with go and is awesome http package"}

		req := r.FormValue("req") //leggo il parametro

		_,err=access.CreateFromJSON(req)
		if err!=nil{
			fmt.Println("error %s",err)
		}else{
			printResult(w,"{res:true}")
			//printResult(w,n.Title + " \n" + n.SubTitle + " \n" + n.Content)
		}

	}else{
		fmt.Println("HTTP Method is wrong")
		printResult(w,"404")
			//si potrebbe ridirezionare ad un errore
	}


}



func printResult(w http.ResponseWriter,resp string){ //non esportata inizia per minuscola


	fmt.Fprintf(w, "%s",resp);
}

func printJSON(w http.ResponseWriter,resp string){


	fmt.Fprintf(w, "%s","{"+strconv.Quote("Res")+":"+resp+"}")
}
