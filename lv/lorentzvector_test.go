package lv_test

import (
	"testing"

	"gonum.org/v1/gonum/spatial/r3"
	
	"github.com/rmadar/go-lorentz-vector/lv"
)

func TestBoost(t *testing.T) {
	// Create an object using (px, py, pz, E)
	vec1 := lv.NewFourVecPxPyPzE(1, 2, 3, 4)
	vec2 := lv.NewFourVecPtEtaPhiM(1, 2, 3, 4)

	// Combination of 4-vectors v1+v2+3*v1
	vec3 := vec1.Scale(3)
	vec3 = vec3.Add(vec2.Add(vec1))

	// Boost vec1 and check that the boost is (0, 0, 0) in its rest frame
	vec1RF := vec1.ToRestFrameOf(vec1)
	boostRF := vec1RF.GetBoost()
	var zero r3.Vec
	if boostRF != zero {
		t.Fatalf("Invalid boost. got=%v, want=%v", boostRF, zero)
	}
}
