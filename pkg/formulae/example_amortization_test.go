package formulae_test

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
	"time"

	"github.com/shopspring/decimal"

	"github.com/bhojpur/finance/pkg/enums/frequency"
	"github.com/bhojpur/finance/pkg/enums/interesttype"
	"github.com/bhojpur/finance/pkg/enums/paymentperiod"
	finance "github.com/bhojpur/finance/pkg/formulae"
)

// This example generates amortization table for a loan of 20 lakhs over 15years at 12% per annum.
func ExampleAmortization_GenerateTable() {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic("location loading error")
	}
	currentDate := time.Date(2009, 11, 11, 4, 30, 0, 0, loc)
	config := finance.Config{

		// start date is inclusive
		StartDate: currentDate,

		// end date is inclusive.
		EndDate:   currentDate.AddDate(15, 0, 0).AddDate(0, 0, -1),
		Frequency: frequency.ANNUALLY,

		// AmountBorrowed is in paisa
		AmountBorrowed: decimal.NewFromInt(200000000),

		// InterestType can be flat or reducing
		InterestType: interesttype.REDUCING,

		// interest is in basis points
		Interest: decimal.NewFromInt(1200),

		// amount is paid at the end of the period
		PaymentPeriod: paymentperiod.ENDING,

		// all values will be rounded
		EnableRounding: true,

		// it will be rounded to nearest int
		RoundingPlaces: 0,

		// no error is tolerated
		RoundingErrorTolerance: decimal.Zero,
	}
	amortization, err := finance.NewAmortization(&config)
	if err != nil {
		panic(err)
	}

	rows, err := amortization.GenerateTable()
	if err != nil {
		panic(err)
	}
	finance.PrintRows(rows)
	// Output:
	// [
	//	{
	//		"Period": 1,
	//		"StartDate": "2009-11-11T04:30:00+05:30",
	//		"EndDate": "2010-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-24000000",
	//		"Principal": "-5364848"
	//	},
	//	{
	//		"Period": 2,
	//		"StartDate": "2010-11-11T00:00:00+05:30",
	//		"EndDate": "2011-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-23356218",
	//		"Principal": "-6008630"
	//	},
	//	{
	//		"Period": 3,
	//		"StartDate": "2011-11-11T00:00:00+05:30",
	//		"EndDate": "2012-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-22635183",
	//		"Principal": "-6729665"
	//	},
	//	{
	//		"Period": 4,
	//		"StartDate": "2012-11-11T00:00:00+05:30",
	//		"EndDate": "2013-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-21827623",
	//		"Principal": "-7537225"
	//	},
	//	{
	//		"Period": 5,
	//		"StartDate": "2013-11-11T00:00:00+05:30",
	//		"EndDate": "2014-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-20923156",
	//		"Principal": "-8441692"
	//	},
	//	{
	//		"Period": 6,
	//		"StartDate": "2014-11-11T00:00:00+05:30",
	//		"EndDate": "2015-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-19910153",
	//		"Principal": "-9454695"
	//	},
	//	{
	//		"Period": 7,
	//		"StartDate": "2015-11-11T00:00:00+05:30",
	//		"EndDate": "2016-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-18775589",
	//		"Principal": "-10589259"
	//	},
	//	{
	//		"Period": 8,
	//		"StartDate": "2016-11-11T00:00:00+05:30",
	//		"EndDate": "2017-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-17504878",
	//		"Principal": "-11859970"
	//	},
	//	{
	//		"Period": 9,
	//		"StartDate": "2017-11-11T00:00:00+05:30",
	//		"EndDate": "2018-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-16081682",
	//		"Principal": "-13283166"
	//	},
	//	{
	//		"Period": 10,
	//		"StartDate": "2018-11-11T00:00:00+05:30",
	//		"EndDate": "2019-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-14487702",
	//		"Principal": "-14877146"
	//	},
	//	{
	//		"Period": 11,
	//		"StartDate": "2019-11-11T00:00:00+05:30",
	//		"EndDate": "2020-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-12702445",
	//		"Principal": "-16662403"
	//	},
	//	{
	//		"Period": 12,
	//		"StartDate": "2020-11-11T00:00:00+05:30",
	//		"EndDate": "2021-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-10702956",
	//		"Principal": "-18661892"
	//	},
	//	{
	//		"Period": 13,
	//		"StartDate": "2021-11-11T00:00:00+05:30",
	//		"EndDate": "2022-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-8463529",
	//		"Principal": "-20901319"
	//	},
	//	{
	//		"Period": 14,
	//		"StartDate": "2022-11-11T00:00:00+05:30",
	//		"EndDate": "2023-11-10T23:59:59+05:30",
	//		"Payment": "-29364848",
	//		"Interest": "-5955371",
	//		"Principal": "-23409477"
	//	},
	//	{
	//		"Period": 15,
	//		"StartDate": "2023-11-11T00:00:00+05:30",
	//		"EndDate": "2024-11-10T23:59:59+05:30",
	//		"Payment": "-29364847",
	//		"Interest": "-3146234",
	//		"Principal": "-26218613"
	//	}
	// ]
}