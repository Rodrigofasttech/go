package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}

// criando um método para a struct quadrado
func (q quadrado) calculo() float64 {
	return (q.lado * q.lado)
}

type circulo struct {
	lado float64
}

func main() {
	math.Dim(1, 2)
	fmt.Println("teste")

	q1 := quadrado{10}
	fmt.Println(q1.calculo()) // chamado o método
}
