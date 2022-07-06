//  Create a rectangle abstraction using struct.
package main

import (
	"errors"
	"fmt"
)

// Rectangle is a struct that represents a rectangle
type Rectangle struct {
	length int
	height int
}

func main() {
	// Get length and height from user
	var length, height int
	println("Enter length: ")
	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}
	println("Enter height: ")
	_, err = fmt.Scan(&height)
	if err != nil {
		panic(err)
	}

	// Create rectangle instance
	rectangle, err := NewRectangle(length, height)
	if err != nil {
		panic(err)
	}

	// Print area and circumference
	println("Area: ", rectangle.Area())
	println("Circumference: ", rectangle.Circumference())

}

// Area returns the area of the rectangle
func (r *Rectangle) Area() int {
	return r.length * r.height
}

// Circumference returns the circumference of the rectangle
func (r *Rectangle) Circumference() int {
	return 2 * (r.length + r.height)
}

// NewRectangle creates a new rectangle instance
func NewRectangle(length int, height int) (*Rectangle, error) {
	// check if length and height are valid
	if length < 0 || height < 0 {
		// return error if length or height is negative
		return nil, errors.New("Invalid length or height! Must be positive!")
	}

	// create new rectangle instance and return it with no error
	return &Rectangle{length, height}, nil
}
