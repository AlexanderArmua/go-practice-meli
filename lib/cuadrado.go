package lib

import (
	"fmt"
	"math"
)

type FormaGeometrica interface {
	Area() float64
	Perimetro() float64
	Escalar(float64)
	Print()
}

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) Area() float64 {
	return math.Pow(c.Lado, 2)
}

func (c *Cuadrado) Perimetro() float64 {
	return 4 * c.Lado
}

func (c *Cuadrado) Escalar(escala float64) {
	c.Lado *= escala
	fmt.Printf("Reescalado Cuadrado x %v\n", escala)
}

func (c *Cuadrado) Print() {
	fmt.Printf("Cuadrado: %v - Area: %v - Perimetro: %v\n", c, c.Area(), c.Perimetro())
}

type Rectangulo struct {
	Alto  float64
	Ancho float64
}

func (c *Rectangulo) Area() float64 {
	return c.Alto * c.Ancho
}

func (c *Rectangulo) Perimetro() float64 {
	return 2*c.Ancho + 2*c.Alto
}

func (c *Rectangulo) Escalar(escala float64) {
	c.Alto *= escala
	c.Ancho *= escala
	fmt.Printf("Reescalado Rectangulo x %v\n", escala)
}

func (c *Rectangulo) Print() {
	fmt.Printf("Rectangulo: %v - Area: %v - Perimetro: %v\n", c, c.Area(), c.Perimetro())
}

func GetArea(f FormaGeometrica) float64 {
	return f.Area()
}

func GetPerimetro(f FormaGeometrica) float64 {
	return f.Perimetro()
}

func SetEscalar(f FormaGeometrica, escala float64) {
	f.Escalar(escala)
}

func PrintForma(f FormaGeometrica) {
	f.Print()
}

func TestCuadrado() {
	var c FormaGeometrica = &Cuadrado{23.2}
	printYEscalar(c)
}

func TestRectangulo() {
	var r FormaGeometrica = &Rectangulo{23.2, 25.2}
	printYEscalar(r)
}

func printYEscalar(forma FormaGeometrica) {
	PrintForma(forma)
	SetEscalar(forma, 10)
	PrintForma(forma)
}
