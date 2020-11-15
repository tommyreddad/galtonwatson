package dist

import (
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// Multinomial implements the multinomial distribution, a generalization of the
// binomial distribution. A multinomial sample consists of a variable number of
// independent samples, each having a fixed probability of being drawn into one
// of a fixed number of distinct categories.
// See: https://en.wikipedia.org/wiki/Multinomial_distribution.
type Multinomial struct {
	// N is the number of trials.
	N float64

	// Prob is a function encoding the event probabilities. For each integer i,
	// Prob(float64(i)) represents the probability of drawing a sample in the
	// i-th category.
	Prob func(float64) float64

	Src rand.Source
}

// Rand returns a random sample drawn from the distribution.
func (m Multinomial) Rand() map[float64]float64 {
	// A multinomial sample N_1, N_2, ... with n trials and event probabilities
	// p_1, p_2, ..., can be generated using a binomial source, using the fact that
	//   N_1 ~ Binomial(n, p_1) .
	// Conditionally upon N_1, then
	//   N_2 ~ Binomial(n - N_1, p_2/(1 - p_1)) .
	// Conditionally upon N_1 and N_2, then
	//   N_3 ~ Binomial(n - N_1 - N_2, p_3/(1 - p_1 - p_2)) ,
	// and so on.
	n := m.N
	currIndex := float64(0.0)
	cumulative := float64(1.0)
	sample := make(map[float64]float64)
	for n > 0 {
		next := distuv.Binomial{
			N:   n,
			P:   m.Prob(currIndex) / cumulative,
			Src: m.Src,
		}
		if binom := next.Rand(); binom > 0 {
			sample[currIndex] = binom
			n = n - binom
		}
		cumulative = cumulative - m.Prob(currIndex)
		currIndex++
	}
	return sample
}
