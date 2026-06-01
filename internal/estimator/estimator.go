package estimator

import "errors"

// WallArea обчислює площу стіни
func WallArea(height, width float64) float64 {
	return height * width
}

// BricksNeeded обчислює кількість цегли
func BricksNeeded(wallArea, brickArea float64) (float64, error) {
	if brickArea == 0 {
		return 0, errors.New("площа цеглини не може бути нульовою")
	}
	return wallArea / brickArea, nil
}
