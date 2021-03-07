const sendPassword = () => {
  let username: HTMLElement = document.getElementById("name");
  let password: HTMLElement = document.getElementById("password");

  console.log(username.value, password.value);
  username.value = "";
  password.value = "";
};
