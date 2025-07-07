function openRegisterForm(event){
    event.preventDefault()
    document.querySelectorAll('.modal').forEach(m => {
        m.style.display = "none"
    })

    document.getElementById("register").style.display = "block"
}


function openLoginForm(event){
    event.preventDefault()
    document.querySelectorAll('.modal').forEach(m => {
        m.style.display = "none"
    })

    document.getElementById("login").style.display = "block"
}


function openStatModal(event, task){
    event.preventDefault()

    access_token = localStorage.getItem("rusEGE_access_token")
    console.log(access_token)
    if (access_token === null){
        openLoginForm(event)
    }
    else{
        fetch(`/api/task/${task}/stat`).then(response => {
            if (response.status == 200){
                response.json().then(response => {
                    console.log(response)
                })
            }
        })
    }
}

function setAuthToken(access_token, refresh_token){
    localStorage.setItem('rusEGE_access_token', access_token)
    localStorage.setItem('rusEGE_refresh_token', refresh_token)
}

function closeModal(modal){
    modal.style.display = "none"
}