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

// Straight represents a straight-bond fixed income security
type Straight struct {
	maturity.Schedule
	Coupon     float64
	Redemption float64
}

// Accrued calculated the accrued interest
func (b *Straight) Accrued() float64 {
	return b.Coupon * b.DayCountFraction()
}

// PresentValue returns the "dirty" bond prices
// (for the "clean" price just subtract the accrued interest)
func (b *Straight) PresentValue(ts term.Structure) float64 {
	dcf := 0.0

	// discount coupon payments
	effCoupon := b.EffectiveCoupon(b.Coupon)
	for _, m := range b.M() {
		dcf += effCoupon * ts.Z(m)
	}

	// discount redemption value
	dcf += b.Redemption * ts.Z(b.Last())

	return dcf
}

// Duration calculates the duration of the bond
// dP/P = -D * dr
func (b *Straight) Duration(ts term.Structure) float64 {
	duration := 0.0

	p := b.PresentValue(ts)
	if p == 0.0 {
		return 0.0
	}

	// discount coupon payments
	effCoupon := b.EffectiveCoupon(b.Coupon)
	for _, m := range b.M() {
		duration += m * effCoupon * ts.Z(m)
	}

	// discount redemption value
	duration += b.Last() * b.Redemption * ts.Z(b.Last())

	return -duration / p
}

// Convexity calculates the modified duration of the bond
// dP/P = -D * dr + 1/2 * C * dr^2
func (b *Straight) Convexity(ts term.Structure) float64 {
	convex := 0.0

	p := b.PresentValue(ts)
	if p == 0.0 {
		return 0.0
	}

	// discount coupon payments
	effCoupon := b.EffectiveCoupon(b.Coupon)
	for _, m := range b.M() {
		convex += m * m * effCoupon * ts.Z(m)
	}

	// discount redemption value
	convex += b.Last() * b.Last() * b.Redemption * ts.Z(b.Last())

	return convex / p
}
