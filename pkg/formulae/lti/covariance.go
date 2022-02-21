package lti

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

import "gonum.org/v1/gonum/mat"

//Covariance contains a discretized matrix to predict the next covariance matrix
// according to p(k+1) = Md * p(k) * Md^T
type Covariance struct {
	Md *mat.Dense
}

//NewCovariance creates a new covariance struct which
//needs to be initialized with a discretizied matrix Md
func NewCovariance(md *mat.Dense) *Covariance {
	return &Covariance{Md: md}
}

//Predict propagates the covariance p(k) to p(k+1)
//according to p(k+1) = Md * p(k) * Md^T;
//additionally noise is added if not nil
func (c *Covariance) Predict(p *mat.Dense, noise *mat.Dense) *mat.Dense {
	// p(k+1) = m * p(k) * m^T + noise
	var pmt, mpmt mat.Dense
	pmt.Mul(p, c.Md.T())
	mpmt.Mul(c.Md, &pmt)

	if noise != nil {
		mpmt.Add(&mpmt, noise)
	}

	return &mpmt
}
