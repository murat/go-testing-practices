package sphere

import (
	"errors"
	"math"
)

type Sphere struct {
	Radius float64
}

func New(r float64) (*Sphere, error) {
	if r < 0 {
		return nil, errors.New("could not initialized sphere")
	}

	return &Sphere{r}, nil
}

func (s *Sphere) Name() string {
	return "Sphere"
}

func (s *Sphere) Volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(s.Radius, 3)
}

func (s *Sphere) Surface() float64 {
	return 4 * math.Pi * math.Pow(s.Radius, 2)
}

func (s *Sphere) Perim() float64 {
	return 2 * math.Pi * s.Radius
}
