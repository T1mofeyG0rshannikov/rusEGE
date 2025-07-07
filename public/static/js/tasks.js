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

document.addEventListener('DOMContentLoaded', function() {
    function getTasks(){
        fetch("/api/tasks/get").then(response => {
            if (response.status == 200){
                response.json().then(response => {
                    tasks = response.tasks
                    console.log(tasks)

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
                        })
                        tasksHTML += `
                            <div id="task${task.number}container">
                                <div class="task-title">
                                    <div>
                                        <h3>Задание ${task.number}</h3>
                                        <span>${ task.description }</span>
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
                })
            }
        })
    }
    getTasks()
    getSeo()
})
