# [routes.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/routes.go)
## para que es?

es para poner las rutas que hara el ser servidor, solo tiene una funcion `setupRoutes() ` la cual regresa un error si algo sucede 
no es muy dificil y bueno eso solo hay que ponerlo en el main

# [types.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/types.go)

## para que es?
es solo para poner ciertas cosas y tener un mayor orden en la estructura del proyecto , ahi ponemos los tipos y otras cosas como variables y constantes y el unico tipo que por ahora hay es el siguiente 
## register
este se hizo para la base de datos
```go
type register struct{
	Password 		string `json:"password"`
	Username		string `json:"username"`
	IP 				  string 

}
```
# [controllers.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/controllers.go)

## para que es ?

es un archivo que se hizo especificamente para hacer los handler request asi se pueden evitar ciertos problemas con algunas cosas

## registerUser()

esta funcion esta hecha para manejar los request de la ruta `/register`
si haces un post request ejecutara la funcion `registerUserCameraDatabase()` 
sino pues te regresa otra respuesta
no es muy dificil de entender.

## loginUser()

esta funcion tambien sirve para manejar los request de la ruta `login`
si haces un post request deberia de ejecutar la funcion `loginUserCameraDatabase()` aun que por ahora esa funcion no esta disponible y de eso deberias de encargarte tu 

<!--------------------->

# [dataControl.go](https://github.com/ranon-rat/IoT-camera-/blob/master/server/dataControl.go)

## para que es ?

En este archivo se realizan ciertas operaciones para funcionar con la base de datos
<!--------------------->
## getConnection()
vale , la funcion con la cual se conecta la base de datos es esta
```go
func getConnection() (*sql.DB,error){}
```
la base de datos con la que se conecta es `iotcameradata.db` ahi en la funcion `getConnection()` se conecta de manera directa asi que no hace falta hacer mucho 
<!--------------------->
## exist()
la funcion `exist()` es para poder  saber si el usuario que se ha tratado de registrar ya existe asi se puede evitar algunos problemas cuando registramos el usuario,
el query que ejecuta esta funcion es
```sql
SELECT COUNT(*) 
		FROM usercameras 
		where username==? || ip ==? ;`
```
y la funcion es 
```go
func exist(user string,ip string,sizeChan chan int) error
```
asi que te puede regresar un error y se debe de usar de manera concurrente `go exist("Fgh","127.0.0.1",sizeChan)`  
<!--------------------->
##  registerUserCameraDatabase()

Esta funcion sirve para registrar el usuario siempre y cuando haga un post request

El query que ejecuta esta funcion es el siguiente 
```go
INSERT INTO 
	usercameras(
		ip,
		password,
		username,
		last_time_login
	)
	VALUES(?1,?2,?3,?4) 
```
ahi agrega informacion importante con la cual se puede trabajar a futuro 

la estructura de la funcion es la siguiente 
```go
func registerUserCameraDatabase(user register,errChan chan error)
```
esta hecha para las goroutines asi que se debe de poner lo tipico `go registerUserCameraDatabase(user,errChan)`
y  con eso bastaria

<!--------------------->
## la estructura de la base de datos es la siguiente

###            usercameras              

|		name            |        type        |	
|-------------------|--------------------|
|id                 |INTEGER PRIMARY KEY |
|ip                 |VARCHAR(64)         |
|password           |TEXT                |
|username           |TEXT                |
|last_time_login    |INTEGER             |

###  userclients

|name | type|
|---|---|
|   cookie |TEXT |

por ahora no hemos implementado esta base de datos pero asi funciona


