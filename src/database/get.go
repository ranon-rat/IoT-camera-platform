package database

import (
	"database/sql"
	"log"
	"path"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

// i need more comments for do something because i cant die aaa
// get a simple connection
func GetConnection() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./sql/iotcameradata.db")
	if err != nil {
		if _, file, _, ok := runtime.Caller(0); ok {
			__dirname := path.Dir(file)
			log.Println("__dirname:", __dirname,ok)
		}
		log.Println(err, "get connection")
		return nil, err
	}
	return db, nil
}
func GetID(user stuff.RegisterCamera, idChan chan string) {
	q := `SELECT username FROM usercameras
		WHERE username=?1 `
	db, err := GetConnection()
	if err != nil {
		log.Println(err.Error())
		idChan <- ""
		return
	}
	defer db.Close()
	idRow, _ := db.Query(q, user.Username)
	id := 0
	for idRow.Next() {
		idRow.Scan(&id)
	}
	idChan <- ""

}
