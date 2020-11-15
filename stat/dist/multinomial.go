package dist

import (
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

// Multinomial implements the multinomial distribution, a generalization of the
// binomial distribution. A multinomial sample consists of a variable number of
// independent samples, each having a fixed probability of being drawn into one
// of a fixed number of distinct categories.
// See: https://en.wikipedia.org/wiki/Multinomial_distribution.
type Multinomial struct {
	// N is the number of trials.
	N uint32

	// CategoryProb is a slice encoding the event probabilities. For each integer i,
	// CategoryProb[i] represents the probability of drawing a sample in the
	// i-th category.
	CategoryProb []float64

	Src rand.Source
}

// Rand returns a random sample drawn from the distribution. The return format
// is a map. The (key, value) pair present in the map indicates `value` number
// of samples drawn from the `key` category.
func (m *Multinomial) Rand() map[uint32]float64 {
	// A multinomial sample N_1, N_2, ... with n trials and event probabilities
	// p_1, p_2, ..., can be generated using a binomial source, using the fact that
	//   N_1 ~ Binomial(n, p_1) .
	// Conditionally upon N_1, then
	//   N_2 ~ Binomial(n - N_1, p_2/(1 - p_1)) .
	// Conditionally upon N_1 and N_2, then
	//   N_3 ~ Binomial(n - N_1 - N_2, p_3/(1 - p_1 - p_2)) ,
	// and so on.
	n := float64(m.N)
	currIndex := uint32(0)
	cumulative := float64(1.0)
	sample := make(map[uint32]float64)
	for n > 0 {
		next := distuv.Binomial{
			N:   n,
			P:   m.CategoryProb[currIndex] / cumulative,
			Src: m.Src,
		}
		if binom := next.Rand(); binom > 0 {
			sample[currIndex] = binom
			n = n - binom
		}
		cumulative = cumulative - m.CategoryProb[currIndex]
		currIndex++
	}
	return sample
}

// TODO: Implement LogProb.
// LogProb computes the natural logarithm of the value of the probability mass function at `x`.
func (m *Multinomial) LogProb(x []float64) float64 {
	return 0.0
}

// Prob computes the value of the probability mass function at `x`.
func (m *Multinomial) Prob(x []float64) float64 {
	return math.Exp(m.LogProb(x))
}

// Mean returns the mean vector of the distribution.
func (m *Multinomial) Mean() []float64 {
	x := make([]float64, len(m.CategoryProb))
	copy(x, m.CategoryProb)
	floats.Scale(float64(m.N), x)
	return x
}

// TODO: Implement CovarianceMatrix.
// CovarianceMatrix returns the covariance matrix of the distribution.
func (m *Multinomial) CovarianceMatrix(dst *mat.SymDense) {
	return
}
