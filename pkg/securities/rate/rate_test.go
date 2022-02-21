package rate_test

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

	"github.com/bhojpur/finance/pkg/securities/rate"
)

var (
	annualRate           = 12.0
	monthlyEffectiveRate = math.Pow(1.01, 12)
	ccRate               = 11.332869
)

func TestContinuous(t *testing.T) {
	if math.Abs(rate.Continuous(annualRate, 1)-ccRate) > 1e-6 {
		t.Errorf("conversion from annual to cc rate failed")
	}
}

func TestAnnual(t *testing.T) {
	if math.Abs(rate.Annual(ccRate, 1)-annualRate) > 1e-6 {
		t.Errorf("conversion from cc to annual rate failed")
	}
}

func TestEffectiveAnnual(t *testing.T) {
	n := 12 // monthly compounding
	if math.Abs(rate.EffectiveAnnual(annualRate, n)-monthlyEffectiveRate) > 1e-6 {
		t.Errorf("conversion from cc to annual rate failed")
	}
}
