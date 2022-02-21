package securities_test

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
	"time"

	securities "github.com/bhojpur/finance/pkg/securities"
	"github.com/bhojpur/finance/pkg/securities/instrument/bond"
	"github.com/bhojpur/finance/pkg/securities/maturity"
	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestInterestSensitivity(t *testing.T) {
	// define zero bond
	bond := bond.Straight{
		Schedule: maturity.Schedule{
			Settlement: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			Maturity:   time.Date(2031, 1, 1, 0, 0, 0, 0, time.UTC),
			Frequency:  1,
			Basis:      "30E360",
		},
		Coupon:     0.0,
		Redemption: 100.0,
	}
	// define flat yield curve
	ts := term.Flat{2.0, 0.0}

	// calculate sensitivity
	pvbp := securities.PVBP(&bond, &ts)
	// fmt.Println("duration=", bond.Duration(&ts))
	// fmt.Println("convexity=", bond.Convexity(&ts))

	ts2 := ts
	pvbpRef := bond.PresentValue(ts2.SetSpread(1.0)) - bond.PresentValue(&ts)

	if math.Abs(pvbp-pvbpRef) > 0.0001 {
		t.Errorf("pvbp calculation failed; got: %v, expected: %v", pvbp, pvbpRef)
	}
}
