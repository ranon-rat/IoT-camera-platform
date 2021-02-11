package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
)
func setupRoutes() error{
	log.Println("setup router")

	r := mux.NewRouter()
	
	r.Handle("/public",http.FileServer(http.Dir("./public")))
	r.HandleFunc("/register",registerUser)
	r.HandleFunc("/login",loginUserCamera)
	r.HandleFunc("/start/{user}",func(w http.ResponseWriter, r *http.Request){
		routesvars:=mux.Vars(r)
		user,err:= routesvars["user"]
		if !err{
			http.Error(w,"user not find",401)
		}
		w.Write([]byte(videoCamera[user]))
	})
	err:=http.ListenAndServe(":8080",r)
	if err!=nil{
		return err
	}
	return nil
}