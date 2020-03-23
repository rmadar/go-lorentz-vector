package lv_test

import (
	"fmt"

	"github.com/rmadar/go-lorentz-vector/lv"
)

func ExampleFourVec_pxPyPzE() {
	p1 := lv.NewFourVecPxPyPzE(1, 2, 3, 4)
	fmt.Printf("p1=%v\n", p1)

	// Output:
	// p1=FourVec{Px: 1, Py: 2, Pz: 3, E:4, M:1.4142135623730951}
}
