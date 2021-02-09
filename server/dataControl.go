package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// get a simple connection
func getConnection() (*sql.DB,error){

	db, err := sql.Open("sqlite3", "./iotcameradata.db")

	if err != nil {
		fmt.Println(err)
		return nil,err
	}
	return db,nil
}

func uploadUserCameraDatabase(user register) error{
	q:=`INSERT INTO 
	usercameras(ip,password,username,last_time_login)
	VALUES(?1,?2,?3,?4) `
	/*
	usercameras:
	id INTEGER PRIMARY KEY,
    ip VARCHAR(50),
    password TEXT,
    username TEXT,
    last_time_login INTEGER
	*/
	db,err:=getConnection()
	if err!=nil{
		log.Println(err)
		return err
	}
	defer db.Close()
	stm,err:=db.Prepare(q)
	if err!=nil{
		log.Println(err)
		return  err
	}
	r, err := stm.Exec(&user.IP, &user.Password, &user.User,)
	if err != nil {
		log.Println(err)
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("se esperaba una sola fila omg")
	}

	fmt.Println(q)
	return nil
}