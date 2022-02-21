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
	"fmt"
	"math"
	"testing"
	"time"

	securities "github.com/bhojpur/finance/pkg/securities"
	"github.com/bhojpur/finance/pkg/securities/instrument/bond"
	"github.com/bhojpur/finance/pkg/securities/instrument/option"
	"github.com/bhojpur/finance/pkg/securities/maturity"
	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestYields(t *testing.T) {

	testData := []struct {
		B              *bond.Straight
		Quote          float64
		ExpectedIRR    float64
		ExpectedSpread float64
	}{
		{
			// bond details
			// ISIN CH0224396983 (quote per 2021-04-01)
			B: &bond.Straight{
				Schedule: maturity.Schedule{
					Settlement: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC),
					Maturity:   time.Date(2026, 5, 28, 0, 0, 0, 0, time.UTC),
					Frequency:  1,
				},
				Redemption: 100.0,
				Coupon:     1.25,
			},
			Quote:          109.70,
			ExpectedIRR:    -0.574,
			ExpectedSpread: 0.0,
		},
		{
			// ISIN CH0193265995 (quote per 2021-04-16)
			B: &bond.Straight{
				Schedule: maturity.Schedule{
					Settlement: time.Date(2021, 4, 15, 0, 0, 0, 0, time.UTC),
					Maturity:   time.Date(2022, 9, 21, 0, 0, 0, 0, time.UTC),
					Frequency:  1,
				},
				Redemption: 100.0,
				Coupon:     1.00,
			},
			Quote:          102.22,
			ExpectedIRR:    -0.54,
			ExpectedSpread: 12.432,
		},
	}

	// term structure (parameters per 2021-04-01 for CH govt bonds)
	term := term.NelsonSiegelSvensson{
		-0.266372,
		-0.471343,
		5.68789,
		-5.12324,
		5.74881,
		4.14426,
		0.0, // spread
	}

	// loop over tests
	for nr, test := range testData {

		// IRR
		irr, err := securities.Irr(test.Quote+test.B.Accrued(), test.B)
		// fmt.Println(irr)
		if err != nil {
			fmt.Println(err)
			t.Errorf("irr failed for test nr %d", nr)
		}

		// fmt.Println("Remaining years:", test.B.RemainingYears())

		if math.Abs(irr-test.ExpectedIRR) > 0.1 {
			t.Errorf("wrong IRR for test nr %d, got %f, expected %f", nr, irr, test.ExpectedIRR)
		}

		// Z-Spread
		spread, err := securities.Spread(test.Quote+test.B.Accrued(), test.B, &term)
		// fmt.Println(zspread)
		if err != nil {
			fmt.Println(err)
			t.Errorf("zspread failed for test nr %d", nr)
		}

		if math.Abs(spread-test.ExpectedSpread) > 0.1 {
			t.Errorf("wrong Z-Spread for test nr %d, got %f, expected %f", nr, spread, test.ExpectedSpread)
		}
	}
}

func TestImpliedVola(t *testing.T) {
	testOption := option.Indian{
		option.Call,
		110.0,
		100.0,
		2.0,
		0.0,
		math.Pi,
	}
	ts := term.Flat{2.0, 0.0}

	expected := 0.3

	tests := []struct {
		OptionType int
		Price      float64
	}{
		{
			OptionType: option.Call,
			Price:      25.1291,
		},
		{
			OptionType: option.Put,
			Price:      11.2080,
		},
	}

	for _, test := range tests {
		testOption.Type = test.OptionType
		vola, err := securities.ImpliedVola(test.Price, &testOption, &ts)
		if err != nil || math.Abs(vola-expected) > 0.0001 {
			t.Error("Type", test.OptionType, "Got", vola, "Expected", expected)
		}
	}
}
