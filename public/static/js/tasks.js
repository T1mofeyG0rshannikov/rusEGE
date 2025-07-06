function startTraining(task){
    var ruleIds = []
    document.getElementById(`task${task}container`).querySelectorAll("input[type=checkbox]").forEach(input => {
        if (input.checked){
            ruleIds.push(input.value)
        }
    })

    console.log(ruleIds)

    const url = `/task/${task}?task=${task}&rule_ids=${ruleIds.join('&rule_ids=')}`
    console.log(url)
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
                                    <input type="checkbox" value="${r.id}">
                                    <span>${r.rule}</span>
                                </label>
                            </div>`
                        })
                        tasksHTML += `
                            <div id="task${task.number}container">
                                <button class="button" onclick="startTraining(${task.number})">Задание ${task.number}</button> - ${ task.description }
                                <ul>
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
