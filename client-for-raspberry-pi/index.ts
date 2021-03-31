import { StillCamera } from "pi-camera-connect";

// Take still image and save to disk
const runApp = async () => {
  const stillCamera = new StillCamera();

  const image: Buffer = await stillCamera.takeImage();

  console.log(image.toString("base64"));
};

runApp();
