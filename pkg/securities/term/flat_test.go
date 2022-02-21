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

func TestFlat(t *testing.T) {

	rate := 1.25
	spread := 10.0

	f := term.Flat{rate, spread}

	if math.Abs(f.Rate(0.0)-f.Rate(math.Pi)) > 0.00001 {
		t.Errorf("yield curve is not flat")
	}
	if math.Abs(f.Rate(math.Pi)-(rate+spread*0.01)) > 0.00001 {
		t.Errorf("rate is not correctly calculated")
	}
	if math.Abs(f.Z(math.Pi)-math.Exp(-(rate+spread*0.01)*0.01*math.Pi)) > 0.00001 {
		t.Errorf("discount factor is not correctly calculated")
	}
}
