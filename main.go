package main

import (
	"fmt"

	"integrador.com/punto"
)

func main() {
	a := punto.Puntos{}
	b := punto.Punto{3, 2}
	c := punto.Punto{4, 2}
	d := punto.Punto{6, 2}
	e := punto.Punto{1, 2}

	a = append(a, b, c, d, e)

	for i, value := range a {
		p := &a[i]
		if (*p).X > (*p).Y {
			p.Swap(value.X, value.Y)
		}
		fmt.Printf("%+v", *p)
	}
}
