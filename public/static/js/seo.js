function getSeo(){
    fetch("/api/indexseo").then(response => {
        if (response.status == 200){
            response.json().then(response => {
                const seo = response.seo;
                console.log(seo)
                try{
                    document.querySelector("#image").src = seo.image
                }
                catch{
                    
                }

                try{
                    document.querySelector("#title").textContent = seo.title
                } catch {

                }
                try{
                    document.querySelector("#about").textContent = seo.about
                } catch{

                }
                
                try{
                    document.querySelector("#logo").src = seo.logo
                } catch{
                    
                }

                try{
                    document.querySelector("#fipi-link").href = seo.fipi_link
                } catch{
                    
                }
            })
        }
    })
}