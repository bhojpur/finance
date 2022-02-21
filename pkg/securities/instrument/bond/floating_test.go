package bond_test

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
	"github.com/bhojpur/finance/pkg/securities/maturity"
	"github.com/bhojpur/finance/pkg/securities/rate"
	"github.com/bhojpur/finance/pkg/securities/term"
)

var (
	date         = time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)
	floatingBond = bond.Floating{
		Schedule: maturity.Schedule{
			Settlement: date,
			Maturity:   date.AddDate(0, 12, 0),
			Frequency:  2,
		},
		Rate:       2.00,
		Redemption: 100.0,
	}
	floatTwo = bond.Floating{
		Schedule: maturity.Schedule{
			Settlement: date.AddDate(0, 3, 0),
			Maturity:   date.AddDate(0, 12, 0),
			Frequency:  2,
		},
		Rate:       2.00,
		Redemption: 100.0,
	}
	floatingTerm = term.Flat{
		rate.Continuous(2.0, floatingBond.Schedule.Compounding()),
		0.0,
	}
)

func TestFloating_PresentValue(t *testing.T) {

	testData := []struct {
		Floating bond.Floating
		Expected float64
	}{
		{
			Floating: floatingBond,
			Expected: 100.0,
		},
		{
			Floating: floatTwo,
			Expected: 100.4975,
		},
	}

	for i, test := range testData {
		// PresentValue delivers the "clean" price
		dirty := test.Floating.PresentValue(&floatingTerm)
		expected := test.Expected
		if math.Abs(dirty-expected) > 0.01 {
			t.Errorf("test nr %d, got %f, expected %f", i, dirty, expected)
		}

	}

}

func TestFloating_Accrued(t *testing.T) {
	accrued := floatTwo.Accrued()
	expected := 0.5
	if math.Abs(accrued-expected) > 0.001 {
		t.Errorf("got %f, expected %f", accrued, expected)
	}
}

func TestFloating_DurationConvexity(t *testing.T) {

	testData := []struct {
		Floating          bond.Floating
		ExpectedDuration  float64
		ExpectedConvexity float64
	}{
		{
			Floating:          floatingBond,
			ExpectedDuration:  -0.5,
			ExpectedConvexity: 0.25,
		},
		{
			Floating:          floatTwo,
			ExpectedDuration:  -0.25,
			ExpectedConvexity: 0.06,
		},
	}

	for nr, test := range testData {

		duration := test.Floating.Duration(&floatingTerm)
		if math.Abs(duration-test.ExpectedDuration) > 0.01 {
			t.Errorf("test nr %d, duration failed, got %f, expected %f", nr, duration, test.ExpectedDuration)
		}
		convex := test.Floating.Convexity(&floatingTerm)
		if math.Abs(convex-test.ExpectedConvexity) > 0.01 {
			t.Errorf("test nr %d, convexity failed, got %f, expected %f", nr, convex, test.ExpectedConvexity)
		}
	}
}
