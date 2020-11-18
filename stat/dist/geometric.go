package dist

import (
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// Geometric implements the geometric distribution, a distribution which models
// the number of failures observed before the first success in independent
// Bernoulli trials.
// See: https://en.wikipedia.org/wiki/Geometric_distribution
type Geometric struct {
	// p is the probability of success in a single trial.
	p float64

	backingUniform *distuv.Uniform
}

// NewGeometric returns a Geometric, representing a geometric random variable
// with parameter `p` and non-negative integer support.
func NewGeometric(p float64, src rand.Source) *Geometric {
	return &Geometric{
		p: p,
		backingUniform: &distuv.Uniform{
			Min: 0,
			Max: 1,
			Src: src,
		},
	}
}

// Rand returns a random sample drawn from the distribution.
// This function must return float64 in order to implement distuv.Rander.
func (g Geometric) Rand() float64 {
	// A geometric sample can be generated efficiently using the inversion method.
	return math.Floor(math.Log(g.backingUniform.Rand()) / math.Log(1.0-g.p))
}
