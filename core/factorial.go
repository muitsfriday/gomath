package core

import "strconv"

// Fac ...
type Fac struct {
	n       uint64
	inverse bool
}

// FacGroup a sequence of non eval factorial
type FacGroup struct {
	group []Fac
}

// NewFac create new factorial value given unsign integer n
func NewFac(n uint64) *Fac {
	f := new(Fac)
	f.n = n
	f.inverse = false
	return f
}

// convert Fac to string
func (f *Fac) String() string {
	return strconv.FormatUint(f.n, 10) + "!"
}
