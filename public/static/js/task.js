var progressBarPercent = 0;
var index = 0;
var maxSteps = 0;
var errors = 0;

function getErrorText(errorCount) {
    const titles = ['ошибка', 'ошибки', 'ошибок'];
    const cases = [2, 0, 1, 1, 1, 2];
    const index = (errorCount % 100 > 4 && errorCount % 100 < 20) ? 2 : cases[(errorCount % 10 < 5) ? errorCount % 10 : 5];
    return titles[index];
}

function updateProgressBar() {
    const percent = Math.max((index / maxSteps) * 100, progressBarPercent);
    progressBarPercent = percent;
    document.querySelector('.progressbar__value').style.width = percent + '%';
}

function getTaskUserWords(task){
    authRetry(getTaskUserWordsAPI)(task).then(response => {
        console.log(response)
        if (response.status === 200){
            let wordsHTML = ``;
            response.data.words.forEach(w => {
                wordsHTML += `<li><div>${w.word} <img class="basket" onclick="deleteUserWord(this, ${w.id})" src="/static/images/cart.png" /></div></li>`
            });

            
            const modal = document.getElementById("userwordsmodal")
            modal.querySelector("ul").innerHTML = wordsHTML
            modal.querySelector(".modal-content").innerHTML += `<img src="/static/images/plus.webp" class="plus" onclick="prepareAddUserWord(${task})" />`
            modal.querySelector(".task-num").textContent = task
        }
        else if (response.status === 401){
            openLoginForm()
        }
    })
}

function prepareAddUserWord(task){
    const modal = document.getElementById("userwordsmodal")
    const addSpace = document.createElement("div")

    authRetry(getTaskRulesAPI)(task).then(response => {
        console.log(response)
        const rules = response.data.rules;

        let rulesHTML = ``;

        rules.forEach(r => {
            rulesHTML += `<option value="${r.id}">${r.rule}</option>`
        })
    
        addSpace.innerHTML = `
            <select>${rulesHTML}</select>
            <span class="hint">Выберите правило</span><br>

            <input type="text" placeholder="Введите слово" /><br>

            <div><input type="checkbox" /> Исключение</div>

            <span class="hint">Введите слово, указав спорную букву заглавным символом. Если на месте пропуска нет буквы, укажите символ "-" </span>
            
            <input type="text" placeholder="Варианты" /><br>
            <span class="hint">Введите буквы, которые могут быть на месте пропуска через запятую. Если на месте пропуска нет буквы, укажите символ "-" </span>
    
            <button style="margin-left: 0;" onclick="addUserWord(${task})" class="button">Добавить</button>
        `;
        addSpace.classList.add("add-userword")
        modal.querySelector(".modal-content").appendChild(addSpace)
    
        modal.querySelector(".plus").remove()
    })
}

function addUserWord(task){
    const modal = document.getElementById("userwordsmodal")
    const word = modal.querySelector("input").value

    const letters = modal.querySelectorAll("input")[2].value
    const rule = modal.querySelector("select").value
    const exception = modal.querySelector("input[type=checkbox]").checked

    authRetry(createUserWordAPI)(task, +rule, word, exception, letters).then(response => {
        console.log(response)
        if (response.status === 200){
            const w = response.data.word
            modal.querySelector("ul").innerHTML += `<li><div>${w.word} <img class="basket" onclick="deleteUserWord(this, ${w.id})" src="/static/images/cart.png" /></div></li>`
            modal.querySelector(".add-userword").remove()
            modal.querySelector(".modal-content").innerHTML += `<img src="/static/images/plus.webp" class="plus" onclick="prepareAddUserWord(${task})" />`
        }
    })
}

function openUserwordsModal() {
    document.getElementById("userwordsmodal").style.display = "block"
    const task = window.location.pathname.split('/')[2]
    getTaskUserWords(task)
}

document.addEventListener('DOMContentLoaded', function() {
    getSeo()
    checkIsAuth()

    let words = []
    const wordErrors = new Set()

    function nextWord(){
        if (words.length == 1){
            document.querySelector(".word-wrapper").innerHTML = `
                <p class="success">Отлично!</p>
                <p style="font-size: 20px;">Вы допустили ${errors} ${getErrorText(errors)}</p>
                <div class="finish-buttons">
                    <button class="button" onclick="window.location.reload()">Ещё раз</button>
                    <button class="button" onclick="window.location.href='/'">На главную</button>
                </div>
            `
            return
        }
        words = words.slice(1)
        setWord()
    }

    function checkLetter(correct, element){
        const word = words[0];

        if (correct){
            index++;
            element.classList.add("correct")

            if (!wordErrors.has(word.word)){
                authRetry(deleteUserErrorAPI)(word).then(response => {
                    console.log(response)
                })
            }

            setTimeout(function(){
                document.querySelector("#options").classList.remove("shake")
                document.querySelectorAll(".option").forEach((o, ind) => {
                    o.classList.remove(word.options[ind].correct ? "correct" : "wrong")
                })
                nextWord()
            }, 1000)
        }
        if (!correct){
            element.classList.add("wrong")
            errors++;
            wordErrors.add(word.word)
            words.push(word)
            document.querySelector("#options").classList.add("shake")
        
            document.querySelectorAll(".option").forEach((o, ind) => {
                if (word.options[ind].correct){
                    o.classList.add("correct")
                }
            })
 
            authRetry(createErrorAPI)(word).then(response => {
                console.log(response)
            })

            function openDescription(word){
                const modal = document.getElementById("description")
                modal.style.display = "block"
                modal.querySelector(".description-content").textContent = word.description
            }

            setTimeout(function(){
                document.querySelector("#options").classList.remove("shake")
                document.querySelectorAll(".option").forEach((o, ind) => {
                    o.classList.remove(word.options[ind].correct ? "correct" : "wrong")
                })
                if (word.description){
                    openDescription(word)
                }
                else{
                    nextWord()
                }
            }, 1000)
        }

        updateProgressBar()
    }

    function getWords(){
        const url = window.location.href.split('?')[1]
        authRetry(getWordsAPI)(url).then(response => {
            words = response.data.words
            console.log(words)
            maxSteps = words.length
            setWord()
        })
    }
    getWords()

    function setWord(){
        const word = words[0]
        document.getElementById("word").textContent = word.word
        document.getElementById("rule").textContent = word.rule
        
        if (word.exception){
            document.getElementById("rule").textContent += ' (Исключение)'
        }

        let optionsHTML = ``;
        word.options.forEach(option => {
            optionsHTML += `<span class="option">${option.letter}</span>`
        });
        document.getElementById("options").innerHTML = optionsHTML
        document.querySelectorAll(".option").forEach((o, ind) => {
            o.addEventListener('click', function(event){
                checkLetter(word.options[ind].correct, this)
            })
        })
    }

    document.querySelector("#description").addEventListener("click", function (){
        document.getElementById("description").style.display = "none"
        nextWord()
    })
})

function deleteUserWord(img, wordId){
    authRetry(deleteUserWordAPI)(wordId).then(response => {
        if (response.status === 200){
            img.closest("li").remove()
        }
    })
}