package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func encryptData(data string) *string {
	sum := sha256.Sum256([]byte(data)) // we encript the data
	v := hex.EncodeToString(sum[:])
	return &v
}

// i need more comments for do something because i cant die aaa
// get a simple connection
func getConnection() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./iotcameradata.db")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
func exist(user string, ip string, sizeChan chan int) error {
	q := `SELECT COUNT(*) 
		FROM usercameras 
		WHERE username=?1 OR ip =?2 ;`
	var size int
	// GET A CONNECTION
	db, err := getConnection()
	if err != nil {
		log.Println(err)
		close(sizeChan)
		return err
	}
	HowMany, err := db.Query(q, user, ip)
	if err != nil {
		log.Println(err)
		close(sizeChan)
		return err
	}
	defer HowMany.Close()

	for HowMany.Next() {
		err = HowMany.Scan(&size)
		if err != nil {
			log.Println(err)
			close(sizeChan)
			return err

		}

	}
	log.Println(size)
	sizeChan <- size

	return nil

}

// register func
func registerUserCameraDatabase(user register, errChan chan error) {
	sizeChan := make(chan int)
	// we check if the username of the camera already exist
	go exist(user.Username, *encryptData(user.IP), sizeChan)
	if <-sizeChan > 0 {
		errChan <- errors.New("sorry but that user has already registered")
		return
	}
	if len(user.Username) == 0 || len(user.Password) == 0 {
		errChan <- fmt.Errorf("some value is empty\nusername:%s\npassword%s",
			user.Username,
			user.Password,
		)
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
	/*
		this is how the table is
		__________________________________________
		|               usercameras              |
		|----------------------------------------|
		|		name        |        type        |
		|-------------------|--------------------|
		|id                 |INTEGER PRIMARY KEY |
		|ip                 |VARCHAR(64)         |
		|password           |TEXT                |
		|username           |TEXT                |
		|last_time_login    |INTEGER             |
		|----------------------------------------|
	*/
	// we get the connection
	db, err := getConnection()
	if err != nil {
		log.Println(err)
		errChan <- err
		return
	}

	defer db.Close()
	//we use stm to avoid attacks

	stm, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
		errChan <- err
		return
	}
	defer stm.Close()
	//then we run the query
	r, err := stm.Exec(
		encryptData(user.IP),
		encryptData(user.Password),
		user.Username,
		time.Now().UnixNano()/int64(time.Hour),
	)
	if err != nil {
		log.Println(err)
		errChan <- err
		return
	}
	// if more than one file is affected we return an error
	i, _ := r.RowsAffected()
	if i != 1 {
		errChan <- fmt.Errorf("idk why a row has been afected lol\n the query was %s \n the ip was %s \n the password was %s \n the username was %s", q,
			user.IP,
			user.Password,
			user.Username,
		)
		return
	}

	close(errChan)
}

// login func
func loginUserCameraDatabase(user register, validChan chan bool) {
	q := `SELECT COUNT(*) FROM usercameras  
	WHERE username = ?1 AND password= ?2;`
	// get the connection
	db, err := getConnection()
	if err != nil {
		log.Println(err)
		validChan <- false

		return
	}
	defer db.Close()
	// make the consult
	// encript the data
	valid, err := db.Query(q, user.Username,
		encryptData(user.Password))
	// review the results
	var i int
	for valid.Next() {
		// change the value of i
		err = valid.Scan(&i)
		if err != nil {
			log.Println(err)
			validChan <- false
			return
		}
	}
	println(i)
	validChan <- i > 0

}
func updateUsages(user register) {
	q := `UPDATE usercameras
		SET last_time_login = ?1
		WHERE username =?2;`
	db, err := getConnection()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	db.Exec(q, time.Now().UnixNano()/int64(time.Hour), user.Username)

}
