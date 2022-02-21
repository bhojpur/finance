package swap

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
	"fmt"

	"github.com/bhojpur/finance/pkg/securities/term"
)

// FxRate returns the swap rate which equals the current exchange rate multiplied
// by the ratio of the relative borrowing costs in the two currencies
func FxRate(currentFx float64, cashflows, maturities []float64, tsLong, tsShort term.Structure) (float64, error) {
	var short, long float64
	for i, t := range maturities {
		long += cashflows[i] * tsLong.Z(t)
		short += cashflows[i] * tsShort.Z(t)
	}
	return currentFx * short / long, nil
}

// InterestRate returns the swap rate. The swap rate c is given by the number that makes
// the value of the swap V(0;c,T) equal to zero at initiation.
func InterestRate(maturities []float64, compounding int, ts term.Structure) (float64, error) {
	var sum float64
	for _, t := range maturities {
		sum += ts.Z(t)
	}
	if sum == 0.0 {
		return 0.0, fmt.Errorf("sum of zero-coupon bonds across maturities is zero")
	}
	swaprate := float64(compounding) * ((1.0 - ts.Z(maturities[len(maturities)-1])) / sum) * 100.0
	return swaprate, nil
}
