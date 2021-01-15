package vector

type Point interface {
	Add(vector Vector) Point
	X() float64
}

type Vector interface{

}

type Tuple struct {
	x, y, z float64
}

func NewPoint(x, y, z float64) Point {
	return &Tuple{x, y, z}
}

func NewVector(x, y, z float64) Vector {
	return Tuple{x, y, z}
}

func (t *Tuple) X() float64 {
	return t.x
}

func (t *Tuple) Add(v Vector) Point {
	t1 := v.(Tuple)

	return &Tuple{
		x: t.x + t1.x,
		y: t.y + t1.y,
		z: t.z + t1.z,
	}
}