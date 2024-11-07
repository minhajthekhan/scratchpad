package main

import (
	"fmt"

	"github.com/minhajthekhan/patterns/specifications/example/pkg/legos"
)

func main() {

	spec := legos.NewLegoSpecification(
		[]string{"white", "red"},
		[]legos.LegoDimension{
			legos.NewLegoDimension("6x4", 1),
			legos.NewLegoDimension("5x5", 2),
		},
		1000,
	)

	q, args := spec.AsSQL()
	fmt.Println(q, args)
}
