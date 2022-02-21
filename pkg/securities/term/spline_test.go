package term_test

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
	"testing"

	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestSpline(t *testing.T) {

	// use NSS term structure as reference
	refTerm := term.NelsonSiegelSvensson{
		-0.266372,
		-0.471343,
		5.68789,
		-5.12324,
		5.74881,
		4.14426,
		0.0, // spread
	}

	maturities := []float64{0.25, 0.5, 1.0, 2.0, 3.0, 5.0, 7.0, 10.0, 15.0, 20.0}
	var refRate, refZ []float64
	for _, t := range maturities {
		refRate = append(refRate, refTerm.Rate(t))
		refZ = append(refZ, refTerm.Z(t))
	}

	// create spline term structure
	spread := 0.0
	spline := term.NewSpline(maturities, refZ, spread)

	// test spline approximation: rates
	sum := 0.0
	for i, t := range maturities {
		sum += math.Pow(spline.Rate(t)-refRate[i], 2.0)
	}

	if math.Abs(sum) > 0.000001 {
		t.Errorf("splines do not accurately interpolte rates of yield curve; got: %v, expected: %v", sum, 0.0)
	}

	// test spline approximation: rates
	sum = 0.0
	for i, t := range maturities {
		sum += math.Pow(spline.Z(t)-refZ[i], 2.0)
	}

	if math.Abs(sum) > 0.000001 {
		t.Errorf("splines do not accurately interpolte discount factors Z of yield curve; got: %v, expected: %v", sum, 0.0)
	}
}
