async function refreshTokens(){
    const response = await fetch(`/api/refresh-token/${localStorage.getItem("rusEGE_refresh_token")}`, {
        method: "POST",
    })
    
    let data = null
    try{
        data = await response.json();
    }
    catch {

    }
    
    console.log(data)
    setAuthToken(data.access_token, data.refresh_token)
}


function authRetry(func) {
  return async function (...args) {
    const response = await func.apply(this, args);

    if (response.status === 401){
        if (response.data.error.includes("expired")){
            await refreshTokens()
            const response = await func.apply(this, args);

            return response
        }
    }

    return response
  };
}

async function getData(response){
    let data = null
    try{
        data = await response.json();
    }
    catch {

    }
    
    return {
      status: response.status,
      data: data,
    };
}

async function createErrorAPI(word){
    const response = await fetch(`/api/word-error/create`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
            "Authorization": localStorage.getItem("rusEGE_access_token")
        },
    body: JSON.stringify({word: word.original})})
    
    return await getData(response)
}

async function deleteUserErrorAPI(word){
    const response = await fetch(`/api/word-error/delete`, {
            method: "DELETE",
            headers: {
                'Content-Type': 'application/json',
                "Authorization": localStorage.getItem("rusEGE_access_token")
            },
        body: JSON.stringify({word_id: word.id})
    })

    return await getData(response)
}


async function getTaskStatAPI(task){
    const response = await fetch(`/api/task/${task}/stat`, {
        headers: {
            'Authorization': localStorage.getItem("rusEGE_access_token")
        }
    })

    return await getData(response)
}

async function getWordsAPI(task){
    const response = await fetch(`/api/words/get?${task}`, {
        method: "GET",
        headers: {
            'Authorization': localStorage.getItem("rusEGE_access_token")
        }
    })

    return await getData(response)
}


async function getRulesStatAPI(taskNum){
    const response = await fetch(`api/rules/get-stat/${taskNum}`, {
        method: "GET",
        headers: {
            'Authorization': localStorage.getItem("rusEGE_access_token")
        }
    })

    return await getData(response)
}

async function getTasksAPI(){
    const response = await fetch("/api/tasks/get")

    return await getData(response)
}

async function getTaskUserWordsAPI(task){
    const response = await fetch(`/api/user-words/get/${task}`, {
        method: "GET",
        headers: {
            'Authorization': localStorage.getItem("rusEGE_access_token")
        }
    })

    return await getData(response)
}

async function getTaskRulesAPI(task){
    const response = await fetch(`/api/rule/get/${task}`)
    return await getData(response)
}

async function deleteUserWordAPI(wordId){
    const response = await fetch(`/api/user-words/delete/${wordId}`, {
        method: "DELETE",
        headers: {
            'Authorization': localStorage.getItem("rusEGE_access_token")
        }
    })

    return await getData(response)
}

async function createUserWordAPI(task, rule, word, exception, letters){
    const response = await fetch(`/api/user-words/create`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem("rusEGE_access_token")
        },
        body: JSON.stringify({
            task: task,
            rule: rule,
            word: word,
            exception: exception,
            letters: letters
        })
    })

    return await getData(response)
}

async function isAuthAPI(){
   const response = await fetch(`/api/user/get`, {
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem("rusEGE_access_token")
        }
    })

    return await getData(response)
}

async function loginAPI(data){
    const response = await fetch("/api/login", {
        method: "POST",
        headers: {
            'Content-type': 'application/json'
        },
        body: data
    })
    
    return await getData(response)
}