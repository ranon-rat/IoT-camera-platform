package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//******** CAMERA *******************
func registerUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		var newUser registerCamera
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.IP = r.Header.Get("x-forwarded-for")
		// if the password is empty send you this
		if len(newUser.Username) < 4 || len(newUser.Password) < 4 {
			http.Error(w, "your password or your username is empty or is less than 4 characters", 406)
			return
		}
		errChan, sizeChan := make(chan bool), make(chan int, 1)
		// creo que esto deberia de marcarnos si sizeChan esta cerrado o no
		// asi creo que PODRIAMOS manejar  los errores

		// we check if the username of the camera already exist
		go exist(newUser.Username, *encryptData(newUser.IP), sizeChan)
		if <-sizeChan > 0 {
			http.Error(w, "that user has been already registered", 409)
			return // manage the errors
		}
		//this register the user for the database
		go registerUserCameraDatabase(newUser, errChan)
		if <-errChan {
			http.Error(w, "internal server error", 500)
			return
		}
		//if everything is fine send you this
		w.Write([]byte("now you are registered "))

		break

	default:
		http.Error(w, "you cant do that 😡", 405)
		break
	}
}

// this is for login the user and send you that
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
		http.Error(w, "something is wrong , verify your password, or user ", 502)
		// generate the token

		break
	default:
		http.Error(w, "you cant do that 😡", 405)
		break

	}
}


func loginClientFromCamera(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "POST":
		w.Write([]byte("nice"))
		break
	case "GET":
		http.ServeFile(w,r,"./frontend/view/index.html")
	default:
		http.Error(w, "you cant do that 😡", 405)

	}

}


//========WEBSOCKETS===========\\
func receiveImages(w http.ResponseWriter, r *http.Request) {
	upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ := upgrade.Upgrade(w, r, nil)

	controlData(ws)
}

// this is the web camera is for receive the video and verify the token
func controlData(conn *websocket.Conn) {
	valid, name := false, ""
	var user streamCamera
	for {
		_, userJSON, err := conn.ReadMessage()
		if err != nil {
			delete(videoCamera, name) // if the client close the conn we delete the user from the map called videoCamera
			return
		}
		json.Unmarshal(userJSON, user) // this is for decode the formulary
		if valid {
			videoCamera[name] = user.Image // if all is good this add the video to the variable
			log.Println("we did it ")
		} else {
			verifyToken(user, &valid,&name) // if not we need to verify something for that

		}

	}
}


