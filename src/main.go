package main

import (
	"log"

	"github.com/ranon-rat/IoT-camera-platform/server/src/router"
)
func main(){
	log.Println(router.SetupRoutes())

}