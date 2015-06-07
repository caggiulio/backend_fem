package server

import(
"fmt"
"strconv"
"net/http"
"access"
"time"
"user"
"utils"
"dbhelper"
) 



type FEMbackend struct{
	mConfiguration utils.Configuration
	mDBHelper dbhelper.DBHelper

}

func (back *FEMbackend) Initialize() {
	back.mConfiguration=utils.LoadConfiguration()
	back.mDBHelper.SetConfiguration(back.mConfiguration)
}



func (back FEMbackend) HandlerHW(w http.ResponseWriter, r *http.Request) { //esportata inizia per maiuscola

	printResult(w,"Hello World!")

}



func (back FEMbackend) HandlerGetAccess(w http.ResponseWriter, r *http.Request) { //esportata inizia per maiuscola

	var n access.Access


	if (r.Method == "GET"){
		fmt.Println("HTTP Method is correct")
		utils.Log(utils.ASSERT, "ProgettoFEM Backend", "HTTP Method is correct")

		n.Create(3,"Garage Door","Salvatore",strconv.FormatInt(time.Now().Unix(),10),9)



    s,err:=n.GetJSON() //ignoro un ritorno con _
    if err!=nil{
    	fmt.Println("error %s",err)
    }


    printResult(w,s)
}else{
	fmt.Println("HTTP Method is wrong")
	utils.Log(utils.WARNING, "ProgettoFEM Backend", "HTTP Method is wrong")
}
}

func (back FEMbackend) HandlerSaveAccess(w http.ResponseWriter, r *http.Request) { //esportata inizia per maiuscola

	var n access.Access
	var err error


	if (r.Method == "POST"){
		fmt.Println("HTTP Method is correct")
		utils.Log(utils.ASSERT, "ProgettoFEM Backend", "HTTP Method is correct")


		//{"door":"Garage Door","who":"Salvatore","time":"1433697822","idhouse":3}
		req := r.FormValue("req") //leggo il parametro
		fmt.Println(req)

		n,err=access.CreateFromJSONRequest(req)
		if err!=nil{
			fmt.Println("error %s",err)




		}else{
			


			coloumns := make([]string, 0, 0)
			values := make([]string, 0, 0)


			coloumns = append(coloumns, "door")
			coloumns = append(coloumns, "time")
			coloumns = append(coloumns, "id_house")
			coloumns = append(coloumns, "who")
			
			values = append(values, n.Door)
			values = append(values, n.Date)
			values = append(values, strconv.FormatInt(int64(n.House), 10))
			values = append(values, n.Who)

			r:=back.mDBHelper.Insert("access",coloumns,values)
			utils.Log(utils.ASSERT, "ProgettoFEM Backend", r)
			fmt.Println(r)

			printResult(w,"{res:true}")


			//printResult(w,n.Title + " \n" + n.SubTitle + " \n" + n.Content)
		}

	}else{
		fmt.Println("HTTP Method is wrong")
		utils.Log(utils.WARNING, "ProgettoFEM Backend", "HTTP Method is wrong")
		printResult(w,"404")
			//si potrebbe ridirezionare ad un errore
	}


}


func (back FEMbackend) HandlerSaveUser(w http.ResponseWriter, r *http.Request) { //esportata inizia per maiuscola

	
	
	var u user.User


	if (r.Method == "POST"){
		fmt.Println("HTTP Method is correct")
		utils.Log(utils.ASSERT, "ProgettoFEM Backend", "HTTP Method is correct")


		//{"title":"Go is stunning!!","sub":"Go http package features","content":"Great http services with go and is awesome http package"}

		req := r.FormValue("req") //leggo il parametro
		fmt.Println(req)

		u.Create(0,"test","1234")
		


		

		coloumns := make([]string, 0, 0)
		values := make([]string, 0, 0)


		coloumns = append(coloumns, "name")
		coloumns = append(coloumns, "registrationid")
		
		values = append(values, u.Name)
		values = append(values, u.RegID)

		r:=back.mDBHelper.Insert("users",coloumns,values)
		utils.Log(utils.ASSERT, "ProgettoFEM Backend", r)
		fmt.Println(r)

		printResult(w,"{res:true}")
		

	}else{
		fmt.Println("HTTP Method is wrong")
		utils.Log(utils.WARNING, "ProgettoFEM Backend", "HTTP Method is wrong")
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
