package punto

type Punto struct {
	X int
	Y int
}
type Puntos []Punto

func (pointerToPunto *Punto) Swap(x, y int) {
	(*pointerToPunto).X, (*pointerToPunto).Y = y, x
}
