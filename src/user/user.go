package user

import ("encoding/json"
        )


type User struct{

    ID 	int     `json:"id"`    //indicano come sono chiamti i campi nel json
    Name string	`json:"name"`
    RegID string	`json:"registrationid"`

}


func (a *User) Create (id int,name string, regid string) {
	a.ID=id
	a.Name=name
	a.RegID=regid
}

func CreateFromJSON(jsonRequest string) (User,error) {
	
	var n User

	err := json.Unmarshal([]byte(jsonRequest), &n) //parso il json

    return n,err
}



func (n *User) GetJSON() (string ,error){ //ritorni multipli

	 b, err := json.Marshal(n) //faccio il json, dichiarazione e assegnazione nello stesso istante con :=
       
     if err!=nil{
     	return "",err
     }
     	

     return string(b),err  
    	
}