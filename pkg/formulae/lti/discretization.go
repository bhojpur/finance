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

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

//IdealDiscretization returns a discretized matrix Md = exp(A*t) * M.
//If M is nil, then it just returns exp(A*t).
func IdealDiscretization(A *mat.Dense, dt float64, M *mat.Dense) (*mat.Dense, error) {
	// A_d = exp(A*dt)
	d, err := discretize(A, dt)
	if err != nil {
		return nil, errors.New("discretization of A failed")
	}

	// M_d = exp(A*dt) * M
	var md mat.Dense
	if M != nil {
		md.Mul(d, M)
	} else {
		md.CloneFrom(d)
	}
	return &md, nil
}

//RealDiscretization returns a discretized matrix Md = Int_0^T exp(A*t) * M dt.
func RealDiscretization(A *mat.Dense, dt float64, M *mat.Dense) (*mat.Dense, error) {
	// M_d = Int_0^T exp(A*dt) * M dt
	md, err := integrate(A, M, dt)
	if err != nil {
		return nil, errors.New("discretization of M failed")
	}

	return md, nil
}
