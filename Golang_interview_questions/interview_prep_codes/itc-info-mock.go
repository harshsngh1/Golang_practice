// Generate a sequence 1,2,4,7,11, 16, 22 ... 100

// generate squares for every number divisible by 5

// take all squares add on mat - routine

// compute area occupied by these squares

// print area

package main

import "fmt"

func main() {

	nums := generateNum()

	sq := genSquares(nums)

	mt := &Mat{}

	c := make(tShapeCh)

	go add(mt, c)

	for _, v := range sq {

		sh := &Square{float64(v)}

		c <- sh

	}

	for _, v := range mt.shapes {

		fmt.Println(v.getArea())

	}

}

type (
	Shape interface {
		getArea() float64
	}

	Square struct {
		side float64
	}

	Mat struct {
		shapes []Shape
	}

	tShapeCh chan Shape
)

func (s *Square) getArea() float64 {

	return s.side * s.side

}

func (m *Mat) place(shape Shape) {

	m.shapes = append(m.shapes, shape)

}

func generateNum() []int {

	res := make([]int, 0)

	current := 1

	for i := 1; i <= 10; i++ {

		res = append(res, current)

		current += i

		fmt.Println(current)

	}

	return res

}

func genSquares(data []int) []int {

	res := make([]int, 0)

	for _, v := range data {

		if v%5 == 0 {

			res = append(res, v*v)

		}

	}

	return res

}

func add(mt *Mat, data tShapeCh) {

	for c := range data {

		mt.place(c)

	}

}
