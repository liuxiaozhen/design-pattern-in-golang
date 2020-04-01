package simplefactory

import (
	"errors"
)

type Shape interface {
	draw() string
}

func NewShape(shapeType string) (Shape, error) {
	if shapeType == "Rectangle" {
		return new(Rectangle), nil
	} else if shapeType == "Square" {
		return new(Square), nil
	} else if shapeType == "Circle" {
		return new(Circle), nil
	}
	return nil, errors.New("shapeType error")
}

type Rectangle struct{}

func (*Rectangle) draw() string {
	return "draw a Rectangle"
}

type Square struct{}

func (*Square) draw() string {
	return "draw a Square"
}

type Circle struct{}

func (*Circle) draw() string {
	return "draw a Circle"
}
