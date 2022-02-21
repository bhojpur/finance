package filters

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

type switchElem struct {
	filters []Filter
	value   interface{}
}

func (s *switchElem) Check(v interface{}) bool {
	for _, f := range s.filters {
		if f.Check(v) {
			s.value = f.Update(v)
			return true
		}
	}
	return false
}

func (s *switchElem) Update(v interface{}) interface{} {
	return s.value
}

//NewSwitch accepts a list of filters and returns Switch Filter.
//The Switch Filter evaluates all filters in sequence and
//returns true if any of the Filters is true.
func NewSwitch(fs ...Filter) Filter {
	filters := []Filter{}
	for _, f := range fs {
		filters = append(filters, f)
	}
	return &switchElem{filters: filters}
}
