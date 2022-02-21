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

// NelsonSiegelSvensson represents a spot-rate term structure
type NelsonSiegelSvensson struct {
	B0     float64 `json:"b0"`
	B1     float64 `json:"b1"`
	B2     float64 `json:"b2"`
	B3     float64 `json:"b3"`
	T1     float64 `json:"t1"`
	T2     float64 `json:"t2"`
	Spread float64 `json:"spread"`
}

// SetSpread sets spread for the calculation of the discount factors
func (nss *NelsonSiegelSvensson) SetSpread(s float64) Structure {
	nss.Spread = s
	return nss
}

// R returns the continuous compounded spot rate in percent for a maturity of m years
func (nss *NelsonSiegelSvensson) Rate(m float64) float64 {
	if m == 0.0 {
		m = 1e-7
	}
	cc := nss.B0
	cc += nss.B1 * ((1.0 - math.Exp(-m/nss.T1)) * nss.T1 / m)
	cc += nss.B2 * (((1.0 - math.Exp(-m/nss.T1)) * nss.T1 / m) - math.Exp(-m/nss.T1))
	cc += nss.B3 * (((1.0 - math.Exp(-m/nss.T2)) * nss.T2 / m) - math.Exp(-m/nss.T2))
	// cc += nss.B1 * ((1.0 - math.Exp(-m/nss.T1)) / (m / nss.T1))
	// cc += nss.B2 * (((1.0 - math.Exp(-m/nss.T1)) / (m / nss.T1)) - math.Exp(-m/nss.T1))
	// cc += nss.B3 * (((1.0 - math.Exp(-m/nss.T2)) / (m / nss.T2)) - math.Exp(-m/nss.T2))
	return cc + nss.Spread*0.01
}

// z return the discount factor for a maturity of m years (spread not considered)
func (nss *NelsonSiegelSvensson) Z(m float64) float64 {
	return math.Exp(-(nss.Rate(m) + nss.Spread*0.01) * 0.01 * m)
}

// Rate returns the annually compounded spot rate in percents for a maturity of m years (spread not considered)
// func (nss *NelsonSiegelSvensson) Rate(m float64) float64 {
// 	return (math.Exp(nss.r(m)*0.01) - 1.0) * 100.0
// }

// Z return the discount factor to discount the cash flows for a maturity of m years.
// coumpounding is the compounding frequency (if 0, set to 1 by default)
// The discount factor includes the zero-volatility spread.
// func (nss *NelsonSiegelSvensson) Z(m float64, compounding int) float64 {
// 	n := 1.0
// 	if compounding > 0 {
// 		n = float64(compounding)
// 	}
// 	return math.Pow(1.0+(nss.Rate(m)*0.01+nss.Spread*0.0001)/n, -m*n)
// }

// F is the forward discount factor at time T1=m years
// func (n *NelsonSiegelSvensson) F(m, t float64) float64 {
// 	return z(m+t) / z(m)
// }
