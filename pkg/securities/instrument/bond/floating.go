package bond

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
	"github.com/bhojpur/finance/pkg/securities/maturity"
	"github.com/bhojpur/finance/pkg/securities/term"
)

// Floating represents a floating-rate bond
type Floating struct {
	maturity.Schedule
	// Rate is the current rate in percent for next coupon payment
	// which is known today
	Rate       float64
	Redemption float64
}

// Accrued calculated the accrued interest
func (f *Floating) Accrued() float64 {
	return f.Rate * f.Schedule.DayCountFraction()
}

// PresentValue returns the "dirty" bond prices (for the "clean" price just subtract the accrued interest)
func (f *Floating) PresentValue(ts term.Structure) float64 {
	pv := 0.0

	// discount face value at next reset date
	effRate := f.EffectiveCoupon(f.Rate)
	pv += (f.Redemption + effRate) * ts.Z(f.Next())

	return pv
}

// Duration calculates the duration of the floating-rate bond
// dP/P = -D * dr
func (f *Floating) Duration(ts term.Structure) float64 {
	p := f.PresentValue(ts)
	if p == 0.0 {
		return 0.0
	}

	// discount redemption value
	duration := f.Next() * (f.Redemption + f.EffectiveCoupon(f.Rate)) * ts.Z(f.Next())

	return -duration / p
}

// Convexity calculates the modified duration of the bond
// dP/P = -D * dr + 1/2 * C * dr^2
func (f *Floating) Convexity(ts term.Structure) float64 {
	p := f.PresentValue(ts)
	if p == 0.0 {
		return 0.0
	}

	convex := f.Next() * f.Next() * (f.Redemption + f.EffectiveCoupon(f.Rate)) * ts.Z(f.Next())

	return convex / p
}
