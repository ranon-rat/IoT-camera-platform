package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		var newUser registerCamera
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.IP = r.Header.Get("x-forwarded-for")
		mes := make(chan codeHTTP)
		go registerUserCameraDatabase(newUser, mes)
		c := <-mes
		w.Write([]byte(fmt.Sprintf("%s %d", c.Message, c.Code)))

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
		codeMessage, valid, token := make(chan codeHTTP), make(chan bool), make(chan string)
		// check if all is okay

		go loginUserCameraDatabase(oldUser, codeMessage, valid)
		c := <-codeMessage

		if <-valid {

			go updateUsages(oldUser, codeMessage)         // we update the last time that he send something
			go generateToken(oldUser, codeMessage, token) // generate the token
			w.Write([]byte(<-token))
			return
		}
		w.Write([]byte(fmt.Sprintf("%s %d", c.Message, c.Code)))

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
