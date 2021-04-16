package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/ranon-rat/IoT-camera-platform/server/src/database"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

//========WEBSOCKETS===========\\
func ReceiveImages(w http.ResponseWriter, r *http.Request) {
	stuff.Upgrade.CheckOrigin = func(r *http.Request) bool { return true }
	ws, _ := stuff.Upgrade.Upgrade(w, r, nil)

	ControlData(ws)
}

// this is the web camera is for receive the video and verify the token
func ControlData(conn *websocket.Conn) {
	valid, name := false, ""
	var user stuff.StreamCamera
	for {
		_, userJSON, err := conn.ReadMessage()
		if err != nil {
			delete(stuff.VideoCamera, name) // if the client close the conn we delete the user from the map called videoCamera
			return
		}
		json.Unmarshal(userJSON, user) // this is for decode the formulary
		if valid {
			stuff.VideoCamera[name] = user.Image // if all is good this add the video to the variable
			log.Println("we did it ")
		} else {
			database.VerifyToken(user, &valid, &name) // if not we need to verify something for that

		}

	}
}