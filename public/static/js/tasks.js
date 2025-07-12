function startTraining(task){
    var ruleIds = []
    document.getElementById(`task${task}container`).querySelectorAll("input[type=checkbox]").forEach(input => {
        if (input.checked){
            ruleIds.push(input.value)
        }
    })

    const url = `/task/${task}?task=${task}&rule_ids=${ruleIds.join('&rule_ids=')}`
    window.location.href = url
}

function loadRulesStatHint(elem){
    if (localStorage.getItem("rusEGE_access_token") === null){
        elem.querySelector(".text").textContent = "Для того, чтобы увидеть статистику по правилам внутри задания войдите в аккаунт"
    } else{
        taskNum = elem.getAttribute("data-task")
        authRetry(getRulesStatAPI)(taskNum).then(response => {
            let hintHTML = `Сколько ошибок вы допустили по темам. Это поможет вам заострить внимание на важных моментах<br><br>`;
            console.log(response)
            response.data.stat.forEach(r => {
                hintHTML += `<span>${r.rule} - ${r.errors}</span><br>`
            })

            elem.querySelector(".text").innerHTML = hintHTML
        })
    }
}

function getTasks(){
    getTasksAPI().then(response => {
        tasks = response.data.tasks
        console.log(tasks)

        let hintHTML = `Сколько ошибок вы допустили по темам. Это поможет вам заострить внимание на важных моментах<br><br>`;
        let tasksHTML = ``;

        tasks.forEach(task => {
            let rulesHTML = ``;
            task.rules.forEach(r => {
                rulesHTML += `<div class="checkbox">
                    <label>
                        <input type="checkbox" checked value="${r.id}">
                        <span>${r.rule}</span>
                    </label>
                </div>`

                hintHTML += `<span>${r.rule} - ${r.errors}</span><br>`
            })
            tasksHTML += `
                <div id="task${task.number}container">
                    <div class="task-title">
                        <div>
                            <h3>Задание ${task.number}</h3>
                            <div class="task-description">
                                <span>${ task.description }</span>
                                <div data-task="${task.number}" class="context-help">
                                    <div class="text">
                                        ${hintHTML}
                                    </div>

                                    <img src="/static/images/ico_localhelp.png" />
                                </div>
                            </div>
                        </div>

                        <div>
                            <a href="#" onclick="openStatModal(event, ${task.number})">Узнать статистику</a>
                            <button class="button" onclick="startTraining(${task.number})">Начать тренировку</button>
                        </div>
                    </div>

                    <ul class="rules">
                        ${rulesHTML}
                    </ul>
                </div>`
        });
        document.querySelector("ul").innerHTML = tasksHTML

        document.querySelectorAll(".context-help").forEach(i => {
            i.addEventListener("mouseenter", () => loadRulesStatHint(i))
        })
    })
}

document.addEventListener('DOMContentLoaded', function() {
    checkIsAuth()
    getTasks()
    getSeo()
})
