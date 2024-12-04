package main

import (
	"bufio"
	"calculator/aritmetica"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var aritmetica aritmetica.Aritmetica
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nBIENVENIDO A LA CALCULADORA!")

	for {
		fmt.Println("\nPOR FAVOR, SELECCIONE UNA OPCION:")
		fmt.Println("1. SUMAR")
		fmt.Println("2. RESTAR")
		fmt.Println("3. DIVIDIR")
		fmt.Println("4. MULTIPLICAR")
		fmt.Println("5. SALIR")

		// Leer la opción del usuario
		opcion := leerEntero("OPCIÓN", scanner)
		if opcion == 5 {
			fmt.Println("\nGRACIAS POR UTILIZAR LA CALCULADORA!")
			break
		}

		// Leer los dos números
		aritmetica.A = leerEntero("PRIMER NÚMERO", scanner)
		aritmetica.B = leerEntero("SEGUNDO NÚMERO", scanner)

		// Realizar la operación seleccionada
		switch opcion {
		case 1:
			fmt.Printf("\nEL RESULTADO DE LA SUMA ES: %d\n", aritmetica.Sumar())
		case 2:
			fmt.Printf("\nEL RESULTADO DE LA RESTA ES: %d\n", aritmetica.Restar())
		case 3:
			result, err := aritmetica.Dividir()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\nEL RESULTADO DE LA DIVISIÓN ES: %.2f\n", result)
			}
		case 4:
			fmt.Printf("\nEL RESULTADO DE LA MULTIPLICACIÓN ES: %d\n", aritmetica.Multiplicar())
		default:
			fmt.Println("\nOPCIÓN INCORRECTA, POR FAVOR, DIGITE UNA OPCIÓN VÁLIDA!")
		}
	}
}

func leerEntero(mensaje string, scanner *bufio.Scanner) int {
	for {
		fmt.Printf("\nDIGITE LA %s: ", mensaje)
		scanner.Scan()
		valor, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("\nSOLO SE PUEDE DIGITAR NÚMEROS!")
		} else {
			return valor
		}
	}
}
