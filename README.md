# galtonwatson

**[This module is under early development and should not be used by anyone at this time.]**

**galtonwatson** is a Go module containing efficient algorithms for generating (conditioned) Galton-Watson trees and computing some of their properties. By implication, this module can be used to generate uniformly random samples from many classes of trees, including:
- uniformly random binary trees of a given size,
- uniformly random *d*-ary trees of a given size,
- uniformly random unordered labeled trees of a given size,
- uniformly random ordered trees of a given size,
- etc.

## Import
```go
import "github.com/tommyreddad/galtonwatson"
```

## Usage

```go
import "github.com/tommyreddad/galtonwatson"

func main() {
    // Generates a uniformly random binary tree of size 100.
    T := galtonwatson.GaltonWatson{
        N: 100,
        Prob: func (i float64) float64 {
            if i == 0 {
                // 0 children with probability 1/4
                return 0.25
            }
            if i == 1 {
                // 1 child with probability 1/2
                return 0.5
            }
            if i == 2 {
                // 2 children with probability 1/4
                return 0.25
            }
            return 0.0
        }
    }
    T.Rand()
}
```

## References

- Luc Devroye (2011): [Simulating size-constrained Galton-Watson trees.](http://luc.devroye.org/gw-simulation.pdf)
- Svante Janson (2012): [Simply generated trees, conditioned Galton-Watson trees, random allocations and condensation.](https://projecteuclid.org/euclid.ps/1331216239)