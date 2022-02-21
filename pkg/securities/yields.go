package securities

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
	"github.com/bhojpur/finance/pkg/formulae/rootfind"
	"github.com/bhojpur/finance/pkg/securities/term"
)

var (
	Precision = 6
)

// Irr calculates the internal rate of return of a security
func Irr(investment float64, s Security) (float64, error) {
	f := func(irr float64) float64 {
		return s.PresentValue(&term.Flat{irr, 0.0}) - investment
	}

	root, err := rootfind.Brent(f, -20.0, 20.0, Precision)
	return root, err
}

// Spread calculates the implied static (zero-volatility) spread
func Spread(investment float64, s Security, ts term.Structure) (float64, error) {
	f := func(spread float64) float64 {
		value := s.PresentValue(ts.SetSpread(spread))
		return value - investment
	}

	root, err := rootfind.Brent(f, -10.0, 10000.0, Precision)
	return root, err
}

// ImpliedVola calculates the implied volatility for a given option price
func ImpliedVola(price float64, o Option, ts term.Structure) (float64, error) {
	f := func(vola float64) float64 {
		o.SetVola(vola)
		value := o.PresentValue(ts)
		return value - price
	}

	root, err := rootfind.Brent(f, 0.0, 1000.0, Precision)
	return root, err
}
