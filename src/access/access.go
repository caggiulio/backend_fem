package access

import ("encoding/json"
        "time")


type Access struct{

    ID 	int     `json:"id"`    //indicano come sono chiamti i campi nel json
    Door string	`json:"door"`
    Who string	`json:"who"`
    Date time.Time   `json:"time"`

}


func (a *Access) Create (id int,door string, who string, date time.Time) {
	a.ID=id
	a.Door=door
	a.Who=who
    a.Date = date	
}

func CreateFromJSON(jsonRequest string) (Access,error) {
	
	var n Access

	//err := json.Unmarshal([]byte(jsonRequest), &n) //parso il json

    //return n,err
    return n,nil

}



func (n *Access) GetJSON() (string ,error){ //ritorni multipli

	 b, err := json.Marshal(n) //faccio il json, dichiarazione e assegnazione nello stesso istante con :=
       
     if err!=nil{
     	return "",err
     }
     	

     return string(b),err  
    	
}