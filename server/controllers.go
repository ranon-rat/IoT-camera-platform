package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		var newUser registerCamera
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.IP = r.Header.Get("x-forwarded-for")
		errChan := make(chan error)

		go registerUserCameraDatabase(newUser, w)
		if <-errChan != nil {
			log.Println(<-errChan)
			err := <-errChan
			w.Write([]byte(err.Error()))
			return
		}
		break

	default:
		http.Error(w, "you cant do that 😡", 405)
		break
	}
}

func loginUserCamera(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// we setup the values from the struct registerCamera
		var oldUser registerCamera
		json.NewDecoder(r.Body).Decode(&oldUser)
		oldUser.IP = r.Header.Get("x-forwarded-for")
		//  we use this for asynchronous communication
		valid, token := make(chan bool), make(chan string)
		// check if all is okay
		go loginUserCameraDatabase(oldUser, w, valid)
		if <-valid {

			go updateUsages(oldUser, w)         // we update the last time that he send something
			go generateToken(oldUser, w, token) // generate the token
			// and send the token
			w.Write([]byte(<-token))

			return
		}
		break
	default:
		http.Error(w, "you cant do that 😡", 405)
		break

	}
}

/*
func controlData(conn *websocket.Conn, user registerCamera) {
	videoCamera[user.Username] = defaultImage
	for{
		message,m,err:=conn.ReadMessage()
		if err!=nil{
			log.Println(err)
			break
		}


	}
}
*/
