import requests
#routes
url_register = "http://localhost:8080/register"
url_login = "http://localhost:8080/login"
#body request
register_or_login = {
    "username": "pai",
    "password": "123"
}
#error managment
res = requests.post(url_register, register_or_login)
if res.content == "sorry but that user has already registered":
    res = requests.post(url_login, register_or_login)
