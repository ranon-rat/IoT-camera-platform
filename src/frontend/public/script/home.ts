// @ts-nocheck
const sendPassword = () => {
  let username: HTMLElement = document.getElementById("name");
  let password: HTMLElement = document.getElementById("password");

  fetch("/", {
    method: "POST",
    body: JSON.stringify({
      username: username.value,
      password: password.value,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });
  console.log(username.value, password.value);
  username.value = "";
  password.value = "";
};
