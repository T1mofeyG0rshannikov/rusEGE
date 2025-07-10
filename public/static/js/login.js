function submitLogin(event){
    event.preventDefault();

    const modal = document.getElementById("login")
    const form = modal.querySelector("form")
    const data = new FormData(form)

    console.log(data)

    fetch("/api/login", {
        method: "POST",
        headers: {
            'Content-type': 'application/json'
        },
        body: data
    }).then(response => {
        console.log(response)
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
                modal.style.display = "none"
            })
        }
    })
}
