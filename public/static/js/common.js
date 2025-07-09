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
        authRetry(getTaskStatAPI)(task).then(response => {
            console.log(response)

            document.getElementById("taskStat").style.display = "block"
            
            let statHTML = ``;
            if (response.data.stat != null){
                for (word of response.data.stat){
                    statHTML += `<li>${word.word} - ${word.errors}</li>`
                }
            }
            document.getElementById("taskStat").querySelector("ul").innerHTML = statHTML         
        })
    }
}

function setAuthToken(access_token, refresh_token){
    localStorage.setItem('rusEGE_access_token', access_token)
    localStorage.setItem('rusEGE_refresh_token', refresh_token)
}

function closeModal(modal){
    modal.closest(".modal").style.display = "none"
}