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

document.addEventListener('DOMContentLoaded', function() {
    getSeo()

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
