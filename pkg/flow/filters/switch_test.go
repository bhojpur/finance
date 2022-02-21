package filters_test

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
	"testing"

	"github.com/bhojpur/finance/pkg/flow/filters"
)

func TestSwitch(t *testing.T) {
	var config = []struct {
		Name   string
		Values []interface{}
		Left   filters.Filter
		Right  filters.Filter
		Checks []bool
		Wants  []interface{}
	}{
		{
			Name:   "Float64Collar",
			Values: []interface{}{0.0, 0.5, 1.0, 1.1, -1.0, -1.2, 0.0},
			Left:   &filters.AboveFloat64{1.0},
			Right:  &filters.BelowFloat64{-1.0},
			Checks: []bool{false, false, false, true, false, true, false},
			Wants:  []interface{}{nil, nil, nil, 1.1, nil, -1.2, nil},
		},
	}

	for _, cfg := range config {
		// internal consistency checks
		or := filters.NewSwitch(cfg.Left, cfg.Right)
		if len(cfg.Values) != len(cfg.Checks) {
			t.Error("internal test consistency failed (values vs. Checks)")
		}
		if len(cfg.Values) != len(cfg.Wants) {
			t.Error("internal test consistency failed (values vs. Wants)")
		}
		// process values stream and check results for check and update
		for i, _ := range cfg.Values {
			check := or.Check(cfg.Values[i])
			if check != cfg.Checks[i] {
				fmt.Printf("Name: %s. Got %v. Expected %v.\n", cfg.Name, check, cfg.Checks[i])
				t.Error("check failed")
			}
			if check {
				value := or.Update(cfg.Values[i])
				if value != cfg.Wants[i] {
					fmt.Printf("Name: %s. Got %v. Expected %v.\n", cfg.Name, value, cfg.Wants[i])
					t.Error("update failed")
				}
			}
		}
	}
}
