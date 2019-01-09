// This package is used to implement Cipolla's algorithm, which solves
// x² ≡ n mod p
// We need this method in order to reconstruct y from x according to the curve at hand
package bn256

import (
	"math/big"
)

type point struct{ x, y big.Int }

func mul(a, b point, ω2, p big.Int) (z point) {
	var w big.Int
	z.x.Mod(z.x.Add(z.x.Mul(&a.x, &b.x), w.Mul(w.Mul(&a.y, &a.y), &ω2)), &p)
	z.y.Mod(z.y.Add(z.y.Mul(&a.x, &b.y), w.Mul(&b.x, &a.y)), &p)
	return
}

func cipolla(n, p big.Int) (R1, R2 big.Int, ok bool) {
	if big.Jacobi(&n, &p) != 1 {
		return
	}
	var one, a, ω2 big.Int
	one.SetInt64(1)
	for ; ; a.Add(&a, &one) {
		// big.Int Mod uses Euclidean division, result is always >= 0
		ω2.Mod(ω2.Sub(ω2.Mul(&a, &a), &n), &p)
		if big.Jacobi(&ω2, &p) == -1 {
			break
		}
	}

	var r, s point
	r.x.SetInt64(1)
	s.x.Set(&a)
	s.y.SetInt64(1)
	var e big.Int
	for e.Rsh(e.Add(&p, &one), 1); len(e.Bits()) > 0; e.Rsh(&e, 1) {
		if e.Bit(0) == 1 {
			r = mul(r, s, ω2, p)
		}
		s = mul(s, s, ω2, p)
	}
	R2.Sub(&p, &r.x)
	return r.x, R2, true
}
