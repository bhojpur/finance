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

	"github.com/bhojpur/finance/pkg/securities/instrument/swap"
	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestFxRate(t *testing.T) {
	// Example:
	// A US firm issues 100 million Euro-denominated 5 year note with coupon c=4%.
	// The firm exchanges the proceeds into $126.73 million at the current
	// exchange rate M0 = 1.2673$/Eur
	// Every 6 months, the firm must pay EUR 2 mil (100 mil * 4%/2). In addition,
	// at T=5, the firm must pay back Eur 100 mil principal

	// Plain vanilly FX swap to hedge exchange risk: what is swap rate K?

	// Assume rates are constant across maturities
	termEur := term.Flat{4.0, 0.0}
	termUsd := term.Flat{6.0, 0.0}

	m0 := 1.2673

	swapRate, err := swap.FxRate(m0,
		[]float64{2.0, 2.0, 2.0, 2.0, 2.0, 2.0, 2.0, 2.0, 2.0, 102.0},
		[]float64{0.5, 1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5.0},
		&termUsd,
		&termEur,
	)
	if err != nil {
		t.Error(err)
	}

	expectedRate := 1.389

	// fmt.Println("Swaprate", swapRate, "Expected", expectedRate)

	if math.Abs(swapRate-expectedRate) > 0.001 {
		t.Error("fx swap rate calculation is wrong; got:", swapRate, "expected:", expectedRate)
	}

}

func TestInterestRate(t *testing.T) {

	// calculating 5-year semi-annual CHF Swap Rate
	// IRS CHF 5Y (CH0002113865)

	// Parameters for CHF yield curve at Nov-30-2021
	// added 9bps z-spread for counterparty risk
	term := term.NelsonSiegelSvensson{
		-0.43381,
		-0.308942,
		4.83643,
		-4.10991,
		4.65211,
		3.33637,
		9.0,
	}

	maturities := []float64{
		3, 4, 5, 6, 7, 8, 9, 10,
	}
	expectedSwapRate := []float64{
		-0.50, -0.42, -0.35, -0.28, -0.21, -0.15, -0.10, -0.06,
	}

	for i, m := range maturities {
		tm := []float64{}
		for k := 0.5; k <= m; k += 0.5 {
			tm = append(tm, k)
		}
		swapRate, err := swap.InterestRate(tm, 2, &term)
		if err != nil {
			t.Error(err)
		}
		expectedRate := expectedSwapRate[i]
		if math.Abs(swapRate-expectedRate) > 0.05 {
			t.Error("interest swap rate calculation is wrong; maturity:", m, "got:", swapRate, "expected:", expectedRate)
		}
	}

}
