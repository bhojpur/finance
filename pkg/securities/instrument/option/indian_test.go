package option_test

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

	"github.com/bhojpur/finance/pkg/securities/instrument/option"
	"github.com/bhojpur/finance/pkg/securities/term"
)

var (
	testOption = option.Indian{
		option.Call,
		110.0,
		100.0,
		2.0,
		0.0,
		0.3,
	}
	ts = term.Flat{2.0, 0.0}
)

func TestN(t *testing.T) {
	tests := []struct {
		Value    float64
		Expected float64
	}{
		{
			Value:    0.5311,
			Expected: 0.7023,
		},
		{
			Value:    -0.5311,
			Expected: 0.2977,
		},
		{
			Value:    0.1068,
			Expected: 0.5425,
		},
		{
			Value:    -0.1068,
			Expected: 0.4575,
		},
	}

	for _, test := range tests {
		got := option.N(test.Value)
		if math.Abs(got-test.Expected) > 0.0001 {
			t.Errorf("N function failed; Got %v, Expected %v", got, test.Expected)
		}
	}
}

func TestD1(t *testing.T) {
	value := option.D1(testOption.S, testOption.K, testOption.T, testOption.Q, testOption.Vola, &ts)
	expected := 0.5311
	if math.Abs(value-expected) > 0.0001 {
		t.Errorf("D1 function failed; Got %v, Expected %v", value, expected)
	}
}

func TestD2(t *testing.T) {
	d1 := option.D1(testOption.S, testOption.K, testOption.T, testOption.Q, testOption.Vola, &ts)
	value := option.D2(d1, testOption.T, testOption.Vola)
	expected := 0.1068
	if math.Abs(value-expected) > 0.0001 {
		t.Errorf("D2 function failed; Got %v, Expected %v", value, expected)
	}
}

func TestIndian_PresentValue1(t *testing.T) {
	call := testOption
	call.Type = option.Call

	value := call.PresentValue(&ts)
	expected := 25.1291

	if math.Abs(value-expected) > 0.0001 {
		t.Errorf("pricing of Indian call failed; Got %v, Expected %v", value, expected)
	}
}

func TestIndian_PresentValue2(t *testing.T) {
	put := testOption
	put.Type = option.Put

	value := put.PresentValue(&ts)
	expected := 11.2080

	if math.Abs(value-expected) > 0.0001 {
		t.Errorf("pricing of Indian put failed; Got %v, Expected %v", value, expected)
	}
}

func TestIndian_Greeks(t *testing.T) {

	testData := []struct {
		Type          int
		ExpectedDelta float64
		ExpectedGamma float64
		ExpectedRho   float64
		ExpectedVega  float64
	}{
		{
			Type:          option.Call,
			ExpectedDelta: 0.7023,
			ExpectedGamma: 0.0074,
			ExpectedRho:   104.2505,
			ExpectedVega:  53.8985,
		},
		{
			Type:          option.Put,
			ExpectedDelta: -0.2977,
			ExpectedGamma: 0.0074,
			ExpectedRho:   -87.9074,
			ExpectedVega:  53.8985,
		},
	}

	tolerance := 0.0001

	for _, test := range testData {
		opt := testOption
		opt.Type = test.Type

		delta := opt.Delta(&ts)
		if math.Abs(delta-test.ExpectedDelta) > tolerance {
			t.Errorf("delta is not correct; Got %v, Expected %v", delta, test.ExpectedDelta)
		}

		gamma := opt.Gamma(&ts)
		if math.Abs(gamma-test.ExpectedGamma) > tolerance {
			t.Errorf("gamma is not correct; Got %v, Expected %v", gamma, test.ExpectedGamma)
		}
		rho := opt.Rho(&ts)
		if math.Abs(rho-test.ExpectedRho) > tolerance {
			t.Errorf("rho is not correct; Got %v, Expected %v", rho, test.ExpectedRho)
		}
		vega := opt.Vega(&ts)
		if math.Abs(vega-test.ExpectedVega) > tolerance {
			t.Errorf("vega is not correct; Got %v, Expected %v", vega, test.ExpectedVega)
		}
	}

}
