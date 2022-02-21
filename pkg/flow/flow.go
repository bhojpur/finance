package flow

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

	"github.com/bhojpur/finance/pkg/flow/filters"
	"github.com/bhojpur/finance/pkg/flow/observer"
)

//New return an Observer that receive the results of the flow
func New(nf filters.Filter, s Source) observer.Observer {
	return s.Run(nf)
}

//Source is the interface for input for the flow
type Source interface {
	Run(f filters.Filter) observer.Observer
}

//Func implements the Source interface and regularly calls a function
type Func struct {
	Fn      func() interface{}
	Refresh time.Duration
}

//Run calls the given function in regular intervals
func (f *Func) Run(nf filters.Filter) observer.Observer {
	o := observer.NewObserver()
	c := time.Tick(f.Refresh)
	go func() {
		for {
			select {
			case <-c:
				if v := f.Fn(); nf.Check(v) {
					o.Notify(nf.Update(v))
				}
			case <-o.Control().C:
				o.Control().D <- true
				return
			}
		}
	}()
	return o
}

//Chan implements the Source interface and provides the input for the flow
type Chan struct {
	Ch chan interface{}
}

//Run passed the channel data to the filters
func (c *Chan) Run(nf filters.Filter) observer.Observer {
	o := observer.NewObserver()
	go func() {
		for {
			select {
			case v := <-c.Ch:
				if nf.Check(v) {
					o.Notify(nf.Update(v))
				}
			case <-o.Control().C:
				o.Control().D <- true
				return
			}
		}
	}()
	return o
}
