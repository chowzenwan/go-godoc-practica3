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

	//Driver.Get("http://automationpractice.com/index.php/")

	return nil
}

func buscamosElTrmino(product string) (err error) {
	pageobjects.WriteText(product)
	pageobjects.ClickSearch()

	//searchField, err := Driver.FindElement(selenium.ByCSSSelector, "input#search_query_top")
	//if err != nil {
	//	return
	//}
	//searchField.SendKeys(product)
	//
	//btnSearch, err := Driver.FindElement(selenium.ByCSSSelector, "button[name='submit_search']")
	//if err != nil {
	//	return
	//}
	//Driver.SetImplicitWaitTimeout(time.Second * 5)
	//btnSearch.Click()

	return nil
}

func seleccionamosElModoLista() (err error) {
	pageobjects.ClickListView()

	//btnListView, err := Driver.FindElement(selenium.ByCSSSelector, "i.icon-th-list")
	//if err != nil {
	//	return
	//}
	////Driver.SetImplicitWaitTimeout(time.Second * 5)
	//btnListView.Click()

	return nil
}

func AgregamosAlCarrito() (err error) {
	pageobjects.AddToCart()

	//btnAddToCart, err := Driver.FindElement(selenium.ByCSSSelector, "a.ajax_add_to_cart_button")
	//if err != nil {
	//	return
	//}
	//
	//btnAddToCart.Click()

	return nil
}

func debemosVerElSiguienteMensaje(message string) (err error) {
	cartMessage := pageobjects.GetAddToCartMessage()

	//Driver.SetImplicitWaitTimeout(time.Second * 5)
	//cartMessage, err := Driver.FindElement(selenium.ByCSSSelector, "span[class*=ajax_cart_product_txt]:not([class*='unvisible'])")
	//if err != nil {
	//	return
	//}
	//
	//cartMessageInnerText, _ := cartMessage.GetAttribute("innerText")
	//cartMessageInnerText = strings.TrimSpace(cartMessageInnerText)
	////fmt.Println("este es el mensaje de la web:", cartMessageInnerText)
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
		//Driver = support.WDInit()
	})

	s.AfterScenario(func(i interface{}, e error) {
		sc := i.(*gherkin.Scenario)
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "-"))

		pageobjects.TakeScreenShot(fileName)
		//
		//shot, _ := Driver.Screenshot()
		//
		//support.SaveImage(shot, fileName)
		//
		//fmt.Println(fileName)

		//Driver.Quit()
		pageobjects.ClosePage()

	})
}
