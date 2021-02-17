package main

import (
	"encoding/json"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		var newUser registerCamera
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.IP = r.Header.Get("x-forwarded-for")
		errChan := make(chan bool)
		go registerUserCameraDatabase(newUser, errChan)
		if len(newUser.Username) == 0 || len(newUser.Password) == 0 {
			http.Error(w, "your password or your username is empty", 406)
			return
		}
		sizeChan := make(chan int, 1)
		// creo que esto deberia de marcarnos si sizeChan esta cerrado o no
		// asi creo que PODRIAMOS manejar  los errores

		// we check if the username of the camera already exist
		go exist(newUser.Username, *encryptData(newUser.IP), sizeChan)

		if <-sizeChan > 0 {
			http.Error(w, "that user has been already registered", 409)
			return // manage the errors
		}
		if <-errChan {

			http.Error(w, "something is bad , try again in other moment", 502)
			return
		}
		w.Write([]byte("now you are registered "))

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
		errChan, valid, token := make(chan bool), make(chan bool), make(chan string)
		// check if all is okay

		go loginUserCameraDatabase(oldUser, valid)

		if <-valid {
			go updateUsages(oldUser, errChan) // we update the last time that he send something
			go generateToken(oldUser, token, errChan)
			if !<-errChan {
				http.Error(w, "something is bad,try again in other moment ", 502)
				return
			}
			w.Write([]byte(<-token))

			return
		}
		http.Error(w, "something is wrong , verify your password, or user\n ", 502)
		// generate the token

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
