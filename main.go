package main

import "fmt"

func main() {
	a := Puntos{}
	b := Punto{3, 2}
	c := Punto{4, 2}
	d := Punto{6, 2}
	e := Punto{1, 2}

	a = append(a, b, c, d, e)

	for i, value := range a {
		punto := &a[i]
		if (*punto).x > (*punto).y {
			punto.swap(value.x, value.y)
		}
		fmt.Printf("%+v", *punto)
	}
}
