package simulation_test

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
)

type pi struct {
	Rng *rand.Rand
}

func NewPi() pi {
	rng := rand.New(rand.NewSource(99))
	return pi{rng}
}

func (p pi) Measurement() float64 {
	x, y := p.Rng.Float64(), p.Rng.Float64()
	if x*x+y*y <= 1.0 {
		return 1.0
	}
	return 0.0
}

func TestEngine(t *testing.T) {
	engine := simulation.New(
		NewPi(),
		1e6,
	)
	err := engine.Run()
	if err != nil {
		t.Error(err)
	}
	average, err := engine.Estimate()
	if err != nil {
		t.Error(err)
	}
	if math.Abs(average-math.Pi/4.0) > 0.001 {
		t.Errorf("monte carlo estimate failed; got: %v, expected: %v", average, math.Pi/4.0)
	}
	stderror, err := engine.StdError()
	if err != nil {
		t.Error(err)
	}
	if math.Abs(stderror-0.0004) > 0.0001 {
		t.Errorf("monte carlo stderror failed; got: %v, expected: %v", stderror, 0.0004)
	}
	lower, upper, err := engine.CI()
	if err != nil {
		t.Error(err)
	}
	if math.Abs((lower-average)/stderror-(-1.96)) > 0.001 {
		t.Errorf("monte carlo lower CI failed; got: %v, expected: %v", lower, -1.96)
	}
	if math.Abs((upper-average)/stderror-1.96) > 0.001 {
		t.Errorf("monte carlo upper CI failed; got: %v, expected: %v", upper, 1.96)
	}

}
