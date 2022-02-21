package rootfind

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
	"math"
	"testing"
)

func TestBrent(t *testing.T) {
	cases := []struct {
		f             func(float64) float64
		intervalStart float64
		intervalEnd   float64
		precision     int
		roots         []float64
		expectedErr   error
	}{
		{
			func(x float64) float64 {
				return (x + 3) * math.Pow(x-1, 2)
			},
			-100000, 100000, 4,
			[]float64{-3, 1},
			nil,
		},
		{
			func(x float64) float64 {
				return (x + 3) * math.Pow(x-2, 2)
			},
			-2, 1.5, 5,
			[]float64{-3, 2},
			ErrRootIsNotBracketed,
		},
		{
			func(x float64) float64 {
				return math.Pow(x, 4) - 2*math.Pow(x, 2) + 0.25
			},
			0, 1, 6,
			[]float64{0.366025403784438},
			nil,
		},
		{
			func(x float64) float64 {
				return -10 + math.Pow(x, 2)
			},
			-10000, 10000, 5,
			[]float64{-3.162278, 3.162278},
			nil,
		},
		{
			func(x float64) float64 {
				return -10 + 100*math.Pow(x, 2)
			},
			-1, 1, 5,
			[]float64{-0.316227, 0.316227},
			nil,
		},
	}
	for _, c := range cases {
		root, err := Brent(c.f, c.intervalStart, c.intervalEnd, c.precision)
		if err != c.expectedErr {
			t.Errorf("expected %v, got %v", c.expectedErr, err)
		}
		if err != nil {
			continue
		}
		matched := false
		i := 0
		acceptance := math.Pow10(-c.precision)
		for i < len(c.roots) && !matched {
			matched = c.roots[i]-acceptance <= root && root <= c.roots[i]+acceptance
			i++
		}
		if !matched {
			t.Errorf("expected roots are %v, got %f", c.roots, root)
		}
	}
}
