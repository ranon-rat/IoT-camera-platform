package database

import (
	"fmt"
	"log"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

// we generate the token
func GenerateToken(user stuff.RegisterCamera, tokenChan chan string) {

	q := `UPDATE usercameras 
			SET token = ?1 
			WHERE username = ?2 ;
	`
	// we get a connection
	db, err := GetConnection()
	if err != nil {
		log.Println(err)
		close(tokenChan)
		return // manage the errors
	}
	// generate the token
	token := *stuff.EncryptData(fmt.Sprintf("%f%s%f", rand.Float64()*1000, (*stuff.EncryptData(user.Password) + *stuff.EncryptData(user.Username)), rand.Float64()*1000))
	defer db.Close()
	// prepare the sentence
	stm, _ := db.Prepare(q)

	stm.Exec(stuff.EncryptData(token), user.Username)

	// and send the token

	tokenChan <- (token)

}
