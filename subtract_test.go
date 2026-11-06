package calculator

import "testing"

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive operands", 5, 3, 2},
		{"zero is right identity", 7, 0, 7},
		{"equal operands cancel", 4, 4, 0},
		{"both operands negative", -2, -3, 1},
		{"negative minus positive", -7, 4, -11},
		{"positive minus negative", 7, -4, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.a, tt.b); got != tt.want {
				t.Errorf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}