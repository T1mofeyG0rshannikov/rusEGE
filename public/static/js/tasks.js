document.addEventListener('DOMContentLoaded', function() {
    function getTasks(){
        fetch("/api/tasks/get").then(response => {
            if (response.status == 200){
                response.json().then(response => {
                    tasks = response.tasks
                    console.log(tasks)

                    let tasksHTML = ``;
                    tasks.forEach(task => {
                        tasksHTML += `<div><button class="button"><a href="/task/${task.number}">Задание ${task.number}</a></button> - ${ task.description }</div>`
                    });
                    document.querySelector("ul").innerHTML = tasksHTML
                })
            }
        })
    }
    getTasks()
    getSeo()
})