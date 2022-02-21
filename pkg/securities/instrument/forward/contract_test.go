package forward_test

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

	"github.com/bhojpur/finance/pkg/securities/instrument/forward"
	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestContract(t *testing.T) {
	ts := term.Flat{2.0, 0.0}
	forwardPrice, _ := forward.ZeroBondPrice(1.0, 2.0, &ts)
	contract := forward.Contract{
		K: forwardPrice,
		F: ts.Z(2.0) / ts.Z(1.0),
		T: 1.0,
	}
	value := contract.PresentValue(&ts)
	expected := 0.0
	if math.Abs(value-expected) > 0.00001 {
		t.Errorf("wrong forward contract value; got: %v, expected: %v", value, expected)
	}
}

func TestZeroBondPrice(t *testing.T) {
	ts := term.Flat{2.0, 0.0}
	forwardPrice, err := forward.ZeroBondPrice(1.0, 2.0, &ts)
	if err != nil {
		t.Error(err)
	}
	expected := ts.Z(2.0) / ts.Z(1.0)
	if math.Abs(forwardPrice-expected) > 0.00001 {
		t.Errorf("wrong zero-bond forward price; got: %v, expected: %v", forwardPrice, expected)
	}
}

func TestFx(t *testing.T) {
	// t.Errorf("not implemented yet")
}

func TestStockPrice(t *testing.T) {
	// t.Errorf("not implemented yet")
}
