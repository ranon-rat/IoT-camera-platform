package main

import (
	"encoding/json"
	"net/http"
)
func registerUser(w http.ResponseWriter, r *http.Request){
	var newUser register
	switch r.Method {
	case "POST":
		json.NewDecoder(r.Body).Decode(&newUser)
		break

		
	default :
		break;
	}
	

}
