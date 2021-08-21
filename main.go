package main

import (
	"fmt"
	Ord "sort"
	tiempo "time"
)

type ByInt []int

func (a ByInt) Len() int           { return len(a) }
func (a ByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInt) Less(i, j int) bool { return a[i] < a[j] }

func Mayor(n []int8) int8 {
	var x int8 = 0
	for i := 0; i < len(n); i++ {
		if x < n[i] {
			x = n[i]
		}
	}
	return x
}
func BusquedaBinaria(n []int, buscando int, a int, b int) (string, int) {
	if a < b {
		return "No se encuentra", -1
	}

	med := (a + b) / 2

	valor := n[med]
	if valor == buscando {
		return "Esta en la posicion =>", med
	}
	if valor > buscando {
		a = med
	} else {
		b = med
	}

	if n[a] == buscando {
		return "Esta en la posicion =>", a
	}

	if n[b] == buscando {
		return "Esta en la posicion =>", b
	}
	if a > b {
		a--
	}
	b++

	return BusquedaBinaria(n, buscando, a, b)

}

func BusquedaBinariaARRAYNoOrden(n []int, buscando int, a int, b int) (string, int) {
	if a < b {
		return "No se encuentra", -1
	}

	med := (a + b) / 2

	valor := n[med]
	if valor == buscando {
		return "Esta en la posicion =>", med
	}
	alto := n[a]
	if alto == buscando {
		return "Esta en la posicion =>", a
	}
	bajo := n[b]
	if bajo == buscando {
		return "Esta en la posicion =>", b
	}
	if a > b {
		a--
	}
	b++

	return BusquedaBinariaARRAYNoOrden(n, buscando, a, b)

}

func main() {

	y := []int{5, 89, 63, 21, 1, 0, 47, 57, 65, 53, 32, 34, 23, 56, 34, 23, 6, 11}
	w := []int{5, 89, 63, 21, 1, 0, 47, 57, 65, 53, 32, 34, 23, 56, 34, 23, 6, 11}
	Ord.Sort(ByInt(y))
	s := tiempo.Now()
	fmt.Println(BusquedaBinariaARRAYNoOrden(w, 47, len(w)-1, 0))
	fmt.Println("Sin Ordenar")
	f := tiempo.Now()
	elapsed := f.Sub(s)
	fmt.Println(elapsed)

	st := tiempo.Now()
	fmt.Println(BusquedaBinaria(y, 47, len(y)-1, 0))
	fmt.Println("Ordenado")
	fn := tiempo.Now()
	e := fn.Sub(st)
	fmt.Println(e)
}
