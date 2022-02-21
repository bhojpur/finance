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
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func NewTestDiscrete() (*Discrete, error) {
	sys, _ := NewSystem(
		mat.NewDense(2, 2, []float64{0, 1, 0, 0}), // A
		mat.NewDense(2, 1, []float64{0, 1}),       // B
		mat.NewDense(1, 2, []float64{1, 0}),       // C
		mat.NewDense(1, 1, []float64{0}),          // D
	)
	dt := 0.1
	//return NewDiscrete(sys.A, sys.B, nil, nil, dt)
	return sys.Discretize(dt)
}

func TestPredict(t *testing.T) {
	sys, err := NewTestDiscrete()
	if err != nil {
		t.Error("Internal error in creating test system")
	}

	state := mat.NewVecDense(2, []float64{0, 1}) // x = position, velocity
	input := mat.NewVecDense(1, []float64{2})    // u = accelartion
	newState := sys.Predict(state, input)
	if err != nil {
		fmt.Println(err)
		t.Error("Predict returned error")
	}

	//
	expectedState := mat.NewVecDense(2, []float64{0.11, 1.2})
	if !mat.EqualApprox(newState, expectedState, 1e-4) {
		fmt.Println("Returned:", newState)
		fmt.Println("Expected:", expectedState)
		t.Error("Predict returned wrong state")
	}
}
