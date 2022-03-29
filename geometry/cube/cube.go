package cube

import (
	"errors"

	"math"
)

type Cube struct {
	Lenght float64
}

func New(n float64) (*Cube, error) {
	if n < 0 {
		return nil, errors.New("could not initialized cube")
	}

	return &Cube{n}, nil
}

func (c *Cube) Name() string {
	return "cube"
}

func (c *Cube) Volume() float64 {
	return math.Pow(c.Lenght, 3)
}

func (c *Cube) Surface() float64 {
	return math.Pow(c.Lenght, 2) * 6
}

func (c *Cube) Perim() float64 {
	return c.Lenght * 12
}
