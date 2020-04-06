# go3dplot

go3dplot is a Go package to drawing simple 3D plots.

## Installation

1. Install [gnuplot](http://www.gnuplot.info/)
2. go get github.com/borgishmorg/go3dplot

## Example

```golang
package main

import (
	"github.com/borgishmorg/go3dplot"
)

func main() {
	x := []float64{0, 1, 2, 3, 4, 5}
	y := []float64{0, 1, 2, 6, 4, 5, 6, 7, 8, 9, 10}

	u := make([][]float64, len(x))
	for i := range u {
		u[i] = make([]float64, len(y))
		for j := range u[i] {
			u[i][j] = x[i] * y[j]
		}
	}

	drawer := go3dplot.GetGnuplotDrawer()
	drawer.Draw(x, y, u, "example")
}
```