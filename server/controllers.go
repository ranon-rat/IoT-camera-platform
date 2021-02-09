package main

import (
	"encoding/json"
	"net/http"
)
func registerUser(w http.ResponseWriter, r *http.Request){
	
	switch r.Method {
	case "POST":
	var newUser register
	json.NewDecoder(r.Body).Decode(&newUser)

	break
	default :
	
	w.Write([]byte("you cant do that 😡"))
	break
	}
	

}
