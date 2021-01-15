package vector

import (
	"testing"
)

func TestPointAdd(t *testing.T) {
	tests := []struct {
		point Point
		vector Vector
		expected Point
	}{
		{
			point: NewPoint(0, 0, 0),
			vector: NewVector(1, 1, 1),
			expected: NewPoint(1, 1, 1),
		},
	}
	for _, tt := range tests {
		actual := tt.point.Add(tt.vector)
		if actual.X() != tt.expected.X() {
			t.Errorf("Actual %v did not match Expected %v", actual, tt.expected)
		}
		if actual != tt.expected {
			t.Errorf("two different pointers, silly")
		}
	}
}