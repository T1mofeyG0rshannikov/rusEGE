function getWordErrors(){
    getWordErrorsAPI().then(response => {
        stat = response.data.stat
        console.log(response)

        let tasksHTML = '';
        stat.forEach(task => {
            if (task.words != null){
                let wordsHTML = ``;
                task.words.forEach(w => {
                    wordsHTML += `<div>
                        ${w.word} - ${w.error_count}
                    </div>`
                })
                tasksHTML += `
                    <div id="task${task.number}container">
                        <div class="task-title">
                            <h3>Задание ${task.task}</h3>
                        </div>

                        <ul class="rules">
                            ${wordsHTML}
                        </ul>
                    </div>`
            }
        });
        document.querySelector("ul").innerHTML = tasksHTML
    })
}

document.addEventListener('DOMContentLoaded', function() {
    checkIsAuth()
    getWordErrors()
    getSeo()
})
