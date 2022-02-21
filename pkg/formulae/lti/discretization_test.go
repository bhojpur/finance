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

func TestIdealDiscretization(t *testing.T) {
	dt := 0.1
	var config = []struct {
		A    *mat.Dense
		T    float64
		M    *mat.Dense
		Want *mat.Dense
	}{
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0,
			}),
			T: dt,
			M: nil,
			Want: mat.NewDense(2, 2, []float64{
				1, dt,
				0, 1,
			}),
		},
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0,
			}),
			T: dt,
			M: mat.NewDense(2, 2, []float64{
				0, 1,
				1, 0,
			}),
			Want: mat.NewDense(2, 2, []float64{
				dt, 1,
				1, 0,
			}),
		},
	}

	for _, cfg := range config {
		got, err := IdealDiscretization(cfg.A, cfg.T, cfg.M)
		if err == nil {
			if !mat.EqualApprox(got, cfg.Want, 1e-8) {
				fmt.Println("received=", got)
				fmt.Println("expected=", cfg.Want)
				t.Error("ideal discretization returned wrong result")
			}
		} else {
			fmt.Println(err)
			t.Error("error received in ideal discretization")
		}

	}

}

func TestRealDiscretization(t *testing.T) {
	dt := 0.1
	var config = []struct {
		A    *mat.Dense
		T    float64
		M    *mat.Dense
		Want *mat.Dense
	}{
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0,
			}),
			T: dt,
			M: mat.NewDense(2, 1, []float64{
				0,
				1,
			}),
			Want: mat.NewDense(2, 1, []float64{
				0.5 * dt * dt,
				dt,
			}),
		},
	}

	for _, cfg := range config {
		got, err := RealDiscretization(cfg.A, cfg.T, cfg.M)
		if err == nil {
			if !mat.EqualApprox(got, cfg.Want, 1e-8) {
				fmt.Println("received=", got)
				fmt.Println("expected=", cfg.Want)
				t.Error("ideal discretization returned wrong result")
			}
		} else {
			fmt.Println(err)
			t.Error("error received in ideal discretization")
		}

	}

}
