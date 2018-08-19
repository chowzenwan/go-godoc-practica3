# go-godog-practica3

Primer intento de hacer BBD y POM usando Godog! :D

Run :
    go get github.com/DATA-DOG/godog/cmd/godog

O ponerlo en vendor


Descargar Chrome Driver yo lo coloque en la carpeta del proyecto(investigando sobre el tema)
http://chromedriver.chromium.org/downloads

Run: 
    godog --format=cucumber >log/report.json
    
    

Para hacer un reporte

Usar el de cucumber js:
npm install cucumber-html-reporter --save-dev

Run:
    node log/reporter.js
    






