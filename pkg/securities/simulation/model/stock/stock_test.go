package stock_test

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
	"math/rand"
	"testing"

	"github.com/bhojpur/finance/pkg/securities/instrument/option"
	"github.com/bhojpur/finance/pkg/securities/simulation"
	"github.com/bhojpur/finance/pkg/securities/simulation/model/stock"
	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestStock_IndianCall(t *testing.T) {

	// define term structure
	term := term.Flat{
		2.0,
		0.0,
	}

	// define parameters for ho lee model
	S := 110.0
	sigma := 0.3
	T := 2.0
	N := 2 * 255
	K := 100.0
	r := term.Rate(T) / 100.0

	max := func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}

	// create stock model to simulate stock prices
	model := stock.New(r, S, sigma, T, N, func(stockPrices []float64) float64 {
		return term.Z(T) * max(stockPrices[N-1]-K, 0.00)
	})

	// overwriting default rng to make monte carlo more deterministic for testing
	model.Rng = rand.New(rand.NewSource(99))

	// calculate reference price with Black Scholes
	refCall := option.Indian{
		option.Call, S, K, T, 0.0, sigma,
	}
	refCallValue := refCall.PresentValue(&term)

	// perform monte carlo simulation with ho lee model
	engine := simulation.New(model, 1e5)
	err := engine.Run()
	if err != nil {
		t.Error(err)
	}
	indianCall, err := engine.Estimate()
	if err != nil {
		t.Error(err)
	}
	if math.Abs(indianCall-refCallValue) > 0.05 {
		t.Errorf("stock model failed to price the Indian call option; got: %v, expected: %v", indianCall, refCallValue)
	}
}
