function submitRegister(event, form){
    event.preventDefault();

    const data = new FormData(form)

    console.log(data)

    fetch("/api/register", {
        method: "POST",
        body: data
    }).then(response => {
        if (!response.ok){
            response.json().then(response => {
                console.log(response)
                form.querySelector(".error").textContent = response.message
            })
        }
        else{
            response.json().then(response => {
                console.log(response)
                setAuthToken(response.access_token, response.refresh_token)
                form.style.display = "none"
            })
        }
    })
}
