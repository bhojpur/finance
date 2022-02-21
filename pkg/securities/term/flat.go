package term

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

import "math"

// Flat represents a flat term structure, i.e. constant rate across maturities
type Flat struct {
	R      float64 `json:"r"`
	Spread float64 `json:"spread"`
}

// SetSpread sets the spread in bps
func (f *Flat) SetSpread(spread float64) Structure {
	f.Spread = spread
	return f
}

// Rate returns the continuously compounded spot rate in percent
func (f *Flat) Rate(t float64) float64 {
	return f.R + f.Spread*0.01
}

// Z returns the discount factor for the given maturity t
func (f *Flat) Z(t float64) float64 {
	return math.Exp(-(f.Rate(t) * 0.01) * t)
}