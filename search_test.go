package practica3

import (
	"github.com/DATA-DOG/godog"
	"practica3/pageobjects"
	"fmt"
	"regexp"
	"strings"
	"github.com/DATA-DOG/godog/gherkin"
)

//var Driver selenium.WebDriver


func seCargaLaPginaDeBsqueda() (err error){
	pageobjects.LoadPage("http://automationpractice.com/index.php/")

	return nil
}

func buscamosElTrmino(product string) (err error) {
	pageobjects.WriteText(product)
	pageobjects.ClickSearch()

	return nil
}

func seleccionamosElModoLista() (err error) {
	pageobjects.ClickListView()

	return nil
}

func AgregamosAlCarrito() (err error) {
	pageobjects.AddToCart()

	return nil
}

func debemosVerElSiguienteMensaje(message string) (err error) {
	cartMessage := pageobjects.GetAddToCartMessage()

	if cartMessage != message {
		return fmt.Errorf("Esperaba: %v - Obtenido: %v", message, cartMessage)
	}

	return nil
}


func FeatureContext(s *godog.Suite) {
	s.Step(`^Se carga la página de búsqueda$`, seCargaLaPginaDeBsqueda)
	s.Step(`^Buscamos el término "([^"]*)"$`, buscamosElTrmino)
	s.Step(`^Seleccionamos el modo lista$`, seleccionamosElModoLista)
	s.Step(`^agregamos al carrito$`, AgregamosAlCarrito)
	s.Step(`^Debemos ver el siguiente mensaje "([^"]*)"$`, debemosVerElSiguienteMensaje)


	s.BeforeScenario(func(interface{}) {
		pageobjects.InitWebDriver()
	})

	s.AfterScenario(func(i interface{}, e error) {
		sc := i.(*gherkin.Scenario)
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "-"))

		pageobjects.TakeScreenShot(fileName)

		pageobjects.ClosePage()

	})
}
