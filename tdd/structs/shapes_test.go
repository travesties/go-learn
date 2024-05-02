package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := &Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}

	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := &Rectangle{12.0, 6.0}
	//	checkArea(t, rectangle, 72.0)
	//})

	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//})

	// Table driven tests are helpful when testing around interfaces or
	// testing a unit that has many different cases.
	// https://go.dev/wiki/TableDrivenTests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		// Using t.Run with named test cases provides clarity around which tests
		// fail. If a test fails, the output will include the subtest name, e.g,
		// TestArea/Rectangle, and including the shape struct in the output string
		// makes it clear which test failed. This ensures that a developer can
		// quickly see which test failed.
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
