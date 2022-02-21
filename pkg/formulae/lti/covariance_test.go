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

func TestCovariancePredict(t *testing.T) {

	md := mat.NewDense(3, 3, []float64{
		0, 1, 0,
		0, 0, 1,
		1, 0, 0,
	})
	systemNoise := NewCovariance(md)

	p := mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})
	pNext := systemNoise.Predict(p, nil)

	expected1 := mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})
	if !mat.EqualApprox(pNext, expected1, 1e-8) {
		fmt.Println("received:", pNext)
		fmt.Println("expected:", expected1)
		t.Error("predict failed without noise")
	}

	// test with extra noise
	noise := mat.NewDense(3, 3, []float64{
		0.1, 0.1, 0.3,
		0.1, 0.2, 0.1,
		0.3, 0.1, 0.1,
	})
	pNext = systemNoise.Predict(p, noise)

	expected2 := mat.NewDense(3, 3, []float64{
		1.1, 0.1, 0.3,
		0.1, 1.2, 0.1,
		0.3, 0.1, 1.1,
	})
	if !mat.EqualApprox(pNext, expected2, 1e-8) {
		fmt.Println("received:", pNext)
		fmt.Println("expected:", expected2)
		t.Error("predict failed with noise")
	}

}
