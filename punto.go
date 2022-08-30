package main

type Punto struct {
	x int
	y int
}
type Puntos []Punto

func (pointerToPunto *Punto) swap(x, y int) {
	(*pointerToPunto).x, (*pointerToPunto).y = y, x
}
