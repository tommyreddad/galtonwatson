package galtonwatson

import (
	"golang.org/x/exp/rand"

	"github.com/tommyreddad/galtonwatson/stat/dist"
	"github.com/tommyreddad/galtonwatson/tree"
)

// GaltonWatson implements a conditioned Galton-Watson tree generator. The
// conditioned Galton-Watson tree is specified by an offspring distribution Prob
// and a node count N. The generated tree is the Galton-Watson tree with offspring
// distribution Prob conditioned to have N nodes.
type GaltonWatson struct {
	// N is the number of nodes in the tree.
	N float64

	// Prob is the probability mass function of the offspring distribution.
	Prob func(float64) float64

	Src rand.Source
}

// Rand returns a random sample drawn from the distribution. The generation
// algorithm is due to Devroye (2011), see: http://luc.devroye.org/gw-simulation.pdf.
func (gw GaltonWatson) Rand() *tree.Node {
	M := dist.Multinomial{
		N:    gw.N,
		Prob: gw.Prob,
		Src:  gw.Src,
	}

	n := uint32(gw.N)
	// Generate multinomial conditionally upon the total sum being n-1, by the rejection method.
	var Mcond map[float64]float64
	for total := uint32(0); total != n-1; {
		Mcond = M.Rand()
		total = 0
		for k, v := range Mcond {
			total += uint32(k) * uint32(v)
		}
	}

	// Compute the first pass at the Xi array of DFS-order offspring.
	Xi := make([]uint32, n)
	{
		i := 0
		for k, v := range Mcond {
			for j := uint32(0); j < uint32(v); j++ {
				Xi[i] = uint32(k)
				i++
			}
		}
	}
	rand.Shuffle(int(n), func(i, j int) {
		Xi[i], Xi[j] = Xi[j], Xi[i]
	})

	// Compute the S array and keep track of the first minimum encounter.
	S := make([]int32, n+1)
	S[0] = 1
	minIndex := uint32(1)
	for i := uint32(1); i < n+1; i++ {
		S[i] = S[i-1] + int32(Xi[i-1]) - 1
		if S[i] < S[minIndex] {
			minIndex = i
		}
	}

	// Rotate Xi according to Dvoretzky-Motzkin.
	Xi = append(Xi[minIndex:n], Xi[0:minIndex]...)

	// Build the tree using the correct Xi in DFS order.
	rootNode := tree.New(0)
	{
		nodeCount := 0
		traversalCount := 0
		currNode := rootNode
		dfsStack := []*tree.Node{currNode}
		for len(dfsStack) > 0 {
			currNode = dfsStack[len(dfsStack)-1]
			dfsStack = dfsStack[:len(dfsStack)-1]
			for i := uint32(0); i < Xi[traversalCount]; i++ {
				nodeCount++
				newNode := tree.New(nodeCount)
				currNode.AppendChild(newNode)
				dfsStack = append(dfsStack, newNode)
			}
			traversalCount++
		}
	}
	return rootNode
}
