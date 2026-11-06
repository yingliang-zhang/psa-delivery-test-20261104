package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive operands", 2, 3, 5},
		{"zero is additive identity", 7, 0, 7},
		{"positive and negative cancel", 5, -5, 0},
		{"both operands negative", -2, -3, -5},
		{"negative plus positive", -7, 4, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}