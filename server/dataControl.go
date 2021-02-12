package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// i need more comments for do something because i cant die aaa
// get a simple connection
func getConnection() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./database/iotcameradata.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
func exist(user string, ip string, w http.ResponseWriter, sizeChan chan int) {
	q := `SELECT COUNT(*) 
		FROM usercameras 
		WHERE username=?1 OR ip =?2 ;`
	var size int
	// GET A CONNECTION
	db, _ := getConnection()

	HowMany, err := db.Query(q, user, ip)
	if err != nil {
		log.Println(err)
		close(sizeChan)
		http.Error(w, "internal server error", 500)
	}
	defer HowMany.Close()

	for HowMany.Next() {
		err = HowMany.Scan(&size)
		if err != nil {
			log.Println(err)
			close(sizeChan)
			http.Error(w, "internal server error", 500)

		}

	}

	sizeChan <- size

}

// register func
func registerUserCameraDatabase(user registerCamera, w http.ResponseWriter) {
	sizeChan := make(chan int)
	if len(user.Username) == 0 || len(user.Password) == 0 {
		http.Error(w, "some value is empty", 406)
		return
	}
	// we check if the username of the camera already exist
	go exist(user.Username, *encryptData(user.IP), w, sizeChan)
	if <-sizeChan > 0 {
		w.Write([]byte("sorry but that user has already registered"))
		return
	}

	// the query for insert the data
	q := `INSERT INTO 
	usercameras(
		ip,
		password,
		username,
		last_time_login
	)
	VALUES(?1,?2,?3,?4) `
	// we get the connection
	db, err := getConnection()
	errorControl(err, w, "internal server error", 500) // manage the errors

	defer db.Close()
	//we use stm to avoid attacks

	stm, err := db.Prepare(q)
	errorControl(err, w, "internal server error", 500) // manage the errors

	defer stm.Close()
	//then we run the query
	r, _ := stm.Exec(
		encryptData(user.IP),
		encryptData(user.Password),
		user.Username,
		time.Now().UnixNano()/int64(time.Hour),
	)
	// if more than one file is affected we return an error
	i, _ := r.RowsAffected()
	if i > 1 {

		log.Printf("idk why a row has been afected lol\n the query was %s \n the ip was %s \n the password was %s \n the username was %s", q,
			*encryptData(user.IP),
			*encryptData(user.Password),
			user.Username,
		)
		http.Error(w, "internal server error", 500)
	}
	w.Write([]byte("all its okay"))

}

// login func
func loginUserCameraDatabase(user registerCamera, w http.ResponseWriter, validChan chan bool) {
	q := `SELECT COUNT(*) FROM usercameras  
	WHERE username = ?1 AND password= ?2;`
	// get the connection
	db, _ := getConnection()

	defer db.Close()
	// make the consult and encrypt the data
	valid, err := db.Query(q, user.Username,
		encryptData(user.Password))
	errorControl(err, w, "internal server error", 500) // manage the errors

	// review the results
	var i int
	for valid.Next() {
		// change the value of i
		err = valid.Scan(&i)
		errorControl(err, w, "internal server error", 500) // manage the errors

	}

	validChan <- i > 0

}

// we generate the token
func generateToken(user registerCamera, w http.ResponseWriter, tokenChan chan string) {
	q := `UPDATE usercameras SET token = ?1 WHERE username = ?2;`
	// we get a connection
	db, err := getConnection()
	errorControl(err, w, "internal server error", 500) // manage the errors

	// generate the token
	token := *encryptData(fmt.Sprintf("%s%d", (*encryptData(user.Password) + *encryptData(user.Username)), rand.Int()))
	defer db.Close()
	// prepare the sentence
	stm, _ := db.Prepare(q)
	stm.Exec(token, user.Username)
	tokenChan <- (token) // and send the token

}

// we update the last time that he send somethings
func updateUsages(user registerCamera, w http.ResponseWriter) {
	// the query
	q := `UPDATE usercameras SET  last_time_login = ?1 WHERE username = ?2;`
	db, err := getConnection()                         // get the connection
	errorControl(err, w, "internal server error", 500) // manage the errors

	defer db.Close()
	db.Exec(q, time.Now().UnixNano()/int64(time.Hour), user.Username) // and exec the query

}
