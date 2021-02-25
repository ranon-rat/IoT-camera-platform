import requests
from websocket import create_connection
import base64


# routes
url = "://localhost:8080"
url_register = f"http{url}/register"
url_login = f"http{url}/login"
# stream
token = ""
# open binary file in read mode
image = base64.encodestring(open('HelloImage.png', 'rb').read())
# body request
register_or_login = {
    "username": "pai",
    "password": "123"
}
# error managment
res = requests.post(url_register, register_or_login)

res = requests.post(url_login, register_or_login)
if res.content != "something is wrong, verify your password, or user":

    token = res.content
    stream_camera = {
        "token": token,
        "image": image
    }
    print(token)
    ws = create_connection(f"ws{url}/videoHandle")
    while True:
        ws.send(stream_camera)
