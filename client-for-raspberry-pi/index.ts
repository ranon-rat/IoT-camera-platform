import { StillCamera } from "pi-camera-connect";
import axios, * from "axios"
import * as fs from "fs";
const host = "192.168.100.21:8080"
const config:from.AxiosRequestConfig = {
    method: "POST",
    data: {
        username: "penisCum",
        password: "1234",
    }
};

let token: string = "";// Take still image and save to disk

axios("http://" + host + "/register", config).then(r => r.data).catch();
axios("http://" + host + "/login", config).then(r => token = r.data).catch();
const runApp = async () =>
{
    const stillCamera = new StillCamera();

    const image = await stillCamera.takeImage();

    console.log(image.toString("base64"));
};

runApp();
