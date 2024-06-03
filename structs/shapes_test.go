package structs

import "testing"

func assert(t testing.TB, shape Shape, want float64) {
	t.Helper()

	got := shape.Area()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("shapes", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{2, 2}, 4},
			{Circle{10}, 314.1592653589793},
		}

		for _, tt := range areaTests {
			assert(t, tt.shape, tt.want)
		}
	})
}
