package term_test

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
	"reflect"
	"testing"

	"github.com/bhojpur/finance/pkg/securities/term"
)

func TestParse(t *testing.T) {

	testData := []struct {
		Data []byte
		Type interface{}
	}{
		{
			Data: []byte(" { \"b0\": -0.596356, \"b1\": -0.153952, \"b2\": 5.79009, \"b3\": -4.69599, \"t1\": 6.5912, \"t2\": 4.63027, \"spread\": 0.0 } "),
			Type: &term.NelsonSiegelSvensson{},
		},
		{
			Data: []byte(" { \"spline\": null, \"spread\": 0.0 } "),
			Type: &term.Spline{},
		},
		{
			Data: []byte(" { \"r\": 0.0, \"spread\": 0.0 } "),
			Type: &term.Flat{},
		},
	}

	for i, test := range testData {
		ts, err := term.Parse(test.Data)
		// fmt.Printf("type=%T\n", ts)
		// dataout, _ := json.Marshal(ts)
		// fmt.Println(string(dataout))
		if err != nil {
			t.Errorf("test %d: %v", i+1, err)
		}
		if reflect.TypeOf(ts) != reflect.TypeOf(test.Type) {
			t.Errorf("test %d: parse returned wrong type: got: %T, expected %T", i+1, ts, test.Type)
		}
	}
}
