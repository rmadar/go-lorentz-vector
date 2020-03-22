package lv

import (
	"math"
	"github.com/golang/geo/r3"
)

// 4-vector type
type FourVec struct {
	Pvec r3.Vector
	P4   float64
}

// Creator of the type FourVec using (px, py, pz, e)
func NewFourVecPxPyPzE(px, py, pz, e float64) (v FourVec) {
	v.Pvec = r3.Vector{px, py, pz}
	v.P4 = e
	return v
}

// Creator of the type FourVec using (px, py, pz, m)
func NewFourVecPxPyPzM(px, py, pz, m float64) (v FourVec) {
	v.Pvec = r3.Vector{px, py, pz}
	v.P4 = math.Sqrt(v.P2() + math.Pow(m, 2))
	return v
}

// Creator of type FourVec using (pT, Eta, Phi and M)
func NewFourVecPtEtaPhiM(pt, eta, phi, m float64) (v FourVec) {
	v.Pvec = r3.Vector{pt*math.Cos(phi), pt*math.Sin(phi), pt*math.Sinh(eta)}
	v.P4  = math.Sqrt(v.P2() + math.Pow(m, 2))
	return v
}

// Creator of type FourVec using (pT, Eta, Phi and E)
func NewFourVecPtEtaPhiE(pt, eta, phi, e float64) (v FourVec) {
	v.Pvec = r3.Vector{pt*math.Cos(phi), pt*math.Sin(phi), pt*math.Sinh(eta)}
	v.P4 = e
	return v
}

// Get Px
func (v *FourVec) Px() (float64){
	return v.Pvec.X
}

// Get Py
func (v *FourVec) Py() (float64){
	return v.Pvec.Y
}

// Get Pz
func (v *FourVec) Pz() (float64){
	return v.Pvec.Z
}

// Get E
func (v *FourVec) E() (float64){
	return v.P4
}

// Transverse momentum
func (v *FourVec) Pt() (float64) {
	px, py := v.Pvec.X, v.Pvec.Y
	return math.Sqrt(px*px + py*py)
}

// Get Eta
func (v *FourVec) Eta() (float64){
	p, pz := v.P(), v.Pz()
	return 0.5*math.Log( (p+pz)/(p-pz) );
}

// Get Phi
// FIX-ME: need to check if the s1.Angle is [0, 2pi] as HEP convention
func (v *FourVec) Phi() (float64){
	pt := r3.Vector{v.Px(), v.Py(), 0}
        Ox := r3.Vector{1, 0, 0}
        return Ox.Angle(pt).Radians()
}

// Get rapidity
func (v *FourVec) Rapidity() (float64) {
	e, pz := v.E(), v.Pz()
	return 0.5*math.Log( (e+pz)/(e-pz) );
}

// Squared distance of the 3-vector
func (v *FourVec) P2() (float64) {
	return v.Pvec.Norm2() 
}

// Distance of the 3-vector
func (v *FourVec) P() (float64) {
	return v.Pvec.Norm() 
}

// Transverse energy
func (v *FourVec) Et() (float64) {
	e2  := v.E() * v.E()
	pt2 := v.Pt() * v.Pt()
	p2  := v.P2()
	return math.Sqrt( e2 * pt2/p2 )
}

// Lorentz scalar product
func (v *FourVec) Dot(u FourVec) (float64) {
	pv, pu := v.Pvec, u.Pvec
	return u.P4 * v.P4 - pv.Dot(pu)
}	

// Invariant mass ('lorentz norm' of the 4-vector) 
func (v *FourVec) M() (float64) {
	return math.Sqrt(v.Dot(*v))
}

// Get DeltaR
func (v *FourVec) DeltaR(u FourVec) (float64) {
	dphi := v.Phi() - u.Phi()
	deta := v.Eta() - u.Eta()
	return math.Sqrt(dphi*dphi + deta*deta)
}

// Get DeltaPhi
func (v *FourVec) DeltaPhi(u FourVec) (float64) {
	return math.Acos(math.Cos(v.Phi() - u.Phi()))
}

// Get 3D boost
func (v *FourVec) GetBoost() (r3.Vector){
	return v.Pvec.Mul(1./v.P4)
}

// Apply Lorentz boost
// (FIX-ME: improve notation since gamma2 != gamma*gamma)
func (v *FourVec) ApplyBoost(b r3.Vector) (vb FourVec){

	// Transformation parameters
	v_p := v.Pvec 
	b2     := b.Norm2()
	bp     := b.Dot(v_p)
	gamma  := 1.0 / math.Sqrt(1.0 - b2)
	gamma2 := (gamma - 1.0)/b2 

	// Boost the 4-vector
	vb.Pvec = v_p.Add( b.Mul(gamma2*bp + gamma*v.P4) )
	vb.P4  = gamma*(v.P4 + bp)
	return vb
}

// Four-vector addition
func (v *FourVec) Add(vec FourVec) (vsum FourVec) {
	vsum.Pvec = v.Pvec.Add(vec.Pvec)
	vsum.P4 = v.P4 + vec.P4
	return vsum
}

// Four-vector addition
func (v *FourVec) Subtract(vec FourVec) (vdiff FourVec) {
	vdiff.Pvec = v.Pvec.Sub(vec.Pvec)
	vdiff.P4 = v.P4 - vec.P4
	return vdiff
}

// Four-vector addition
func (v *FourVec) Multiply(a float64) (vprod FourVec) {
	vprod.Pvec = v.Pvec.Mul(a)
	vprod.P4 = a * v.P4
	return vprod
}

