package swap_test

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

	"github.com/bhojpur/finance/pkg/securities/instrument/bond"
	"github.com/bhojpur/finance/pkg/securities/instrument/swap"
	"github.com/bhojpur/finance/pkg/securities/maturity"
	"github.com/bhojpur/finance/pkg/securities/rate"
	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestInterestRateSwap(t *testing.T) {
	// Parameters for CHF yield curve at Nov-30-2021
	// added 9bps z-spread for counterparty risk
	term := term.NelsonSiegelSvensson{
		-0.43381,
		-0.308942,
		4.83643,
		-4.10991,
		4.65211,
		3.33637,
		10.0,
	}

	maturities := []int{
		3, 4, 5, 6, 7, 8, 9, 10,
	}
	swapRate := []float64{
		-0.50, -0.42, -0.35, -0.28, -0.21, -0.15, -0.10, -0.06,
	}

	date := time.Date(2021, 12, 3, 0, 0, 0, 0, time.UTC)
	for i, m := range maturities {
		scheduleFloating := maturity.Schedule{
			Settlement: date,
			Maturity:   date.AddDate(m, 0, 0),
			Frequency:  2,
			Basis:      "ACT360",
		}
		scheduleFixed := maturity.Schedule{
			Settlement: date,
			Maturity:   date.AddDate(m, 0, 0),
			Frequency:  1,
			Basis:      "30E360",
		}

		floatingLeg := bond.Floating{
			Schedule:   scheduleFloating,
			Rate:       rate.Annual(term.Rate(0.5), 2),
			Redemption: 100.0,
		}

		fixedLeg := bond.Straight{
			Schedule:   scheduleFixed,
			Coupon:     swapRate[i],
			Redemption: 100.0,
		}

		swapSecurity := swap.InterestRateSwap{
			Floating: floatingLeg,
			Fixed:    fixedLeg,
		}

		value := swapSecurity.PresentValue(&term)

		if math.Abs(value) > 0.25 {
			t.Error("value of interest rate swap is wrong; maturity:", m, "got:", value, "expected: 0.0")
		}
	}

}
