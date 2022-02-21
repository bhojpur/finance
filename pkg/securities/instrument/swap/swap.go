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
	"github.com/bhojpur/finance/pkg/securities/instrument/bond"
	"github.com/bhojpur/finance/pkg/securities/term"
)

// InterestRateSwap implements a plain vanilla fixed-for-floating interest
// rate swap contract. The interest rate swap is an agreement between
// two counterpaties in which one counterparty agrees to make n fixed payments
// per year at an (annualized) fixed rate c up to a maturity date T,
// while at the same time the other counterparty commits to make payments
// linked to a floating rate index.
type InterestRateSwap struct {
	// Floating rate bond (long position)
	Floating bond.Floating
	// Fixed rate bond (short position) with swap rate as coupon
	Fixed bond.Straight
}

// PresentValue returns the value of the forward contract
func (s *InterestRateSwap) PresentValue(ts term.Structure) float64 {
	return s.Floating.PresentValue(ts) - s.Fixed.PresentValue(ts)
}
