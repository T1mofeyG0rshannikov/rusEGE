async function submitLogin(event){
    event.preventDefault();

    console.log("fwfww")
    const modal = document.getElementById("login")
    const form = modal.querySelector("form")
    const data = new FormData(form)

    const object = {};
    data.forEach((value, key) => object[key] = value);
    const json = JSON.stringify(object);

    response = await loginAPI(json)

    console.log(response)

    if (response.status === 200){
        setAuthToken(response.data.access_token, response.data.refresh_token)
        modal.style.display = "none"
        checkIsAuth()
    }
    else{
        console.log(response)
        form.querySelector(".error").textContent = response.data.message
    }
}
