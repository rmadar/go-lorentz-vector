package lv

import (
	"math"
)

// 4-vector type
type FourVec struct {
	Px float64
	Py float64
	Pz float64
	E  float64
}

// Creator of the type FourVec using (px, py, pz, E)
func NewFourVecPxPyPzE(px, py, pz, e float64) (v FourVec) {
	v.Px = px
	v.Py = py
	v.Pz = pz
	v.E  = e
	return v
}

// Creator of the type FourVec using (px, py, pz, M)
func NewFourVecPxPyPzM(px, py, pz, m float64) (v FourVec) {
	v.Px = px
	v.Py = py
	v.Pz = pz
	v.E  = math.Sqrt(v.P2() + math.Pow(m, 2))
	return v
}

// Creator of type FourVec using (pT, Eta, Phi and M)
func NewFourVecPtEtaPhiM(pt, eta, phi, m float64) (v FourVec) {
	v.Px = pt  * math.Cosh(eta)
	v.Py = eta * math.Sinh(eta)
	v.Pz = phi * pt * math.Cos(phi) * v.Px
	v.E  = math.Sqrt(v.P2() + math.Pow(m, 2))
	return v
}

// Creator of type FourVec using (pT, Eta, Phi and E)
func NewFourVecPtEtaPhiE(pt, eta, phi, e float64) (v FourVec) {
	v.Px = pt  * math.Cosh(eta)
	v.Py = eta * math.Sinh(eta)
	v.Pz = phi * pt * math.Cos(phi) * v.Px
	v.E  = e
	return v
}

// Get Eta
func (v *FourVec) Eta() (float64){
	return 1.
}

// Get Phi
func (v *FourVec) Phi() (float64){
	return 1.
}

// Squared distance of the 3-vector
func (v *FourVec) P2() (float64) {
	return math.Pow(v.Px, 2) + math.Pow(v.Py, 2) + math.Pow(v.Pz, 2)
}

// Distance of the 3-vector
func (v *FourVec) P() (float64) {
	return math.Sqrt(v.P2())
}

// Transverse momentum
func (v *FourVec) Pt() (float64) {
	return math.Sqrt(v.P2() - math.Pow(v.Pz, 2))
}

// Transverse momentum
func (v *FourVec) Et() (float64) {
	return 1.0
}

// Invariant mass
func (v *FourVec) M() (float64) {
	return math.Sqrt(math.Pow(v.E, 2) - v.P2())
}

// Get DeltaR
func (v *FourVec) DeltaR(u FourVec) (dr float64) {
	return 1.0
}

// Get DeltaPhi
func (v *FourVec) DeltaPhi(u FourVec) (dphi float64) {
	return 1.0
}

// Get 3D boost
func (v *FourVec) GetBoost() (bx, by, bz float64){
	bx = v.Px/v.E
	by = v.Py/v.E
	bz = v.Pz/v.E
	return bx, by, bz
}

// Apply Lorentz boost
func (v *FourVec) ApplyBoost(bx, by, bz float64) (vb FourVec){

	// Transformation parameters
	b2     := bx*bx + by*by + bz*bz
	gamma  := 1.0 / math.Sqrt(1.0 - b2)
	bp     := bx*v.Px + by*v.Py + bz*v.Pz
	gamma2 := (gamma - 1.0)/b2;

	// Boost the 4-vector
	vb.Px = v.Px + gamma2*bp*bx + gamma*bx*v.E
	vb.Py = v.Py + gamma2*bp*by + gamma*by*v.E
	vb.Pz = v.Pz + gamma2*bp*bz + gamma*bz*v.E
	vb.E  = gamma*(v.E + bp)
	return vb
}

// Four vector addition
func (v *FourVec) Add(vec FourVec) (vsum FourVec) {
	vsum.Px = v.Px + vec.Px
	vsum.Py = v.Py + vec.Py
	vsum.Pz = v.Pz + vec.Pz
	vsum.E  = v.E  + vec.E
	return vsum
}

// Four vector addition
func (v *FourVec) Subtract(vec FourVec) (vdiff FourVec) {
	vdiff.Px = v.Px - vec.Px
	vdiff.Py = v.Py - vec.Py
	vdiff.Pz = v.Pz - vec.Pz
	vdiff.E  = v.E  - vec.E
	return vdiff
}

// Four vector addition
func (v *FourVec) Multiply(a float64) (vprod FourVec) {
	vprod.Px = a * v.Px
	vprod.Py = a * v.Py
	vprod.Pz = a * v.Pz
	vprod.E  = a * v.E
	return vprod
}

