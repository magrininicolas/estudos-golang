package main

import "fmt"

type Scaler interface {
	Scale(n int)
}

type Coordinate struct {
	x, y int
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c *Coordinate) Scale(n int) {
	c.x *= n
	c.y *= n
}

func alteraMap(mapa map[string]int) {
	mapa["nicolas"] = 10
}

func alteraSlice(a []int) {
	a[0] = 999
}

func main() {
	// Map
	mapa := map[string]int{"nicolas": 1}
	alteraMap(mapa)
	fmt.Println(mapa["nicolas"])

	// Slice
	a := []int{1}
	alteraSlice(a)
	fmt.Println(a)

	// Struct
	c1 := Coordinate{2, 4}
	c2 := Coordinate{5, 6}
	coords := []Scaler{&c1, &c2}

	for _, v := range coords {
		fmt.Println(v)
		v.Scale(2)
		fmt.Println(v)
	}
}
