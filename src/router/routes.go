package router

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/IoT-camera-platform/server/src/controllers"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
	//"github.com/ranon-rat/IoT-camera-platform/src/controllers"
)

func SetupRoutes() error {
	log.Println("setup router")
	r := mux.NewRouter()
	
	r.HandleFunc("/", controllers.LoginClientFromCamera)
	r.HandleFunc(`/frontend/{file:[\/\w\d\W]+?}`, func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	r.HandleFunc("/register", controllers.RegisterUser)
	
	r.HandleFunc("/login", controllers.LoginUserCamera)
	r.HandleFunc("/videoHandle", controllers.ReceiveImages)
	r.HandleFunc("/start/{user}", func(w http.ResponseWriter, r *http.Request) {
		routesvars := mux.Vars(r)
		user, err := routesvars["user"]
		if !err {
			http.Error(w, "user not find", 401)
		}
		w.Write([]byte(stuff.VideoCamera[user]))
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}
	return nil
}
