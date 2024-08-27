package main

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("expected %2f, got %2f", want, got)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{name: "circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v expected %g, got %g", tt.shape, tt.want, got)
			}
		})
	}

}
