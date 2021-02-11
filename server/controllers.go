package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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
		go updateUsages(oldUser)
		upgrade.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		go controlData(ws)
		
		return
	}
	break
	default:
	w.Write([]byte("you cant do that 😡"))
	break

	}
}

func controlData(conn *websocket.Conn,user register){
	videoCamera[user.Username]=defaultImage
	/*for{
		message,m,err:=conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			break
		}
		

	}*/


}
