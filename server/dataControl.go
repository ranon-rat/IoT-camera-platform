package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func encryptData(data string) *string {

   	h := sha256.New()// we encript the data
	h.Write([]byte(data))
	v:=hex.EncodeToString(h.Sum(nil))
	return  &v

 
}

// get a simple connection
func getConnection() (*sql.DB,error){

	db, err := sql.Open("sqlite3", "./iotcameradata.db")

	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return db,nil
}
func exist(user string,ip string,sizeChan chan int) error{
	q:=`SELECT COUNT(*) 
		FROM usercameras 
		where username==? || ip ==? `
	db,err:=getConnection()
	if err!=nil{

		log.Println(err)
		close(sizeChan)
		return err
	}
	stm,err:=db.Prepare(q)
	if err!=nil{
		log.Println(err)
		close(sizeChan)	
		return err
	}
	HowMany,err:=stm.Query(user,ip)
	if err!=nil{
		log.Println(err)
		close(sizeChan)	
		return err 
	}
	var size int
	for HowMany.Next(){
		err = HowMany.Scan(&size)
		if err != nil {
			close(sizeChan)
			return err

		}

	}
	sizeChan<-size

	return nil

}

func uploadUserCameraDatabase(user register,errChan chan error) {
	sizeChan:=make(chan int)
	// we check if the username of the camera already exist
	go exist(user.Username,user.IP,sizeChan)
	if <-sizeChan>0{
		errChan<- errors.New("sorry but that user has already registered")
		return 
	}
	// the query for insert the data
	q:=`INSERT INTO 
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
	db,err:=getConnection()
	if err!=nil{
		log.Println(err)
		errChan<-err
		return
	}

	defer db.Close()
	//we use stm to avoid attacks
	
	stm,err:=db.Prepare(q)
	if err!=nil{
		log.Println(err)
		errChan<-err
		return  
	}
	defer stm.Close()
	//then we run the query
	r, err := stm.Exec(
		encryptData(user.IP), 
		encryptData(user.Password), 
		encryptData(user.Username),
	)
	if err != nil {
		log.Println(err)
		errChan<-err
		return
	}
	// if more than one file is affected we return an error
	i, _ := r.RowsAffected()
	if i != 1 {
		errChan<- fmt.Errorf("idk why a row has been afected lol\n the query was %s \n the ip was %s \n the password was %s \n the username was %s",q,
			user.IP,  
			user.Password, 
			user.Username,
		)
		return
	}
	
	close(errChan)
}