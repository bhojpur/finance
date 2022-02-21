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

// It implements a general state-space representation of a linear system.
// A state-space representation can be expressed in a matrix form as
//
// x'(t) = A * x(t) + B * u(t)
// and
// y(t)  = C * x(t) + D * u(t)
//
// where x(t) represents the state and u(t) the control input vectors and the
// matrices are
// A: System matrix,
// B: Control matrix,
// C: Output matrix and
// D: Feedforward matrix

import "gonum.org/v1/gonum/mat"

//LTI represents a general time-continuous state-space LTI system
type LTI interface {
	Observable() (bool, error)
	Controllable() (bool, error)
	Response(x, u *mat.VecDense) *mat.VecDense
}

//Predictor represents a discretized LTI system
type Predictor interface {
	LTI
	Predict(x, u *mat.VecDense) *mat.VecDense
}
