package access

import ("encoding/json"
        )



type Access_Request struct{

    Door string `json:"door"`
    Who string  `json:"who"`
    Date string  `json:"time,int"`
    House int `json:"house"`
}

type Access struct{

    ID 	int     `json:"idaccess"`    //indicano come sono chiamti i campi nel json
    Door string	`json:"door"`
    Who string	`json:"who"`
    Date string  `json:"time,int"`
    House int   `json:"id_house"`

}


func (a *Access) Create (id int,door string, who string, date string,ida int) {
	a.ID=id
	a.Door=door
	a.Who=who
    a.Date = date
    a.House = ida
}

func CreateFromJSON(jsonRequest string) (Access,error) {
	
	var n Access

	err := json.Unmarshal([]byte(jsonRequest), &n) //parso il json

    return n,err


}

func CreateFromJSONRequest(jsonRequest string) (Access,error) {
    
    var n Access
    var a Access_Request

    err := json.Unmarshal([]byte(jsonRequest), &a) //parso il json

    if(err==nil){
        n.Create(0,a.Door,a.Who,a.Date,a.House)
    }

    return n,err
   

}



func (n *Access) GetJSON() (string ,error){ //ritorni multipli

	 b, err := json.Marshal(n) //faccio il json, dichiarazione e assegnazione nello stesso istante con :=
       
     if err!=nil{
     	return "",err
     }
     	

     return string(b),err  
    	
}