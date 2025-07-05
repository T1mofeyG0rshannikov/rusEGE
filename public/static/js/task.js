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

document.addEventListener('DOMContentLoaded', function() {
    getSeo()

    function updateProgressBar() {
        const percent = Math.max((index / maxSteps) * 100, progressBarPercent);
        progressBarPercent = percent;
        document.querySelector('.progressbar__value').style.width = percent + '%';
    }

    function checkLetter(correct){
        if (correct){
            index++;
        }
        if (!correct){
            errors++;
            words.push(words[0])
            document.querySelector("#options").classList.add("shake")
        }
        document.querySelectorAll(".option").forEach((o, ind) => {
            o.classList.add(words[0].Options[ind].correct ? "correct" : "wrong")
        })
        updateProgressBar()
        setTimeout(function(){
            document.querySelector("#options").classList.remove("shake")
            document.querySelectorAll(".option").forEach((o, ind) => {
                o.classList.remove(words[0].Options[ind].correct ? "correct" : "wrong")
            })
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
        }, 1000)
    }

    let words = []
    function getWords(){
        const taskNum = window.location.href.split('/')[4]

        fetch(`/api/words/get/${taskNum}`).then(response => {
            if (response.status == 200){
                response.json().then(response => {
                    words = response.words
                    console.log(words)
                    maxSteps = words.length
                    setWord()
                })
            }
        })
    }
    getWords()

    function setWord(){
        document.getElementById("word").textContent = words[0].word
        document.getElementById("rule").textContent = words[0].rule
        let optionsHTML = ``;
        words[0].Options.forEach(option => {
            optionsHTML += `<span class="option">${option.letter}</span>`
        });
        document.getElementById("options").innerHTML = optionsHTML
        document.querySelectorAll(".option").forEach((o, ind) => {
            o.addEventListener('click', function(event){
                checkLetter(words[0].Options[ind].correct)
            })
        })
    }
})
