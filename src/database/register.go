package database

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

// register func
func RegisterUserCameraDatabase(user stuff.RegisterCamera, okay chan bool) {

	// the query for insert the data
	q := `
		INSERT INTO 
			usercameras(ip,password,username,last_time_login)
			VALUES(?1,?2,?3,?4) `
	// we get the connection
	db, err := GetConnection()
	if err != nil {
		log.Println(err)
		okay <- false
		return // manage the errors
	} // manage the errors

	defer db.Close()
	//we use stm to avoid attacks

	stm, err := db.Prepare(q)
	if err != nil {
		okay <- false
		return
	}

	defer stm.Close()
	//then we run the query
	r, err := stm.Exec(
		stuff.EncryptData(user.IP),
		stuff.EncryptData(user.Password),
		user.Username,
		time.Now().UnixNano()/int64(time.Hour),
	)
	if err != nil {
		okay <- false
		log.Println(err)
		return // manage the errors
	}

	// if more than one file is affected we return an error
	i, err := r.RowsAffected()
	if err != nil {
		okay <- false
		log.Println(err)
		return // manage the errors
	}
	if i > 1 {
		log.Printf("idk why a row has been afected lol\n the query was %s \n the ip was %s \n the password was %s \n the username was %s", q, *stuff.EncryptData(user.IP), *stuff.EncryptData(user.Password), user.Username)
		if err != nil {
			okay <- false
			log.Println(err)
			return // manage the errors
		}
	}
	okay <- true

}