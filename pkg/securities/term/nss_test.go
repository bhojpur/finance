package term_test

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

	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestNelsonSiegelSvensson(t *testing.T) {
	// NSS parameters
	n := term.NelsonSiegelSvensson{
		-0.266372,
		-0.471343,
		5.68789,
		-5.12324,
		5.74881,
		4.14426,
		0.0, // spread
	}

	data := []struct {
		M             float64
		RateInPercent float64
	}{
		{1, -0.782},
		{2, -0.776},
		{3, -0.736},
		{4, -0.677},
		{5, -0.606},
		{6, -0.532},
		{7, -0.459},
		{8, -0.389},
		{9, -0.325},
		{10, -0.267},
		{15, -0.073},
		{20, -0.001},
		{30, -0.007},
	}

	for _, test := range data {
		got := n.Rate(test.M)
		expected := math.Log(1.0+test.RateInPercent/100.0) * 100.0
		if math.Abs(got-expected) > 0.001 {
			t.Errorf("got %f, but wanted %f, failed for maturity %f", got, expected, test.M)
		}
	}
}
