/**
		GDG GoLang Backend Daemon
        Copyright (C) 2014+  Gabriele Baldoni
 */
package gcm

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"

)

const apiKey = ""
const gcmurl ="https://android.googleapis.com/gcm/send"

type GCMRequest struct{
	RegistrationIDs []string 	`json:"registration_ids"`
	CollapseKey string 			`json:"collapse_key,omitempty"`
	DelayWhileIdle bool 		`json:"delay_while_idle,omitempty"`
	Data map[string]string 		`json:"data,omitempty"`
	TimeToLive int 				`json:"time_to_live,omitempty"`
}





type GCMResponse struct {
	MulticastID int64 				`json:"multicast_id"`
	Success int 					`json:"success"`
	Failure int 					`json:"failure"`
	CanonicalIDs int 				`json:"canonical_ids"`
	Results []struct {
		MessageID string 		`json:"message_id"`
		RegistrationID string 	`json:"registration_id"`
		Error string 			`json:"error"`
		}							`json:"results"`

}

func SendGCMRequest(req GCMRequest) GCMResponse {

	var resp GCMResponse




	b, _ := json.Marshal(req)
	data:= string(b)

	client:=&http.Client{}


	r,_:=http.NewRequest("POST",gcmurl,bytes.NewBufferString(data))
	r.Header.Add("Authorization:" , "key=" + apiKey)
	r.Header.Add("Content-Type", "application/json")

	httpResp,_:=client.Do(r)

	fmt.Println(httpResp)

	return resp
}

func ParseGCMResponse(resp GCMResponse) {



}
