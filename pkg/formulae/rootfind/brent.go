package rootfind

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"math"
)

// Brent - Brent's Method finds the root of the given quadratic function f in [a,b].
// The precision is the number of digits after the floating point.
// reference: https://en.wikipedia.org/wiki/Brent%27s_method
func Brent(f func(x float64) float64, a, b float64, precision int) (r float64, err error) {
	var (
		delta            = EpsilonF64 * (b - a) // numerical tolerance
		acceptance       = math.Pow10(-precision)
		fa               = f(a)
		fb               = f(b)
		c                = a
		fc               = fa
		s                float64
		fs               float64
		d                float64
		wasBisectionUsed = true
		absBMinusC       float64
		absCMinusD       float64
		absSMinusB       float64
		tmp              float64
		// swap - a becomes b, b becomes a
		swap = func() {
			tmp = a
			a = b
			b = tmp
			tmp = fa
			fa = fb
			fb = tmp
		}
	)
	if a > b {
		swap()
	}
	if fa*fb > 0 {
		if a >= 0 || b <= 0 {
			return 0, ErrRootIsNotBracketed
		}
		f0 := f(0)
		if f0*fb > 0 && f0*fa > 0 {
			return 0, ErrRootIsNotBracketed
		}
	}
	if math.Abs(fa) < math.Abs(fb) {
		swap()
	}
	for fb != 0 && math.Abs(b-a) > acceptance {
		if fa != fc && fb != fc { // inverse quadratic interpolation
			s = (a*fb*fc)/((fa-fb)*(fa-fc)) + (b*fa*fc)/((fb-fa)*(fb-fc)) + (c*fa*fb)/((fc-fa)*(fc-fb))
		} else { // secant method
			s = b - fb*(b-a)/(fb-fa)
		}
		absBMinusC = math.Abs(b - c)
		absCMinusD = math.Abs(c - d)
		absSMinusB = math.Abs(s - b)
		switch {
		case s < (3*a+b)/4 || s > b,
			wasBisectionUsed && absSMinusB >= absBMinusC/2,
			!wasBisectionUsed && absSMinusB >= absCMinusD/2,
			wasBisectionUsed && absBMinusC < delta,
			!wasBisectionUsed && absCMinusD < delta: // bisection method
			s = (a + b) / 2
			wasBisectionUsed = true
			break
		default:
			wasBisectionUsed = false
			break
		}
		fs = f(s)
		d = c // d is first defined here; is not use in the first step above because wasBisectionUsed set to true
		c = b
		fc = fb
		if fa*fs < 0 {
			b = s
			fb = fs
		} else {
			a = s
			fa = fs
		}
		if math.Abs(fa) < math.Abs(fb) {
			swap()
		}
	}
	return s, nil
}
