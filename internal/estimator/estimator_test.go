package estimator

import "testing"

func TestWallArea(t *testing.T) {
	tests := []struct {
		name     string
		height   float64
		width    float64
		expected float64
	}{
		{"standard wall", 3.0, 4.0, 12.0},
		{"zero dimensions", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WallArea(tt.height, tt.width)
			if result != tt.expected {
				t.Errorf("WallArea() = %v; очікувалося %v", result, tt.expected)
			}
		})
	}
}

func TestBricksNeeded(t *testing.T) {
	tests := []struct {
		name        string
		wallArea    float64
		brickArea   float64
		expected    float64
		expectError bool
	}{
		{"normal calculation", 10.0, 0.5, 20.0, false},
		{"zero brick area", 10.0, 0.0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := BricksNeeded(tt.wallArea, tt.brickArea)
			if (err != nil) != tt.expectError {
				t.Errorf("BricksNeeded() err = %v, очікувалася помилка: %v", err, tt.expectError)
				return
			}
			if !tt.expectError && result != tt.expected {
				t.Errorf("BricksNeeded() = %v; очікувалося %v", result, tt.expected)
			}
		})
	}
}
