package main

import "fmt"

type Poligono interface {
	Area() float64
	ImprimirArea()
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

type Cuadrado struct {
	Lado float64
}

func (t Triangulo) Area() float64 {
	return 0.5 * t.Base * t.Altura
}

func (t Triangulo) ImprimirArea() {
	fmt.Println("Area del Triangulo:", t.Area())
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func (r Rectangulo) ImprimirArea() {
	fmt.Println("Area del Rectangulo:", r.Area())
}

func (c Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (c Cuadrado) ImprimirArea() {
	fmt.Println("Area del Cuadrado:", c.Area())
}

func areaCalculada(p Poligono) {
	p.ImprimirArea()
}

func main() {
	areaCalculada(Triangulo{Base: 10, Altura: 5})
	areaCalculada(Rectangulo{Base: 10, Altura: 5})
	areaCalculada(Cuadrado{Lado: 10})
}
