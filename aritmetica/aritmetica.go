package aritmetica

import "fmt"

// Aritmetica estructura para almacenar los números
type Aritmetica struct {
	A, B int
}

// Métodos para realizar las operaciones
func (arit *Aritmetica) Sumar() int {
	return arit.A + arit.B
}

func (arit *Aritmetica) Restar() int {
	return arit.A - arit.B
}

func (arit *Aritmetica) Dividir() (float64, error) {
	if arit.B == 0 {
		return 0, fmt.Errorf("DIVISION POR 0 NO PERMITIDA!")
	}
	return float64(arit.A) / float64(arit.B), nil
}

func (arit *Aritmetica) Multiplicar() int {
	return arit.A * arit.B
}
