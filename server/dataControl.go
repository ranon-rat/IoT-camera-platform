package main

import (
	"database/sql"
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
func exist(user string, ip string, sizeChan chan int) {

	q := `SELECT COUNT(*) 
		FROM usercameras 
		WHERE username=?1 OR ip =?2 ;` // igual aqui
	var size int
	// GET A CONNECTION
	db, _ := getConnection()

	HowMany, err := db.Query(q, user, ip)
	if err != nil {
		log.Println(err)

		if err != nil {
			log.Println(err)
			close(sizeChan)

			return // manage the errors
		}

	}

	defer HowMany.Close()

	for HowMany.Next() {
		err = HowMany.Scan(&size)
		if err != nil {
			log.Println(err)
			if err != nil {
				log.Println(err)
				close(sizeChan)

				return // manage the errors
			}

		}

	}

	sizeChan <- size

}

// register func
func registerUserCameraDatabase(user registerCamera, okay chan bool) {

	// the query for insert the data
	q := `
	BEGIN TRANSACTION;
		INSERT INTO 
			usercameras(ip,password,username,last_time_login)
			VALUES(?1,?2,?3,?4) 
	END TRANSACTION;`
	// we get the connection
	db, err := getConnection()
	if err != nil {
		log.Println(err)
		okay <- false
		return // manage the errors
	} // manage the errors

	defer db.Close()
	//we use stm to avoid attacks

	stm, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
		okay <- false
		return // manage the errors
	} // manage the errors

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
		log.Printf("idk why a row has been afected lol\n the query was %s \n the ip was %s \n the password was %s \n the username was %s", q, *encryptData(user.IP), *encryptData(user.Password), user.Username)
		if err != nil {
			okay <- false
			log.Println(err)
			return // manage the errors
		}
	}
	okay <- true

}

// login func
func loginUserCameraDatabase(user registerCamera, validChan chan bool) {

	q := `SELECT COUNT(*) FROM usercameras  
	WHERE username = ?1 AND password= ?2;` // aqui no accedemos a la informacion , accedemos a la cantidad de usuarios que coinciden
	// get the connection
	db, _ := getConnection()

	defer db.Close()
	// make the consult and encrypt the data
	valid, err := db.Query(q, user.Username,
		encryptData(user.Password))
	if err != nil {
		log.Println(err)
		validChan <- false
		return // manage the errors
	} // manage the errors

	// review the results
	var i int
	for valid.Next() {
		// change the value of i
		err = valid.Scan(&i)

		if err != nil {
			log.Println(err)
			validChan <- false
			return // manage the errors
		} // manage the errors
	}
	validChan <- i > 0

}

// we generate the token
func generateToken(user registerCamera, tokenChan chan string, okay chan bool) {

	q := `
	BEGIN TRANSACTION;
		UPDATE usercameras 
			SET token = ?1 
			WHERE username = ?2 AND password=?3;
	END TRANSACTION;`
	// we get a connection
	db, err := getConnection()
	if err != nil {
		log.Println(err)
		close(tokenChan)
		okay <- false
		return // manage the errors
	}
	// generate the token
	token := *encryptData(fmt.Sprintf("%f%s%f", rand.Float64()*1000, (*encryptData(user.Password) + *encryptData(user.Username)), rand.Float64()*1000))
	defer db.Close()
	// prepare the sentence
	stm, _ := db.Prepare(q)
	stm.Exec(encryptData(token), user.Username, user.Password)
	tokenChan <- (token) // and send the token
	okay <- true

}

// we update the last time that he send somethings
func updateUsages(user registerCamera, okay chan bool) {
	// the query
	q := `BEGIN TRANSACTION;
			UPDATE usercameras 
				SET  last_time_login = ?1 
				WHERE username = ?2;
		END TRANSACTION;`
	db, err := getConnection() // get the connection
	if err != nil {
		log.Println(err)
		okay <- false
		return // manage the errors
	} // manage the errors
	defer db.Close()
	db.Exec(q, time.Now().UnixNano()/int64(time.Hour), user.Username)
	// and exec the query
	okay <- true

}
func verifyToken(camera streamCamera, valid chan bool, nameChan chan string) {
	q := `SELECT name FROM usercameras 
		WHERE token=?1;
		BEGIN TRANSACTION;
			UPDATE usercameras 
				SET  last_time_login = ?1
				WHERE token = ?1;
		END TRANSACTION;`
	// uso esto para cambiar la ultima ves que se conecto
	db, err := getConnection()
	if err != nil {
		valid <- false
		log.Println(err)
		close(nameChan)
		return
	}
	defer db.Close()
	info, err := db.Query(q, *encryptData(camera.Token), time.Now().UnixNano()/int64(time.Hour))
	if err != nil {
		valid <- false
		log.Println(err)
		close(nameChan)
		return

	}
	names, name := []string{}, ""

	for info.Next() {

		err = info.Scan(&name)
		if err != nil {
			valid <- false
			close(nameChan)
			log.Println(err)
			return
		}
	}
	names = append(names, name)
	valid <- len(names) > 0
	nameChan <- name
}
