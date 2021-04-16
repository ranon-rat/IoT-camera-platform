package database

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

// we update the last time that he send somethings
func UpdateUsages(user stuff.RegisterCamera) {
	// the query
	q := `
		UPDATE usercameras 
			SET  last_time_login = ?1 
			WHERE username = ?2;
		`
	db, err := GetConnection() // get the connection
	if err != nil {
		log.Println(err,"its this")

		return // manage the errors
	} // manage the errors
	defer db.Close()
	db.Exec(q, time.Now().UnixNano()/int64(time.Hour), user.Username)
	// and exec the query

}