package holee_test

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

	"github.com/bhojpur/finance/pkg/securities/simulation"
	"github.com/bhojpur/finance/pkg/securities/simulation/model/holee"
	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestHoLee(t *testing.T) {

	// define term structure for model calibration
	term := term.NelsonSiegelSvensson{
		-0.43381,
		-0.308942,
		4.83643,
		-4.10991,
		4.65211,
		3.33637,
		0.0,
	}

	// define parameters for ho lee model
	sigma := 0.001
	T := 5.0
	N := 5 * 255

	// create ho lee interest rate model with payoff function for zero bonds
	model, err := holee.New(&term, sigma, T, N, func(rates []float64) float64 {
		dt := T / float64(N)
		rate := 0.0
		for i := 0; i < (N - 1); i += 1 {
			rate += rates[i] * dt
		}
		return math.Exp(-rate) * 100.0
	})
	if err != nil {
		t.Error(err)
	}

	// overwriting default rng to make monte carlo more deterministic for testing
	model.Rng = rand.New(rand.NewSource(99))

	// perform monte carlo simulation with ho lee model
	engine := simulation.New(model, 1e4)
	err = engine.Run()
	if err != nil {
		t.Error(err)
	}
	zeroBondEstimate, err := engine.Estimate()
	if err != nil {
		t.Error(err)
	}
	if math.Abs(zeroBondEstimate-term.Z(T)*100.0) > 0.01 {
		t.Errorf("ho lee model failed to calculate zero bond; got: %v, expected: %v", zeroBondEstimate, term.Z(T)*100.0)
	}

}
