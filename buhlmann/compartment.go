package buhlmann

import "math"

type Compartment struct {
	HN2  float64
	AN2  float64
	BN2  float64
	PPN2 float64
	HHe  float64
	AHe  float64
	BHe  float64
	PPHe float64
}

func (c *Compartment) AddStop(ata float64, time float64) {
	if c.PPN2 == 0 {
		c.PPN2 = 0.79
	}

	pBegin := 0.79
	pGas := c.PPN2 * ata
	te := time
	tht := c.HN2
	pComp := pBegin + (pGas-pBegin)*(1-math.Pow(2, -te/tht))

	c.PPN2 = pComp
}

func toATAs(depth float64) float64 {
	return (depth * 0.030242348200148) + 1
}

func toATAf(depth float64) float64 {
	return (depth * 0.02949893905749) + 1
}
