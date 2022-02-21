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

type chain struct {
	fs    []Filter
	Value interface{}
	Flag  bool
}

func (c *chain) Check(v interface{}) bool {
	c.Value = v
	c.Flag = true
	for _, f := range c.fs {
		if f.Check(c.Value) {
			c.Value = f.Update(c.Value)
		} else {
			c.Flag = false
			break
		}
	}
	return c.Flag
}

func (c *chain) Update(v interface{}) interface{} {
	return c.Value
}

//NewChain chains together filters.
func NewChain(filters ...Filter) Filter {
	chainedFilters := make([]Filter, 0)
	for _, f := range filters {
		chainedFilters = append(chainedFilters, f)
	}
	return &chain{fs: chainedFilters}
}
