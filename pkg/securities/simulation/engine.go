package simulation

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
)

const (
	IncompleteSetup = iota
	Initialized
	Running
	ResultsAvailable
)

// Model specifies the interface for using the Monte Carlo engine
type Model interface {
	Measurement() float64
}

// Engine implements the Monte Carlo simulation
type Engine struct {
	Model     Model
	Nsim      int
	Estimates []float64
	Status    int
}

// New creates a Monte Carlo simulation engine for the given model
func New(m Model, nsim int) *Engine {
	e := Engine{
		Model:     m,
		Nsim:      nsim,
		Estimates: make([]float64, nsim),
		Status:    Initialized,
	}
	return &e
}

// Run runs the Monte Carlo simulation
func (e *Engine) Run() error {
	if e.Status != Initialized {
		return fmt.Errorf("Monte Carlo engine not initialized")
	}
	e.Status = Running
	for i := 0; i < e.Nsim; i += 1 {
		e.Estimates[i] = e.Model.Measurement()
	}
	e.Status = ResultsAvailable
	return nil
}

// Estimate returns the estimate of the simulation (average over all simulations)
func (e *Engine) Estimate() (float64, error) {
	if e.Status != ResultsAvailable {
		return 0.0, fmt.Errorf("no results available from Monte Carlo simulation")
	}
	var value float64
	for _, estimate := range e.Estimates {
		value += estimate
	}
	return value / float64(len(e.Estimates)), nil
}

// StdError returns the standard error of the simulations
func (e *Engine) StdError() (float64, error) {
	if e.Status != ResultsAvailable {
		return 0.0, fmt.Errorf("no results available from Monte Carlo simulation")
	}
	average, err := e.Estimate()
	if err != nil {
		return 0.0, err
	}
	stderror := 0.0
	for _, estimate := range e.Estimates {
		stderror += math.Pow(estimate-average, 2.0)
	}
	size := float64(len(e.Estimates))
	return math.Sqrt(stderror/size) / math.Sqrt(size), nil

}

// CI returns the 95% confidence interval
func (e *Engine) CI() (float64, float64, error) {
	average, err := e.Estimate()
	if err != nil {
		return 0.0, 0.0, err
	}
	stderror, err := e.StdError()
	if err != nil {
		return 0.0, 0.0, err
	}

	return average - 1.96*stderror, average + 1.96*stderror, nil
}
