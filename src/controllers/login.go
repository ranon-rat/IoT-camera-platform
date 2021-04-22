package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ranon-rat/IoT-camera-platform/server/src/database"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

// this is for login the user and send you that
func LoginUserCamera(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// we setup the values from the struct registerCamera
		var oldUser stuff.RegisterCamera
		json.NewDecoder(r.Body).Decode(&oldUser)
		oldUser.IP = r.Header.Get("x-forwarded-for")
		//  we use this for asynchronous communication
		valid, token := make(chan bool), make(chan string)
		// check if all is okay

		go database.LoginUserCameraDatabase(oldUser, valid)
		// no es el login

		if <-valid {

			go database.UpdateUsages(oldUser)
			// we update the last time that he send something

			go database.GenerateToken(oldUser, token)

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
func LoginClientFromCamera(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		valid, newUserReal := make(chan bool), stuff.RegisterCamera{}
		json.NewDecoder(r.Body).Decode(&newUserReal)
		go database.LoginUserCameraDatabase(newUserReal, valid)
		if <-valid {

			valueCookie := *stuff.EncryptData(fmt.Sprintf("%s%f%s",
				newUserReal.Password, rand.Float64()*1000,
				newUserReal.Username))

			go database.AddTheCookieToTheDatabase(newUserReal.Username, valueCookie)

			http.SetCookie(w, &http.Cookie{
				Name:    newUserReal.Username,
				Value:   valueCookie,
				Expires: time.Now().AddDate(0, 0, 1),
			})
			return
		}
		break
	case "GET":
		http.ServeFile(w, r, "./frontend/view/index.html")
	default:
		http.Error(w, "you cant do that 😡", 405)

	}

}
