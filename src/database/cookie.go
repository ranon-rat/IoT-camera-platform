package database

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)






func AddTheCookieToTheDatabase(id string, cookieName string) {
	q := `INSERT INTO userclients(id_camera_client,cookie) VALUES(?1,?2);`
	/**

	| name             | type                 |
	| ---------------- | -------------------- |
	| id               | INTEGER PRIMARY KEY, |
	| id_camera_client | INTEGER NOT NULL,    |
	| cookie           | TEXT NOT NULL        |
	*/
	db, err := GetConnection()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()
	db.Exec(q, id, *stuff.EncryptData(cookieName))

}

