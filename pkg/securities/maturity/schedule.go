package maturity

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
	"sort"
	"time"

	"github.com/bhojpur/finance/pkg/formulae/daycount"
)

// Schedule contain the information about the term maturities of fixed income security's cash flows
type Schedule struct {
	// Settlement represent the date of valuation (or settlement)
	Settlement time.Time
	// Maturity represents the maturity date
	Maturity time.Time
	// Frequency is the compounding frequency per year (default: 1x per year)
	Frequency int
	// Basis represents the day count convention (default: "" for 30E/360 ISDA)
	Basis string
}

//Compounding returns the annual compounding frequency
func (m *Schedule) Compounding() int {
	n := 1
	if m.Frequency > 0 {
		n = m.Frequency
	}
	return n
}

//EffectiveCoupon calculates the coupon that is payable given the annual coupon rate
func (m *Schedule) EffectiveCoupon(annualCoupon float64) float64 {
	n := float64(m.Compounding())
	return annualCoupon / n
}

//M returns a slice of the effective maturities in years of the bond's cash flows
func (m *Schedule) M() []float64 {
	maturities := []float64{}

	if m.Compounding() > 12 {
		panic("more than 12 compounding periods not implemented yet")
	}
	step := 12 / m.Compounding()

	// walk back from maturity date to quote date
	quote := m.Settlement
	for current := m.Maturity; current.Sub(quote) > 0; current = current.AddDate(0, -step, 0) {
		frac, err := daycount.Fraction(quote, current, quote.AddDate(1, 0, 0), m.Basis)
		if err != nil {
			panic(err)
		}
		maturities = append(maturities, frac)
	}

	return maturities
}

//Last returns the latest maturity value in years (i.e. the years to maturity)
func (m *Schedule) Last() float64 {
	t := m.M()
	if len(t) == 0 {
		return 0.0
	}
	sort.Float64s(t)
	return t[len(t)-1]
}

//Next returns the next maturity value in years
func (m *Schedule) Next() float64 {
	t := m.M()
	if len(t) == 0 {
		return 0.0
	}
	sort.Float64s(t)
	return t[0]
}

// DayCountFraction returns year fraction since last coupon
func (m *Schedule) DayCountFraction() float64 {
	if m.Maturity.Before(m.Settlement) {
		return 0.0
	}

	d1 := m.Maturity
	d2 := m.Settlement
	d3 := time.Time{}

	// iterate maturity date backwards until last coupon date before settlement date
	if m.Compounding() > 12 {
		panic("more than 12 compounding periods not implemented yet")
	}
	step := 12 / m.Compounding()
	for ; d1.Sub(d2) > 0; d1 = d1.AddDate(0, -step, 0) {
		d3 = d1
	}

	// calculate day count fraction
	frac, err := daycount.Fraction(d1, d2, d3, m.Basis)
	if err != nil {
		panic(err)
	}

	return frac / float64(m.Compounding())
}

// Actual difference between two dates in years
// func ActualDifferenceInYears(start, stop time.Time) float64 {
// 	years := 0.0
// 	// same year, just take difference in days and divided by numbers of days in year
// 	if start.Year() == stop.Year() {
// 		years = float64(stop.YearDay()-start.YearDay()) / DaysInYear(start.Year())
// 	} else {
// 		// "maturity" for current year
// 		years += 1.0 - float64(start.YearDay())/DaysInYear(start.Year())
// 		// "maturity" for last year
// 		years += float64(stop.YearDay()) / DaysInYear(stop.Year())
// 		// "maturity" for years in between
// 		for y := start.Year() + 1; y < stop.Year(); y += 1 {
// 			years += 1.0
// 		}
// 	}
// 	// hour adjustment
// 	years -= float64(start.Hour()) / 24.0 / DaysInYear(start.Year())
// 	years += float64(stop.Hour()) / 24.0 / DaysInYear(stop.Year())
//
// 	return years
// }
