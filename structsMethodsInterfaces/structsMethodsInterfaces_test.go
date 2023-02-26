package perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("gets perimiter", func (t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
	
		if got != want {
			t.Errorf("got %.2f want %2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{shape: Rectangle{width: 12, height: 6}, want: 72},
		{shape: Circle{radius: 10}, want: 314.1592653589793},
		{shape: Triangle{12,6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}