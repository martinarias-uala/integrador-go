package main

import (
	"fmt"

	"integrador.com/punto"
)

func swap(x, y *int) {
	*x, *y = *y, *x
}
func main() {
	a := punto.Puntos{}
	b := punto.Punto{X: 3, Y: 2}
	c := punto.Punto{X: 4, Y: 2}
	d := punto.Punto{X: 6, Y: 2}
	e := punto.Punto{X: 1, Y: 2}

	a = append(a, b, c, d, e)

	for i, _ := range a {
		p := &a[i]
		if (*p).X > (*p).Y {
			swap(&p.X, &p.Y)
		}
		fmt.Printf("%+v", *p)
	}
}
