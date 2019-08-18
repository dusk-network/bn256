bn256
-----

Package bn256 implements a particular bilinear group at the 128-bit security
level. It is a modification of [Cloudflare's implementation of BN256](https://github.com/cloudflare/bn256) 
which in turn represents a substantial improvement of golang's official version at
https://golang.org/x/crypto/bn256 with significantly faster operations (~10 times faster on
amd64 and arm64).

Bilinear groups are the basis of many of the new cryptographic protocols that
have been proposed over the past decade. They consist of a triplet of groups
(G₁, G₂ and GT) such that there exists a function e(g₁ˣ,g₂ʸ)=gTˣʸ (where gₓ is a
generator of the respective group). That function is called a pairing function.

This package specifically implements the Optimal Ate pairing over a 256-bit
Barreto-Naehrig curve as described in
http://cryptojedi.org/papers/dclxvi-20100714.pdf. Its output is compatible with
the implementation described in that paper.

It also includes functionalities for point compression and decompression into 
a 33 byte representation. The additional byte includes information to resolve 
the ambiguity of calculating the Y coordinate from X coordinate using the curve 
equation

### Benchmarks

branch `master`:
```
BenchmarkG1-4        	   10000	    154995 ns/op
BenchmarkG2-4        	    3000	    541503 ns/op
BenchmarkGT-4        	    1000	   1267811 ns/op
BenchmarkPairing-4   	    1000	   1630584 ns/op
```

official version:
```
BenchmarkG1-4        	    1000	   2268491 ns/op
BenchmarkG2-4        	     300	   7227637 ns/op
BenchmarkGT-4        	     100	  15121359 ns/op
BenchmarkPairing-4   	      50	  20296164 ns/op
```

### Note
The original Clouflare's repository includes a `lattice` branch for non-commercial 
use which benchmarks ~10 times faster than the `official version`. Such branch has not 
been migrated to this repository yet.
