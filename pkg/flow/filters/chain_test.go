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

func TestChain(t *testing.T) {
	var config = []struct {
		Name   string
		Values []interface{}
		Chain  []filters.Filter
		Checks []bool
		Wants  []interface{}
	}{
		{
			Name:   "Int",
			Values: []interface{}{1, 1, 2, 2, 2, 3, 3, 3, 3, 4},
			Chain:  []filters.Filter{&filters.OnChange{}, &filters.OnValue{3}},
			Checks: []bool{false, false, false, false, false, true, false, false, false, false},
			Wants:  []interface{}{nil, nil, nil, nil, nil, 3, nil, nil, nil, nil},
		},
		{
			Name:   "Float64",
			Values: []interface{}{1.1, 1.1, 2.1, 2.1, 2.1, 3.5, 3.5, 3.5, 3.5, 4.0},
			Chain:  []filters.Filter{&filters.OnChange{}, &filters.OnValue{3.5}},
			Checks: []bool{false, false, false, false, false, true, false, false, false, false},
			Wants:  []interface{}{nil, nil, nil, nil, nil, 3.5, nil, nil, nil, nil},
		},
		{
			Name:   "String",
			Values: []interface{}{"hello", "hello", "world", "world"},
			Chain:  []filters.Filter{&filters.OnChange{}, &filters.OnValue{"world"}},
			Checks: []bool{false, false, true, false},
			Wants:  []interface{}{nil, nil, "world", nil},
		},
		{
			Name:   "FloatFilters",
			Values: []interface{}{1.0, 2.0, 2.0, 3.5, 5.0, 6.0},
			Chain:  []filters.Filter{&filters.None{}, &filters.AboveFloat64{3.0}, &filters.BelowFloat64{4.0}},
			Checks: []bool{false, false, false, true, false, false},
			Wants:  []interface{}{nil, nil, nil, 3.5, nil, nil},
		},
	}

	for _, cfg := range config {
		// internal consistency checks
		chain := filters.NewChain(cfg.Chain...)
		if len(cfg.Values) != len(cfg.Checks) {
			t.Error("internal test consistency failed (values vs. Checks)")
		}
		if len(cfg.Values) != len(cfg.Wants) {
			t.Error("internal test consistency failed (Values vs. Wants)")
		}
		// process values stream and check results for check and update
		for i, _ := range cfg.Values {
			check := chain.Check(cfg.Values[i])
			if check != cfg.Checks[i] {
				fmt.Printf("Name: %s. Got %v. Expected %v.\n", cfg.Name, check, cfg.Checks[i])
				t.Errorf("check failed")
			}
			if check {
				value := chain.Update(cfg.Values[i])
				if value != cfg.Wants[i] {
					fmt.Printf("Name: %s. Got %v. Expected %v.\n", cfg.Name, value, cfg.Wants[i])
					t.Errorf("update failed")
				}
			}
		}
	}
}
