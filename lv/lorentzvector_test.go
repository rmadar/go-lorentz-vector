package lv_test

import (
	"testing"

	"github.com/golang/geo/r3"
	"github.com/rmadar/go-lorentz-vector/lv"
)

func TestBoost(t *testing.T) {
	// Create an object using (px, py, pz, E)
	vec1 := lv.NewFourVecPxPyPzE(1., 2., 3., 4.)
	vec2 := lv.NewFourVecPtEtaPhiM(1., 2., 3, 4)

	// Combination of 4-vectors v1+v2+3*v1
	vec3 := vec1.Multiply(3)
	vec3 = vec3.Add(vec2.Add(vec1))

	// Boost vec1 and check that the boost is (0, 0, 0) in its rest frame
	boost := vec1.GetBoost()
	vec1RF := vec1.ApplyBoost(boost.Mul(-1))
	boostRF := vec1RF.GetBoost()
	var zero r3.Vector
	if boostRF != zero {
		t.Fatalf("invalid boost. got=%v, want=%v", boostRF, zero)
	}
}