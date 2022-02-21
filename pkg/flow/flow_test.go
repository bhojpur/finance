package flow_test

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
	"time"

	"github.com/bhojpur/finance/pkg/flow"
	"github.com/bhojpur/finance/pkg/flow/filters"
)

var config = []struct {
	Values []interface{}
	Want   interface{}
}{
	{
		Values: []interface{}{1.1, 1.1, 1.1, 2.1, 1.1},
		Want:   2.1,
	},
	{
		Values: []interface{}{1, 1, 1, 2, 1},
		Want:   2,
	},
	{
		Values: []interface{}{"hello", "hello", "hello", "world", "hello"},
		Want:   "world",
	},
}

var observers = []struct {
	Name   string
	TrFunc func(v interface{}) filters.Filter
}{
	{
		Name: "OnChange",
		TrFunc: func(v interface{}) filters.Filter {
			return &filters.OnChange{v}
		},
	},
	{
		Name: "OnValue",
		TrFunc: func(v interface{}) filters.Filter {
			return &filters.OnValue{v}
		},
	},
}

func TestChannelObservers(t *testing.T) {

	for _, cfg := range config {
		for _, observerCfg := range observers {

			// prepare test
			startC := make(chan bool, 1)
			ch := make(chan interface{}, 1)
			values := cfg.Values
			go func() {
				<-startC
				for _, v := range values {
					ch <- v
					time.Sleep(10 * time.Millisecond)
				}
			}()

			// get start value
			var start interface{}
			switch observerCfg.Name {
			case "OnChange":
				start = values[0]
			case "OnValue":
				start = cfg.Want
			}

			// create observer
			observer := flow.New(observerCfg.TrFunc(start), &flow.Chan{ch})
			subscriber := observer.Subscribe()
			startC <- true

			// run test
			select {
			case <-time.After(1 * time.Second):
				str := fmt.Sprintf("%s: Timed out waiting for channel.", observerCfg.Name)
				t.Fatal(str)
			case <-subscriber.C():
				received := subscriber.Value()
				if received != cfg.Want {
					str := fmt.Sprintf("%s: Got %v. Expected %v", observerCfg.Name, received, cfg.Want)
					t.Fatal(str)
				}
			}

			// close
			observer.Close()
		}
	}
}

func TestIntervalObservers(t *testing.T) {

	refresh := 10 * time.Millisecond

	for _, cfg := range config {
		for _, observerCfg := range observers {

			// prepare test
			values := cfg.Values
			var index int
			fn := func() interface{} {
				if index > len(values) {
					t.Error("Ran out of values.")
				}
				v := values[index]
				index++
				return v
			}

			// get start value
			var start interface{}
			switch observerCfg.Name {
			case "OnChange":
				start = values[0]
			case "OnValue":
				start = cfg.Want
			}

			// create observer
			observer := flow.New(observerCfg.TrFunc(start), &flow.Func{fn, refresh})
			subscriber := observer.Subscribe()
			// run test
			select {
			case <-time.After(1 * time.Second):
				str := fmt.Sprintf("%s: Timed out waiting for channel.", observerCfg.Name)
				t.Fatal(str)
			case <-subscriber.C():
				received := subscriber.Value()
				if received != cfg.Want {
					str := fmt.Sprintf("%s: Got %v. Expected %v", observerCfg.Name, received, cfg.Want)
					t.Fatal(str)
				}
			}

			// close
			observer.Close()
		}
	}
}
