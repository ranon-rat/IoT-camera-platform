package main

import (
	"encoding/json"
	"log"
	"net/http"
)
func registerUser(w http.ResponseWriter, r *http.Request){
	
	switch r.Method {
	case "POST":
	
	var newUser register
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.IP=r.Header.Get("x-forwarded-for")
	errChan:=make(chan error)
	
	go registerUserCameraDatabase(newUser,errChan)
	if <-errChan!=nil{
		log.Println(<-errChan)
		return
	}
	break

	default :
	/*
	
	*/
	w.Write([]byte("you cant do that 😡"))
	break
	}
}


func loginUser(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "POST":
	var newUser register
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.IP=r.Header.Get("x-forwarded-for")
	
	break
	default:
	w.Write([]byte("you cant do that 😡"))
	break

	}
}
