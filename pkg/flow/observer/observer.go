package observer

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
	"sync"
)

//Observer defines the interface for the observer design patter
type Observer interface {
	Notify(interface{})
	Subscribe() Subscriber
	Control() control
	Close()
}

//NewObserver returns an implementation of the observer interface
func NewObserver() Observer {
	o := observerI{
		control: NewControl(),
		state:   NewState(),
	}
	return &o
}

//observer
type observerI struct {
	sync.RWMutex //embedded
	control      //embedded
	state        *state
}

//Notify sends out the current value in the observer channel
func (o *observerI) Notify(value interface{}) {
	o.Lock()
	defer o.Unlock()
	o.state.Value = value
	next := NewState()
	o.state.Next = next
	close(o.state.C)
	o.state = o.state.Next
}

//Subscribe returns a new subscriber to access values and listens for events
func (o *observerI) Subscribe() Subscriber {
	o.RLock()
	defer o.RUnlock()
	return &subscriber{state: o.state}
}

//Control
func (o *observerI) Control() control {
	return o.control
}
