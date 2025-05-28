package geomentry

import "math"

// Exported function
func CalculateCircleArea(radius float64) float64 {
    return math.Pi * math.Pow(radius, 2)
}

// Exported constant
const (
    MaxAttempts   = 5
    ErrorMessage = "An error occurred"
)

// Exported type
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) CalculateArea() float64 {
    return r.Width * r.Height
}
