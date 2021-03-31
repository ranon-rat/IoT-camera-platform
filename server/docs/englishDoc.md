# [routes.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/routes.go)
## what is it for?

is to set the routes that the server will do, it only has a `setupRoutes ()` function which returns an error if something happens
It is not very difficult and well, you just have to put it in the main

# [types.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/types.go)

## what is it for?
It is only to put certain things and have a greater order in the structure of the project, there we put the types and other things as variables and constants and the only type that for now is the following
## register
this was made for the database
```go
type  register struct {
Password    string `json:"password"`
Username    string `json:"username"`
Image       string `json:"image"`
IP          string

}
```
# [controllers.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/controllers.go)

## what is it for?

It is a file that was made specifically to make the handler requests so you can avoid certain problems with some things

## registerUser ()

This function is made to handle requests from the `/ register` route
if you make a post request it will execute the function `registerUserCameraDatabase ()`
but then you get another answer
it is not very difficult to understand.

## loginUser ()

This function also serves to handle requests from the `login` route
If you make a post request, you should execute the `loginUserCameraDatabase ()` function, even though for now that function is not available and you should take care of that.

<!--------------------->

# [dataControl.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/dataControl.go)

## what is it for?

In this file certain operations are performed to work with the database
<!--------------------->
## getConnection ()
ok, the function with which the database connects is this
```go
func getConnection () (* sql.DB, error) {}
```
the database with which it connects is `iotcameradata.db` there in the` getConnection () `function it connects directly so you don't need to do much
<!--------------------->

## registerUserCameraDatabase ()

This function serves to register the user as long as it makes a post request

The query that executes this function is the following
```go
INSERT INTO
usercameras
ip,
password,
username,
last_time_login
)
VALUES (? 1,? 2,? 3,? 4)
```
There you add important information with which you can work in the future

the structure of the function is as follows
```go
func registerUserCameraDatabase (user register, errChan chan error)
```
It is made for goroutines so you should put the typical `go registerUserCameraDatabase (user, errChan)`
and with that it would be enough

## loginUserCameraDatabase ()

### what is it for?
It is a function that is made to make a login, it already works and it is quite good

### query
the query that executes this function is
```sql
SELECT COUNT (*) FROM usercameras
WHERE username =? 1 AND password =? 2;
```
### Body

the body of the function is this
```go
func loginUserCameraDatabase (user register, validChan chan bool)
```
It is recommended that you use co `go loginUserCameraDatabase` to control it since it is a function that was made specifically to log in concurrently, it returns a` true` if the query value is greater than 1
## updateUsages ()

### what is it for?
We have done this function for a simple reason ... it is to be able to update and know when was the last time you sent an image, so we can avoid some problems that usually happen

### Query
the query that executes this function is
```sql
UPDATE usercameras
SET last_time_login =? 1
WHERE username =? 2; `
```

## Body
This function is made to work in `goroutines` so it is recommended that you use it that way
```go
func updateUsages (user register)
```


<! --------------------->
## the database structure is as follows

### usercameras

| name | type |
| ------------------- | -------------------- |
| id | INTEGER PRIMARY KEY |
| ip | TEXT |
| password | TEXT |
| username | TEXT |
| last_time_login | INTEGER |

### userclients

| name | type |
| --- | --- |
| cookie | TEXT |

for now we have not implemented this database but it works like this
