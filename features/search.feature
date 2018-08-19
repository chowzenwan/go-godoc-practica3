
Feature: Funcionalidad de búsqueda de productos
  Scenario: Búsqueda y agregar al carrito un producto
    Given Se carga la página de búsqueda
    When Buscamos el término "skirt"
    When Seleccionamos el modo lista
    And agregamos al carrito
    Then Debemos ver el siguiente mensaje "There is 1 item in your cart."


