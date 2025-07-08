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

    function checkLetter(correct, element){
        if (correct){
            index++;
            element.classList.add("correct")
        }
        if (!correct){
            element.classList.add("wrong")
            errors++;
            words.push(words[0])
            document.querySelector("#options").classList.add("shake")
        
            document.querySelectorAll(".option").forEach((o, ind) => {
                if (words[0].Options[ind].correct){
                    o.classList.add("correct")
                }
            })

            fetch(`/api/word-error/create`, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                    "Authorization": localStorage.getItem("rusEGE_access_token")
                },
            body: JSON.stringify({word_id: words[0].id})}).then(response => {
                    response.json().then(response => {
                        console.log(response)
                    })
            })
        }

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
        const url = window.location.href.split('?')[1]
        fetch(`/api/words/get?${url}`).then(response => {
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
        
        if (words[0].exception){
            document.getElementById("rule").textContent += ' (Исключение)'
        }

        let optionsHTML = ``;
        words[0].Options.forEach(option => {
            optionsHTML += `<span class="option">${option.letter}</span>`
        });
        document.getElementById("options").innerHTML = optionsHTML
        document.querySelectorAll(".option").forEach((o, ind) => {
            o.addEventListener('click', function(event){
                checkLetter(words[0].Options[ind].correct, this)
            })
        })
    }
})
