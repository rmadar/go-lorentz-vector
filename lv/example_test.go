package lv_test

import (
	"fmt"

	"github.com/rmadar/go-lorentz-vector/lv"
)

func ExampleNewFourVecPxPyPzE() {
	p := lv.NewFourVecPxPyPzE(1, 2, 3, 4)
	fmt.Printf("p = %v\n", p)

	// Output:
	// p = FourVec{Px: 1, Py: 2, Pz: 3, E:4, M:1.4142135623730951}
}

func ExampleNewFourVecPtEtaPhiM() {
	p := lv.NewFourVecPtEtaPhiM(1, 2, 3, 4)
	fmt.Printf("p = %v\n", p)

	// Output:
	// p = FourVec{Px: -0.9899924966004454, Py: 0.1411200080598672, Pz: 3.626860407847019, E:5.491276392425375, M:3.999999999999999}
}
