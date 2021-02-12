package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
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
func exist(user string, ip string, w chan codeHTTP, sizeChan chan int) {
	q := `SELECT COUNT(*) 
		FROM usercameras 
		WHERE username=?1 OR ip =?2 ;` // igual aqui
	var size int
	// GET A CONNECTION
	db, _ := getConnection()

	HowMany, err := db.Query(q, user, ip)
	if err != nil {
		log.Println(err)
		close(sizeChan)
		errorControl(err, w, "internal server error", 500)

	}
	defer HowMany.Close()

	for HowMany.Next() {
		err = HowMany.Scan(&size)
		if err != nil {
			log.Println(err)
			errorControl(err, w, "internal server error", 500)
			close(sizeChan)

		}

	}

	sizeChan <- size

}

// register func
func registerUserCameraDatabase(user registerCamera, code chan codeHTTP) {
	sizeChan := make(chan int)
	if len(user.Username) == 0 || len(user.Password) == 0 {
		errorControl(errors.New("some value is empty"), code, "some value is empty", 500)
		return
	}
	// we check if the username of the camera already exist
	go exist(user.Username, *encryptData(user.IP), code, sizeChan)
	if <-sizeChan > 0 {
		errorControl(errors.New("that user has been already registered"), code, "that user has been registered", 400)

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
	errorControl(err, code, "internal server error", 500) // manage the errors

	defer db.Close()
	//we use stm to avoid attacks

	stm, err := db.Prepare(q)
	errorControl(err, code, "internal server error", 500) // manage the errors

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
		errorControl(errors.New("internal server error"), code, "intelan server error", 500)

	}
	errorControl(nil, code, "", 0)

}

// login func
func loginUserCameraDatabase(user registerCamera, code chan codeHTTP, validChan chan bool) {
	q := `SELECT COUNT(*) FROM usercameras  
	WHERE username = ?1 AND password= ?2;` // aqui no accedemos a la informacion , accedemos a la cantidad de usuarios que coinciden
	// get the connection
	db, _ := getConnection()

	defer db.Close()
	// make the consult and encrypt the data
	valid, err := db.Query(q, user.Username,
		encryptData(user.Password))
	errorControl(err, code, "internal server error", 500) // manage the errors

	// review the results
	var i int
	for valid.Next() {
		// change the value of i
		err = valid.Scan(&i)
		errorControl(err, code, "internal server error", 500) // manage the errors

	}

	validChan <- i > 0

}

// we generate the token
func generateToken(user registerCamera, code chan codeHTTP, tokenChan chan string) {
	q := `UPDATE usercameras SET token = ?1 WHERE username = ?2;`
	// we get a connection
	db, err := getConnection()
	errorControl(err, code, "internal server error", 500) // manage the errors

	// generate the token
	token := *encryptData(fmt.Sprintf("%s%d", (*encryptData(user.Password) + *encryptData(user.Username)), rand.Int()))
	defer db.Close()
	// prepare the sentence
	stm, _ := db.Prepare(q)
	stm.Exec(token, user.Username)
	tokenChan <- (token) // and send the token

}

// we update the last time that he send somethings
func updateUsages(user registerCamera, code chan codeHTTP) {
	// the query
	q := `UPDATE usercameras SET  last_time_login = ?1 WHERE username = ?2;`
	db, err := getConnection()                            // get the connection
	errorControl(err, code, "internal server error", 500) // manage the errors
	defer db.Close()
	db.Exec(q, time.Now().UnixNano()/int64(time.Hour), user.Username) // and exec the query
}
