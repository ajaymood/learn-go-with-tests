package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("test rectangle perimeter", func(t *testing.T) {
		r := Rectangle{50.0, 10.0}
		got := Perimeter(r)
		want := 120.0
		assertCorrectness(t, got, want)
	})

}

func TestArea(t *testing.T) {
	// checkArea := func(t testing.TB, s Shape, want float64) {
	// 	t.Helper()
	// 	got := s.Area()
	// 	if got != want {
	// 		t.Errorf("got: %g want: %g", got, want)
	// 	}
	// }
	// t.Run("test rectangle area", func(t *testing.T) {
	// 	r := Rectangle{10.0, 20.0}
	// 	want := 200.0
	// 	checkArea(t, r, want)
	// })

	// t.Run("test circle area", func(t *testing.T) {
	// 	c := Circle{10}
	// 	want := 314.1592653589793
	// 	checkArea(t, c, want)
	// })

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{10.0, 20.0}, want: 200.0},
		{name: "circle", shape: Circle{10.0}, want: 314.159265358979},
		{name: "triangle", shape: Triangle{4.0, 8.0}, want: 16.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got: %g want: %g", tt.shape, got, tt.want)
			}
		})
	}
}

func assertCorrectness(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got: %.2f want: %.2f", got, want)
	}
}
