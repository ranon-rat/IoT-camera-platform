package database

import (
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ranon-rat/IoT-camera-platform/server/src/stuff"
)

func VerifyToken(camera stuff.StreamCamera, valid *bool, nameChan *string) {
	q := `SELECT name FROM usercameras 
		WHERE token=?1;
			UPDATE usercameras 
				SET  last_time_login = ?1
				WHERE token = ?1;`
	// uso esto para cambiar la ultima ves que se conecto
	db, err := GetConnection()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	info, err := db.Query(q, *stuff.EncryptData(camera.Token), time.Now().UnixNano()/int64(time.Hour))
	if err != nil {
		log.Println(err)
		return

	}
	names, name := []string{}, ""

	for info.Next() {
		err = info.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}
	}
	names = append(names, name)
	*valid = len(names) > 0
	*nameChan = name
}


// verify the cookies 
func VerifyTheCookie(cookieName string, validCookie chan bool) {
	q := `SELECT COUNT(*) FROM userclients 
			WHERE cookie =?1;
	`
	db, err := GetConnection()
	if err != nil {
		log.Println(err.Error())
		validCookie <- false
		return
	}
	defer db.Close()
	c, _ := db.Query(q, cookieName)
	cSize := 0
	for c.Next() {
		c.Scan(&cSize)
	}
	validCookie <- cSize == 1

}