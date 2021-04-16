package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ranon-rat/IoT-camera-platform/server/src/database"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		var newUser stuff.RegisterCamera
		json.NewDecoder(r.Body).Decode(&newUser)
		newUser.IP = r.Header.Get("x-forwarded-for")
		// if the password is empty send you this

		if len(newUser.Username) < 3 || len(newUser.Password) < 3 {
			http.Error(w, "your password or your username is empty or is less than 4 characters", 406)
			return
		}

		errChan := make(chan bool)
		// creo que esto deberia de marcarnos si sizeChan esta cerrado o no
		// asi creo que PODRIAMOS manejar  los errores

		// we check if the username of the camera already exist

		//this register the user for the database
		go database.RegisterUserCameraDatabase(newUser, errChan)

		//if everything is fine send you this
		w.Write([]byte("now you are registered "))

		break

	default:
		http.Error(w, "you cant do that 😡", 405)
		break
	}
}
