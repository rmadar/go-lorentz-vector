# Go package for Lorentz vector

This repository is mostly to practice. The stucture is probably not correct for a proper go package, but this might evolve. For now, it was just
to better undertand how to create/manipulate types. The maths are mostly taken from [TLorentzVector ROOT class](https://github.com/root-project/root/blob/master/math/physics/src/TLorentzVector.cxx). The repository contains:
  + the code `lorentzvector.go` of the package `lv` (which should be in `$GOPATH/src/<something>`)
  + an example printing some test `main.go`

**To-do**
 + [ ] cross-check the definition of `gamma2` since it is not `gamma*gamma`, change the variable name
 + [ ] cross-check the `v.Phi()` value if this is in [0, 2pi] as in HEP convention
 + [ ] compare values, boosted vectors, etc with other existing codes