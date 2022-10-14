package structs

import "math"

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle...
type Circle struct {
  Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

// Triangle has te dimensions of a Triangle
type Triangle struct {
  Base float64
  Height float64
}

// Area returns the area of the Triangle
func (t Triangle) Area() float64 {
  return (t.Base * t.Height) * 0.5
}

