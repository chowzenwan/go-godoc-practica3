# go-godog-practica3

Primer intento de hacer BBD y POM usando Godog! :D

Run :
    go get github.com/DATA-DOG/godog/cmd/godog

O ponerlo en vendor(leer en godog, no lo he probado)


Descargar Chrome Driver yo lo coloque en la carpeta del proyecto(investigando sobre el tema)
http://chromedriver.chromium.org/downloads

Descargar Selenium Standalone Server
https://www.seleniumhq.org/download/

Run:
    java --jar selenium-server-standalone-3.13.0.jar

GODOG

Run:
    godog
or
    ../../bin/godog


PARA HACER UN REPORTE

Generar un reporte en json

Run:
    godog --format=cucumber >log/report.json

Usar el de cucumber js:
npm install cucumber-html-reporter --save-dev


Run:
    node log/reporter.js

