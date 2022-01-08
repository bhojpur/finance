package formulae

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

import "github.com/shopspring/decimal"

// Flat implements financial methods for facilitating a loan use case, following a flat rate of interest.
type Flat struct{}

// GetPrincipal returns principal amount contribution in a given period towards a loan, depending on config.
func (f *Flat) GetPrincipal(config Config, _ int64) decimal.Decimal {
	dPeriod := decimal.NewFromInt(config.periods)
	minusOne := decimal.NewFromInt(-1)
	return config.AmountBorrowed.Div(dPeriod).Mul(minusOne)
}

// GetInterest returns interest amount contribution in a given period towards a loan, depending on config.
func (f *Flat) GetInterest(config Config, period int64) decimal.Decimal {
	minusOne := decimal.NewFromInt(-1)
	return config.getInterestRatePerPeriodInDecimal().Mul(config.AmountBorrowed).Mul(minusOne)
}

// GetPayment returns the periodic payment to be done for a loan depending on config.
func (f *Flat) GetPayment(config Config) decimal.Decimal {
	dPeriod := decimal.NewFromInt(config.periods)
	minusOne := decimal.NewFromInt(-1)
	totalInterest := config.getInterestRatePerPeriodInDecimal().Mul(dPeriod).Mul(config.AmountBorrowed)
	Payment := totalInterest.Add(config.AmountBorrowed).Mul(minusOne).Div(dPeriod)
	return Payment
}
