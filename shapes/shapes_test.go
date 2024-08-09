package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	// Table driven tests example
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 40.0},
		{Circle{Radius: 10.0}, 62.83185307179586},
		{Square{Side: 4}, 16.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
	// Old non-table driven tests
	// rect := Rectangle{10.0, 10.0}

	// got := rect.Perimeter()
	// want := 40.0

	// if got != want {
	// 	t.Errorf("got %.2f want %.2f", got, want)
	// }
}

func TestArea(t *testing.T) {
	// Table driven tests example
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
		{"Square", Square{2}, 16.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	// Old non-table driven tests
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }
	// t.Run("rectangles", func(t *testing.T) {
	// 	rect := Rectangle{12, 6}
	// 	checkArea(t, rect, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
