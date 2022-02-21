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

// Discrete represents a discrete LTI system.
//
// The parameters are:
// 	A_d: Discretized Ssystem matrix
// 	B_d: Discretized Control matrix
//
//
type Discrete struct {
	Ad *mat.Dense
	Bd *mat.Dense
	C  *mat.Dense
	D  *mat.Dense
}

//NewDiscrete returns a Discrete struct
func NewDiscrete(A, B, C, D *mat.Dense, dt float64) (*Discrete, error) {

	// A_d = exp(A*dt)
	ad, err := discretize(A, dt)
	if err != nil {
		return nil, errors.New("discretization of A failed")
	}

	// B_d = Int_0^T exp(A*dt) * B dt
	bd, err := integrate(A, B, dt)
	if err != nil {
		return nil, errors.New("discretization of B failed")
	}

	return &Discrete{
		Ad: ad,
		Bd: bd,
		C:  C,
		D:  D,
	}, nil
}

// Predict predicts  x(k+1) = A_discretized * x(k) + B_discretized * u(k)
func (d *Discrete) Predict(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {
	// x(k+1) = A_d * x + B_d * u
	return multAndSumOp(d.Ad, x, d.Bd, u)
}

//Response returns the output vector y(t) = C * x(t) + D * u(t)
func (d *Discrete) Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {
	// y(t) = C * x(t) + D * u(t)
	return multAndSumOp(d.C, x, d.D, u)
}

// Controllable checks the controllability of the LTI system.
func (d *Discrete) Controllable() (bool, error) {
	// system is controllable if
	// rank( [B, A B, A^2 B, A^n-1 B] ) = n
	return checkControllability(d.Ad, d.Bd)
}

// Observable checks the observability of the LTI system.
func (d *Discrete) Observable() (bool, error) {
	// system is observable if
	// rank( S=[C, C A, C A^2, ..., C A^n-1]' ) = n
	return checkObservability(d.Ad, d.C)
}
