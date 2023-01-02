package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := Perimeter(r)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

//func TestArea(t *testing.T) {
//
//	// This is a helper function
//	// the helper is decoupled from the concrete types, it just needs the method
//	// Helper functions works well with interfaces
//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{12, 6}
//		checkArea(t, rectangle, 72.0)
//	})
//
//	t.Run("circles", func(t *testing.T){
//		circle := Circle{10.0}
//		checkArea(t, circle, 314.1592653589793)
//	})
//
//}

// Table Testing
func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	// By wrapping each case in a t.Run you will have clearer test
	// output on failures as it will print the name of the case
	for _, tt := range areaTests {
		// using tt.name from the case to use it as the t.Run test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
	// run specfic tests: go test -run TestArea/Rectangle
}
