package main

// VERSION EN ESPAÑOL 1.0

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Estructura Aritmetica para almacenar los números
type Aritmetica struct {
	a, b int
}

// Métodos para realizar las operaciones
func (arit *Aritmetica) sumar() int {
	return arit.a + arit.b
}

func (arit *Aritmetica) restar() int {
	return arit.a - arit.b
}

func (arit *Aritmetica) dividir() float64 {
	if arit.b == 0 {
		fmt.Println("Error: División por cero no permitida.")
		return 0
	}
	return float64(arit.a) / float64(arit.b)
}

func (arit *Aritmetica) multiplicar() int {
	return arit.a * arit.b
}

// Función principal
func main() {
	var aritmetica Aritmetica

	fmt.Println("BIENVENIDO A LA CALCULADORA!")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nPOR FAVOR, SELECCIONE UNA OPCION:")
		fmt.Println("1. SUMAR")
		fmt.Println("2. RESTAR")
		fmt.Println("3. DIVIDIR")
		fmt.Println("4. MULTIPLICAR")
		fmt.Println("5. SALIR")

		// Leer la opción del usuario
		scanner.Scan()
		opcionInput := scanner.Text()
		opcion, err := strconv.Atoi(opcionInput)
		if err != nil {
			fmt.Println("SOLO SE PUEDE DIGITAR NUMEROS!")
			continue
		}

		if opcion > 5 || opcion < 1 {
			fmt.Println("OPCION INCORRECTA, POR FAVOR, DIGITE UNA OPCION VALIDA!")
			continue
		}

		if opcion == 5 {
			fmt.Println("GRACIAS POR UTILIZAR LA CALCULADORA 2! :)")
			break
		}

		// Leer los dos números
		fmt.Print("DIGITE EL PRIMER NUMERO: ")
		scanner.Scan()
		aritmetica.a, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("SOLO SE PUEDE DIGITAR NUMEROS!")
			continue
		}

		fmt.Print("DIGITE EL SEGUNDO NUMERO: ")
		scanner.Scan()
		aritmetica.b, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("SOLO SE PUEDE DIGITAR NUMEROS!")
			continue
		}

		// Realizar la operación seleccionada
		switch opcion {
		case 1:
			fmt.Printf("EL RESULTADO DE LA SUMA ES: %d\n", aritmetica.sumar())
		case 2:
			fmt.Printf("EL RESULTADO DE LA RESTA ES: %d\n", aritmetica.restar())
		case 3:
			result := aritmetica.dividir()
			if result != 0 {
				fmt.Printf("EL RESULTADO DE LA DIVISION ES: %.2f\n", result)
			}
		case 4:
			fmt.Printf("EL RESULTADO DE LA MULTIPLICACION ES: %d\n", aritmetica.multiplicar())
		}
	}
}
