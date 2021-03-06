package main

import (
	"fmt"

	"github.com/rmadar/go-lorentz-vector/lv"
	"go-hep.org/x/hep/fmom"
)

func main() {

	// Quick checks
	fmt.Println("\nPrinting some 4-vector manipulations")
	fmt.Println("====================================")
	quickTest()

	// Cross-check with fmom
	fmt.Println("\n\nComparison with fmom")
	fmt.Println("====================")
	checkAgainstFmom(1., 2., 3., 4.)
}

func quickTest() {

	// Create an object using (px, py, pz, E)
	vec1 := lv.NewFourVecPxPyPzE(1., 2., 3., 4.)
	fmt.Println("\nVec1:", vec1)

	// Using now (Pt, Eta, Phi, M)
	vec2 := lv.NewFourVecPtEtaPhiM(1., 2., 3, 4)
	fmt.Println("\nVec2:", vec2)

	// Combination of 4-vectors v1+v2+3*v1
	vec3 := vec1.Scale(3)
	vec_sum := vec3.Add(vec2.Add(vec1))
	fmt.Println("\nVec[sum]:", vec_sum)

	// Boost vec1 and check that the boost is (0, 0, 0) in its rest frame
	boost := vec1.GetBoost()
	vec1_RF := vec1.ApplyBoost(boost.Scale(-1))
	boost_RF := vec1_RF.GetBoost()
	fmt.Println("\nCheck the boost is (0, 0, 0) in the particle rest frame:",
		boost_RF,
	)
}

func checkAgainstFmom(px, py, pz, e float64) {
	p_fmom := fmom.NewPxPyPzE(px, py, pz, e)
	p_here := lv.NewFourVecPxPyPzE(px, py, pz, e)
	fmt.Println("Pt  : lv=", p_here.Pt(), "vs fmom=", p_fmom.Pt())
	fmt.Println("Eta : lv=", p_here.Eta(), "vs fmom=", p_fmom.Eta())
	fmt.Println("Phi : lv=", p_here.Phi(), "vs fmom=", p_fmom.Phi())
	fmt.Println("Mass: lv=", p_here.M(), "vs fmom=", p_fmom.M())
}
