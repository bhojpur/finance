package stock

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
	"time"
)

// Stock implements the log-normal simulation for the Monte Carlo engine
type Stock struct {
	// R is annual return until maturity
	R float64
	// S0 is the stock value at the time 0
	S0 float64
	// Sigma is the standard deviation of the stock
	Sigma float64
	// T is the maturity (up to which to calculate the interes rates)
	T float64
	// N represents number of steps
	N int
	// Rng is the random number generator (NormFloat64)
	Rng *rand.Rand
	// Payoff returns the discounted payoff for the given simulated rates
	Payoff func([]float64) float64
}

// New creates a new Stock model for simulation
func New(r, s0, sigma, t float64, n int, payoff func([]float64) float64) *Stock {
	s := &Stock{
		R:      r,
		S0:     s0,
		Sigma:  sigma,
		T:      t,
		N:      n,
		Rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Payoff: payoff,
	}
	return s
}

// Measurement implements the Monte Carlo model interface
func (s *Stock) Measurement() float64 {
	n := s.N
	dt := s.T / float64(n)
	stockValues := make([]float64, n)

	// simulate interest rates
	stockValues[0] = s.S0
	for i := 0; i < (n - 1); i += 1 {
		stockValues[i+1] = stockValues[i] * math.Exp((s.R-math.Pow(s.Sigma, 2.0)/2.0)*dt+s.Sigma*math.Sqrt(dt)*s.Rng.NormFloat64())
	}
	return s.Payoff(stockValues)
}
