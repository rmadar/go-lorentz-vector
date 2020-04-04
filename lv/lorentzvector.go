package lv

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/spatial/r3"
)

// 4-vector type
type FourVec struct {
	Pvec r3.Vec
	P4   float64
}

// Numerical tolerence of 1 keV (e.g. |p| is allowed to exceed E by 1 keV)
const precision float64 = 1e-6

// Errors message
var (
	err_PgtE string   = "lv::Lorentz vector not physical: |p|>E"
	err_boost string  = "lv::Boost not physical: |beta|>=1"
	err_invM string   = "lv::Squated invariant mass negative (below the 1e-5 tolerance)"
	err_pTnull string = "lv::NewPtEtaPhi[M,E] pT is zero, Eta NaN (incoming parton?). Please, use NewPxPyPz[E,M]()"
)

// Creator of the type FourVec using (px, py, pz, e)
func NewFourVecPxPyPzE(px, py, pz, e float64) FourVec {
	v := FourVec{
		Pvec: r3.Vec{px, py, pz},
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
		Pvec: r3.Vec{px, py, pz},
		P4:   math.Sqrt(px*px + py*py + pz*pz + m*m),
	}
}

// Creator of type FourVec using (pT, Eta, Phi and E)
func NewFourVecPtEtaPhiE(pt, eta, phi, e float64) FourVec {
	v := FourVec{
		Pvec: r3.Vec{pt * math.Cos(phi), pt * math.Sin(phi), pt * math.Sinh(eta)},
		P4:   e,
	}
	if pt<=0 {
		fmt.Printf("v = %v\n", v)
		panic(err_pTnull)
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
	p := r3.Vec{pt * math.Cos(phi), pt * math.Sin(phi), pt * math.Sinh(eta)}
	if pt<=0 {
		fmt.Printf("Pvec = %v\n", p)
		panic(err_pTnull)
	}
	return FourVec{
		Pvec: p,
		P4:   math.Sqrt(r3.Norm2(p) + m*m),
	}
}

func (v FourVec) String() string {
	return fmt.Sprintf("FourVec{Px: %v, Py: %v, Pz: %v, E:%v, M:%v}",
		v.Px(), v.Py(), v.Pz(), v.E(), v.M(),
	)
}

// Checking physics validity of the Lorentz vector, ie |p|<=E (since E2 = p2 + m2)
func (v FourVec) isPhysical() bool{
	return v.P()<=v.E()+precision
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
	if p==pz {
		fmt.Printf("v = %v\n", v)
		panic(err_pTnull)
	}
	
	return 0.5 * math.Log((p+pz)/(p-pz))
}

// Get Phi, defined as the angle between the (px, py)-vector and the x-axis in [-pi, pi[ interval
func (v FourVec) Phi() float64 {
	return math.Atan2(v.Py(), v.Px())
}

// Get rapidity
func (v FourVec) Rapidity() float64 {
	e, pz := v.E(), v.Pz()
	return 0.5 * math.Log((e+pz)/(e-pz))
}

// Squared norm of the 3-vector
func (v FourVec) P2() float64 {
	return r3.Norm2(v.Pvec)
}

// Norm of the 3-vector, ie momentum
func (v FourVec) P() float64 {
	return r3.Norm(v.Pvec)
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

// Squared invariant mass ('lorentz norm' of the 4-vector)
func (v FourVec) M2() float64 {
	return v.Dot(v)
}

// Invariant mass: m=sqrt(M2) if M2>0, else m=-sqrt(-M2)
func (v FourVec) M() float64 {
	return signedSqrt(v.M2())
}

// Squared transverse mass ('lorentz norm' of the 4-vector with Pz set to zero)
func (v FourVec) MT2() float64 {
	u := NewFourVecPxPyPzE(v.Px(), v.Py(), 0.0, v.E())
	return u.Dot(u)
}

// Invariant mass: mT=sqrt(MT2) if M2>0, else mT=-sqrt(-MT2)
func (v FourVec) MT() float64 {
	return signedSqrt(v.MT2())
}

// Get DeltaR = sqrt(dPhi*2 + dEta*2)
func (v FourVec) DeltaR(u FourVec) float64 {
	dphi := v.DeltaPhi(u)
	deta := v.DeltaEta(u)
	return math.Sqrt(dphi*dphi + deta*deta)
}

// Get DeltaPhi angle between u and v in the (px, py)-plane, in [0, pi[
func (v FourVec) DeltaPhi(u FourVec) float64 {
	return math.Acos(math.Cos(v.Phi() - u.Phi()))
}

// Get DeltaEta
func (v FourVec) DeltaEta(u FourVec) float64 {
	return v.Eta() - u.Eta()
}

// Get vectorial boost beta=(px/E, py/E, pz/E)
func (v FourVec) GetBoost() r3.Vec {
	return v.Pvec.Scale(1. / v.P4)
}

// Apply vectorial Lorentz boost (|beta|<1), defined as
//  p' = p + [(gamma-1)/beta2 * (p.beta) + gamma*E] * beta
//  E' = gamma * (E+p.beta)
func (v FourVec) ApplyBoost(beta r3.Vec) FourVec {

	// First check that v<c
	if r3.Norm(beta)>=1 {
		fmt.Println("beta  =", beta)
		fmt.Println("|beta|=", r3.Norm(beta))
		panic(err_boost)
	}
	
	// Lorentz transformation parameters
	p, E := v.Pvec, v.P4
	beta2 := r3.Norm2(beta)
	beta_dot_p := beta.Dot(p)
	gamma := 1.0 / math.Sqrt(1.0-beta2)
	alpha := (gamma - 1.0) / beta2
	
	// Return the boosted 4-vector
	return FourVec{
		Pvec: p.Add(beta.Scale(alpha*beta_dot_p + gamma*E)),
		P4:   gamma * (E + beta_dot_p),
	}
}

// Get the 4-vector in the frame where u=(0, m), aka the rest frame of u
func (v FourVec) ToRestFrameOf(u FourVec) FourVec {
	return v.ApplyBoost( u.GetBoost().Scale(-1) )
}

// Four-vector addition
func (v FourVec) Add(vec FourVec) FourVec {
	return FourVec{
		Pvec: v.Pvec.Add(vec.Pvec),
		P4:   v.P4 + vec.P4,
	}
}

// Four-vector multiplication with a scalar
func (v FourVec) Scale(a float64) FourVec {
	return FourVec{
		Pvec: v.Pvec.Scale(a),
		P4:   a * v.P4,
	}
}	

// Signed root square function: sign(x)*sqrt(abs(x))
func signedSqrt(x float64) float64 {
	if x<0 {
		return -math.Sqrt(-x)
	} else {
		return math.Sqrt(x)
	}
}
