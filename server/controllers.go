package main

import (
	"encoding/json"
	"fmt"
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
		err:=<-errChan
		w.Write([]byte(err.Error()))
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


func loginUserCamera(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "POST":
	var oldUser register
	json.NewDecoder(r.Body).Decode(&oldUser)
	oldUser.IP=r.Header.Get("x-forwarded-for")
	valid:=make( chan bool )
	// check if all is okay
	go loginUserCameraDatabase(oldUser,valid)
	if <-valid{
		fmt.Println(oldUser.Username)
		videoCamera[oldUser.Username]=defaultImage
		
		fmt.Println(videoCamera)
		w.Write([]byte("you are already register"))
		return
	}
	break
	default:
	w.Write([]byte("you cant do that 😡"))
	break

	}
}
