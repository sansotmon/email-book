package test

import "fmt"

func funcionalgo(a int) (c, d int) {
	return a, a
}

type car struct {
	brand string
	year  int
}

func Test() {
	const pi float64 = 3.14
	const pi2 = 3.1223
	fmt.Println("pi", pi)

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	value, value2 := funcionalgo(2)

	fmt.Println(value, value2)

	defer fmt.Println(areaCuadrado)

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	var array [6]int
	array[1] = 3
	fmt.Println(array, cap(array), len(array))

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	fmt.Println(slice[1:2])

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	m := make(map[string]int)
	m["Josesito"] = 14
	fmt.Println(m)

	valuemake := m["Josesito"]
	fmt.Println(valuemake)

	myCar := car{brand: "Chevy", year: 1554}
	fmt.Println(myCar)

	var ohertCAr car
	ohertCAr.brand = "Focus"
	fmt.Println(ohertCAr)
}
