# Go package for Lorentz vector

**For manipulating Lorentz vectors, I recommand to use the maintained package `go-hep.org/x/hep/fmom` ([documentation](https://godoc.org/go-hep.org/x/hep/fmom))

The full documentation is available on [godoc.org](https://godoc.org/github.com/rmadar/go-lorentz-vector/lv).

### A word of caution

This repository is primarily to practice. The stucture is probably not optimal for a proper package structure, but this might evolve. For now, it is just to better undertand how to create/manipulate the go langage. The maths are mostly taken from [TLorentzVector ROOT class](https://github.com/root-project/root/blob/master/math/physics/src/TLorentzVector.cxx).

### Content 

The lorentz vector object is a four-component object used in high energy physics, and is implemented as a 3D vector and a fourth component `{r3.Vector, P4}` object. The package relies then naturally on `github.com/golang/geo/r3` package (3D vectors).

The repository is structured as follow:
  + `lv/lorentzvector.go`: code of the `lv` package
  + `show-lv/main.go`: short code showing how `lv` package can be used

