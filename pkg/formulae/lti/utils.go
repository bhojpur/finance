package lti

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
	"errors"

	"gonum.org/v1/gonum/mat"
)

//Discretize
// A_discretized = exp(A * t)
func discretize(m *mat.Dense, t float64) (*mat.Dense, error) {
	// m_d = exp(m * t)

	// check if matrix m is square
	if r, c := m.Dims(); r != c {
		return m, errors.New("Discretize: matrix is not square")
	}

	// tmp = m * t
	var tmp mat.Dense
	tmp.Scale(t, m)

	// exp( tmp )
	var md mat.Dense
	md.Exp(&tmp)

	return &md, nil
}

// Integrate
// B_discretized = Int_0^T exp(A t) B dt
// with exp(A t) = Sum_k (A t)^k / k!
func integrate(a *mat.Dense, b *mat.Dense, t float64) (*mat.Dense, error) {

	// B_d = Int_0^T exp(At) * B dt
	// B_d = (Int_0^T exp(At) dt) * B
	// B_d = ( T [ Sum_n (AT)^n-1 ] ) * B

	// (At)
	var at mat.Dense
	at.Scale(t, a)

	// Sum_n (AT)^n-1 / n!
	var x, tmp mat.Dense
	x.CloneFrom(&at)
	x.Zero()
	fac := 1.0
	for n := 1; n < 10; n++ {
		// (AT)^n-1
		tmp.Pow(&at, n-1)
		// n!
		fac = fac * float64(n)
		tmp.Scale(1.0/fac, &tmp)
		//fmt.Println("n=", n, "fac=", fac, "tmp=", tmp)
		x.Add(&tmp, &x)
	}
	//fmt.Println("at=", at)
	//fmt.Println("x=", x)

	// Int * B
	var bd mat.Dense
	bd.Mul(&x, b)

	bd.Scale(t, &bd)

	return &bd, nil
}

// rank calculates rank of matrix using singular value decomposition
func rank(a *mat.Dense) (int, error) {
	var svd mat.SVD
	ok := svd.Factorize(a, mat.SVDNone)
	if !ok {
		return 0, errors.New("rank: factorization failed")
	}
	rank := 0
	for _, value := range svd.Values(nil) {
		if value > 1e-8 {
			rank++
		}
	}
	return rank, nil
}

//checkControllability checks controllability of the LTI system
func checkControllability(a *mat.Dense, b *mat.Dense) (bool, error) {
	// system is controllable if
	// rank( [B, A B, A^2 B, A^n-1 B] ) = n

	// controllability matrix
	n, _ := b.Dims()

	var c, ab mat.Dense
	c.CloneFrom(b)
	ab.CloneFrom(b)

	// create augmented matrix
	for i := 0; i < n-1; i++ {
		ab.Mul(a, &ab)
		var tmp mat.Dense
		tmp.Augment(&c, &ab)
		c.CloneFrom(&tmp)
	}
	//fmt.Println(c)

	// calculate rank
	rank, err := rank(&c)
	if err != nil {
		return false, err
	}
	//fmt.Println("rank(C)=", rank)

	// check
	if rank < n {
		return false, nil
	}
	return true, nil
}

//checkObservability checks observability of the LTI system
func checkObservability(a *mat.Dense, c *mat.Dense) (bool, error) {
	// system is observable if
	// rank( S=[C, C A, C A^2, ..., C A^n-1]' ) = n

	// observability matrix S
	_, n := c.Dims()

	var sb, ca mat.Dense
	sb.CloneFrom(c)
	ca.CloneFrom(c)

	// create stacked matrix
	for i := 0; i < n-1; i++ {
		ca.Mul(&ca, a)
		var tmp mat.Dense
		tmp.Stack(&sb, &ca)
		sb.CloneFrom(&tmp)
	}
	//fmt.Println("S=", s)

	// calculate rank
	rank, err := rank(&sb)
	if err != nil {
		return false, err
	}
	//fmt.Println("rank(S)=", rank)

	// check
	if rank < n {
		return false, nil
	}
	return true, nil
}

// multAndSumOp multiplies A * x and B * u and returns the sum
func multAndSumOp(a *mat.Dense, x *mat.VecDense, b *mat.Dense, u *mat.VecDense) *mat.VecDense {

	// ax = A * x
	var ax mat.VecDense
	ax.MulVec(a, x)

	// bu = B * u
	var bu mat.VecDense
	bu.MulVec(b, u)

	// sum = A * x + B * u
	var sum mat.VecDense
	sum.AddVec(&ax, &bu)

	return &sum
}
