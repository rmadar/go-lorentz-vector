# Go package for Lorentz vector


### Word of caution

This repository is primarily to practice. The stucture is probably not correct for a proper go package, but this might evolve. For now, it is just
to better undertand how to create/manipulate types in go. The maths are mostly taken from [TLorentzVector ROOT class](https://github.com/root-project/root/blob/master/math/physics/src/TLorentzVector.cxx).

### Content 

The lorentz vector object is a four-component object used in high energy physics, and is implemented as a 3D vector and a fourth component `{r2.Vector, P4}` object. The package relies then naturally on `github.com/golang/geo/r3` package (3D vectors).

The repository is structured as follow:
  + `lorentzvector.go`: code of of the package `lv` (which should be in `$GOPATH/src/<something>` in order to be used in a go program)
  + `main.go`: usage example of the `lv` package

### To-do
 + [ ] cross-check the definition of `gamma2` since it is not `gamma*gamma`, change the variable name
 + [ ] cross-check the `v.Phi()` value if this is in [0, 2pi] as in HEP convention
 + [ ] compare values, boosted vectors, etc with other existing codes