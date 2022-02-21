package term

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
	"encoding/json"
	"fmt"
)

var (
	registered = map[Structure][]string{
		&NelsonSiegelSvensson{}: []string{"b0", "b1", "b2", "b3", "t1", "t2", "spread"},
		&Flat{}:                 []string{"r", "spread"},
		&Spline{}:               []string{"spline", "spread"},
	}
)

func Parse(data []byte) (Structure, error) {
	// unmarshal data into map[string]interface{}
	anonymous := make(map[string]interface{})
	err := json.Unmarshal(data, &anonymous)
	if err != nil {
		return nil, err
	}
	for term, keys := range registered {
		for _, key := range keys {
			// fmt.Println("checking", key)
			if _, ok := anonymous[key]; !ok {
				// fmt.Println("key failed", key)
				// outdata, _ := json.MarshalIndent(term, " ", "")
				// fmt.Println(string(outdata))
				goto nextTerm
			}
		}
		// fmt.Printf("unmarshelling data for %T\n", term)
		err = json.Unmarshal(data, term)
		if err != nil {
			return nil, err
		}
		return term, nil
	nextTerm:
		// fmt.Println("nextTerm")
	}
	return nil, fmt.Errorf("parsing into yield curve failed")

}
