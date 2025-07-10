function submitRegister(event){
    event.preventDefault();

    const modal = document.getElementById("register")
    const form = modal.querySelector("form")
    const data = new FormData(form)
    data.delete('repeat-password')

    const object = {};
    data.forEach((value, key) => object[key] = value);
    const json = JSON.stringify(object);

    fetch("/api/register", {
        method: "POST",
        headers: {
            'Content-type': 'application/json'
        },
        body: json
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
                modal.style.display = "none"
            })
        }
    })
}
