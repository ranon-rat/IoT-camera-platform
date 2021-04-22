package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

// login func
func LoginUserCameraDatabase(user stuff.RegisterCamera, validChan chan bool) {

	q := `SELECT COUNT(*) FROM usercameras  
	WHERE username = ?1 AND password= ?2;` // aqui no accedemos a la informacion , accedemos a la cantidad de usuarios que coinciden
	// get the connection
	db, _ := GetConnection()

	defer db.Close()
	// make the consult and encrypt the data
	valid, err := db.Query(q, user.Username,
		stuff.EncryptData(user.Password))
	if err != nil {
		log.Println(err)
		validChan <- false
		return // manage the errors
	} // manage the errors

	// review the results
	var i int

	// change the value of i
	for valid.Next() {
		valid.Scan(&i)
	}

	validChan <- i > 0

}
