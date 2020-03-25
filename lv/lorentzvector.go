package lv

import (
	"fmt"
	"math"
	
	"github.com/golang/geo/r3"
)

// 4-vector type
type FourVec struct {
	Pvec r3.Vector
	P4   float64
}

// Errors message
var err_PgtE string = "lv::Lorentz vector not physical: |p|>E"
var err_boost string = "lv:: Boost not physical: |beta|>=1"


// Creator of the type FourVec using (px, py, pz, e)
func NewFourVecPxPyPzE(px, py, pz, e float64) FourVec {
	v := FourVec{
		Pvec: r3.Vector{px, py, pz},
		P4:   e,
	}
	if v.isPhysical() {
		return v
	} else {
		fmt.Printf("v = %v\n", v)
		panic(err_PgtE)
	}
}

// Creator of the type FourVec using (px, py, pz, m)
func NewFourVecPxPyPzM(px, py, pz, m float64) FourVec {
	return FourVec{
		Pvec: r3.Vector{px, py, pz},
		P4:   math.Sqrt(px*px + py*py + pz*pz + m*m),
	}
}

// Creator of type FourVec using (pT, Eta, Phi and E)
func NewFourVecPtEtaPhiE(pt, eta, phi, e float64) FourVec {
	v := FourVec{
		Pvec: r3.Vector{pt * math.Cos(phi), pt * math.Sin(phi), pt * math.Sinh(eta)},
		P4:   e,
	}
	if v.isPhysical() {
		return v
	} else {
		fmt.Printf("v = %v\n", v)
		panic(err_PgtE)
	}
}

// Creator of type FourVec using (pT, Eta, Phi and M)
func NewFourVecPtEtaPhiM(pt, eta, phi, m float64) FourVec {
	p := r3.Vector{pt * math.Cos(phi), pt * math.Sin(phi), pt * math.Sinh(eta)}
	return FourVec{
		Pvec: p,
		P4:   math.Sqrt(p.Norm2() + m*m),
	}
}

func (v FourVec) String() string {
	return fmt.Sprintf("FourVec{Px: %v, Py: %v, Pz: %v, E:%v, M:%v}",
		v.Px(), v.Py(), v.Pz(), v.E(), v.M(),
	)
}

// Checking physics validity of the Lorentz vector, ie |p|<=E (since E2 = p2 + m2)
func (v FourVec) isPhysical() bool{
	return v.P()<=v.E()
}

// Get Px
func (v FourVec) Px() float64 {
	return v.Pvec.X
}

// Get Py
func (v FourVec) Py() float64 {
	return v.Pvec.Y
}

// Get Pz
func (v FourVec) Pz() float64 {
	return v.Pvec.Z
}

// Get Energy
func (v FourVec) E() float64 {
	return v.P4
}

// Transverse momentum
func (v FourVec) Pt() float64 {
	px, py := v.Pvec.X, v.Pvec.Y
	return math.Sqrt(px*px + py*py)
}

// Get pseudo-rapidity Eta
func (v FourVec) Eta() float64 {
	p, pz := v.P(), v.Pz()
	return 0.5 * math.Log((p+pz)/(p-pz))
}

// Get Phi, defined as the angle between the (px, py)-vector and the x-axis
// FIX-ME: need to check if the s1.Angle is [0, 2pi] as HEP convention
func (v FourVec) Phi() float64 {
	pt := r3.Vector{v.Px(), v.Py(), 0}
	Ox := r3.Vector{1, 0, 0}
	return Ox.Angle(pt).Radians()
}

// Get rapidity
func (v FourVec) Rapidity() float64 {
	e, pz := v.E(), v.Pz()
	return 0.5 * math.Log((e+pz)/(e-pz))
}

// Squared norm of the 3-vector
func (v FourVec) P2() float64 {
	return v.Pvec.Norm2()
}

// Norm of the 3-vector, ie momentum
func (v FourVec) P() float64 {
	return v.Pvec.Norm()
}

// Transverse energy defined as ET = E*pT/p
func (v FourVec) Et() float64 {
	e2 := v.E() * v.E()
	pt2 := v.Pt() * v.Pt()
	p2 := v.P2()
	return math.Sqrt(e2 * pt2 / p2)
}

// Lorentz scalar product defined as v1.v2 = p1.dot(p2) - E1*E3
func (v FourVec) Dot(u FourVec) float64 {
	pv, pu := v.Pvec, u.Pvec
	return u.P4*v.P4 - pv.Dot(pu)
}

// Invariant mass ('lorentz norm' of the 4-vector)
func (v FourVec) M() float64 {
	return math.Sqrt(v.Dot(v))
}

// Get DeltaR = sqrt(dPhi*2 + dEta*2)
func (v FourVec) DeltaR(u FourVec) float64 {
	dphi := v.Phi() - u.Phi()
	deta := v.Eta() - u.Eta()
	return math.Sqrt(dphi*dphi + deta*deta)
}

// Get DeltaPhi
func (v FourVec) DeltaPhi(u FourVec) float64 {
	return math.Acos(math.Cos(v.Phi() - u.Phi()))
}

// Get vectorial boost beta=(px/E, py/E, pz/E)
func (v FourVec) GetBoost() r3.Vector {
	return v.Pvec.Mul(1. / v.P4)
}

// Apply vectorial Lorentz boost (|beta|<1), defined as
//  p' = p + [(gamma-1)/beta2 * (p.beta) + gamma*E] * beta
//  E' = gamma * (E+p.beta)
func (v FourVec) ApplyBoost(beta r3.Vector) FourVec {

	// First check that v<c
	if beta.Norm()>=1 {
		fmt.Println("beta=", beta)
		panic(err_boost)
	}
	
	// Lorentz transformation parameters
	p, E := v.Pvec, v.P4
	beta2 := beta.Norm2()
	beta_dot_p := beta.Dot(p)
	gamma := 1.0 / math.Sqrt(1.0-beta2)
	alpha := (gamma - 1.0) / beta2
	
	// Return the boosted 4-vector
	return FourVec{
		Pvec: p.Add(beta.Mul(alpha*beta_dot_p + gamma*E)),
		P4:   gamma * (E + beta_dot_p),
	}
}

// Get the 4-vector in the frame where u=(0, m), aka the rest frame of u
func (v FourVec) ToRestFrameOf(u FourVec) FourVec {
	return v.ApplyBoost( u.GetBoost().Mul(-1) )
}

// Four-vector addition
func (v FourVec) Add(vec FourVec) FourVec {
	return FourVec{
		Pvec: v.Pvec.Add(vec.Pvec),
		P4:   v.P4 + vec.P4,
	}
}

// Four-vector multiplication with a scalar
func (v FourVec) Multiply(a float64) FourVec {
	return FourVec{
		Pvec: v.Pvec.Mul(a),
		P4:   a * v.P4,
	}
}	
