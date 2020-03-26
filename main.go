package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/mercadolibre/go-practice-meli/lib"
	"golang.org/x/tour/pic"
)

func main() {
	//fmt.Println(RaizCuadrada(10, 10))
	//testPointer()
	//workWithSlice()
	//showPic()
	//showWordCount()
	//probaranonima()
	//testClosures()
	//testFibonacci()
	//lib.TestCuadrado()
	//lib.TestRectangulo()
	//lib.TestSqrt()
	//lib.TestChannelYRoutine()
	lib.TestTreeCompare()
	//lib.TestMutex()
}

func sumaYMultiplicacion(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}

// RaizCuadrada calcula la raiz cuadrada y con la precision podemos pedir que sea mas exacto
func RaizCuadrada(numero float64, precision int) float64 {
	z := numero

	for i := 0; i < precision; i++ {
		z -= ((z * z) - numero) / (2 * z)
	}

	return z
}

// RaizGo calcula la raiz cuadrada usando una funcion de go
func RaizGo(numero float64) float64 {
	return math.Sqrt(numero)
}

func testPointer() {
	var i uint8 = 10
	var j uint8 = 20
	p := &i // Puntero a una variable que contiene un  entero
	fmt.Printf("La variable: 'P' - Apunta a la dirección: %v - Que tiene valor: %v\n", p, *p)

	*p = j      // Cambio el valor de la variable a la que apunta p
	p = &j      // Ahora p en lugar de apuntar a la variable i, apunta a la variable j
	*p = *p * 2 // Modifico la variable j desde p

	fmt.Printf("La variable: 'P' - Apunta a la dirección: %v - Que tiene valor: %v\n", p, *p)
}

func workWithSlice() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0] // Se vacia el slice pero tambien todavia conserva sus datos.
	printSlice(s)

	// Extend its length.
	s = s[:cap(s)]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func showPic() {
	matrix := Pic(5, 5)
	for _, v := range matrix {
		fmt.Println(v)
	}

	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	container := make([][]uint8, dy)

	for x := range container {
		container[x] = make([]uint8, dx)

		for y := range container[x] {
			container[x][y] = xElevadoAY(x, y)
		}
	}

	return container
}

func xMasYDividido2(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func xPorY(x, y int) uint8 {
	return uint8(x * y)
}

func xElevadoAY(x, y int) uint8 {
	return uint8(math.Pow(float64(x), float64(y)))
}

func showWordCount() {
	fmt.Println(wordCount("SAD SAD qwe kjk qwe kjk"))
}

func wordCount(s string) map[string]int {
	words := strings.Fields(s)
	mapa := make(map[string]int)

	for _, word := range words {
		_, ok := mapa[word]
		if ok {
			mapa[word]++
		} else {
			mapa[word] = 1
		}
	}

	return mapa
}

func operarAnonima(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func probaranonima() {
	unaFn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x, 3) + math.Pow(y, 3))
	}

	fmt.Println(operarAnonima(unaFn))
	fmt.Println(operarAnonima(math.Pow))
}

func funcClosure() func(x int) int {
	var suma int

	return func(x int) int {
		suma += x
		return suma
	}
}

func testClosures() {
	positivo := funcClosure()
	negativo := funcClosure()

	for i := 0; i < 15; i++ {
		fmt.Printf("Positivo: %v - Negativo: %v\n", positivo(i), negativo(-i))
	}
}

// 0 1 1 2 3 5 8 13
func fibonacci() func() int {
	var prevValue int = 0 // 0	- 1
	var lastValue int = 0 // 1	- 1

	return func() int {
		if lastValue == 0 { // 0 0 => 0
			lastValue = 1
			return 0
		} else if prevValue == 0 && lastValue == 1 { // 0 1 => 1
			prevValue = 1
			return 1
		} else { // 1 1 => 2
			newValue := prevValue + lastValue

			prevValue = lastValue
			lastValue = newValue

			return lastValue
		}
	}
}

func testFibonacci() {
	fibo := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v\n", fibo())
	}
}

type MyInteger int

func (i MyInteger) Abs() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}

func VariadicSum(nums ...int) int {
	var suma int
	for _, v := range nums {
		suma += v
	}
	return suma
}

func testVariadicsum() {
	fmt.Println(VariadicSum(1))
	fmt.Println(VariadicSum(1, 2))
	fmt.Println(VariadicSum(1, 2, 3))

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(VariadicSum(numeros...))
}
