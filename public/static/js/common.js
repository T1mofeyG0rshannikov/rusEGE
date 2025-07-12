function openRegisterForm(event){
    event.preventDefault()
    document.querySelectorAll('.modal').forEach(m => {
        m.style.display = "none"
    })

    document.getElementById("register").style.display = "block"
}

function openLoginForm(){
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
            if (response.status == 200){
                if (response.data.stat != null){
                    for (word of response.data.stat){
                        statHTML += `<li>${word.word} - ${word.errors}</li>`
                    }
                    statHTML = `<ul>${statHTML}</ul>`
                } else{
                    statHTML += '<p>Поздравляю! Ты не совершил ни одной ошибки</p>'
                }
            }
            document.getElementById("taskStat").querySelector(".stat-content").innerHTML = statHTML         
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

function setAuthToNav(user){
    const profile = document.querySelector(".profile")
    if (user){
        profile.classList.add("auth")
        profile.querySelector(".logo").innerHTML = `<a>${user.username[0].toUpperCase()}</a>`
        
        const userNav = document.querySelector(".user-nav")
        userNav.querySelector(".title a").textContent = user.username
    } else{
        profile.classList.remove("auth")
        profile.querySelector(".logo").innerHTML = `<img src="/static/images/profile.png" />`
    }
}

function checkIsAuth(){
    authRetry(isAuthAPI)().then(response => {
        if (response.status === 200){
            console.log(response)
            const user = response.data.user
            setAuthToNav(user)
        }
    })
}

function logout(){
    localStorage.removeItem("rusEGE_access_token")
    localStorage.removeItem("rusEGE_refresh_token")

    setAuthToNav(null)
}