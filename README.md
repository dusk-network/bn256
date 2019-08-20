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

*Specs*

```
memory         15GiB System memory
processor      Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
bridge         Xeon E3-1200 v6/7th Gen Core Processor Host Bridge/DRAM Registers
```

*Lattices Branch*

```
BenchmarkG1-8              20000             77906 ns/op
BenchmarkG2-8               5000            269775 ns/op
BenchmarkGT-8               3000            508216 ns/op
BenchmarkPairing-8          1000           1299645 ns/op
```

*Master Branch*

```
BenchmarkG1-8              10000            121665 ns/op
BenchmarkG2-8               3000            435977 ns/op
BenchmarkGT-8               2000           1026991 ns/op
BenchmarkPairing-8          1000           1301513 ns/op
```

* official version (as reported by Cloudflare on non-specified hardware)*

```
BenchmarkG1-4        	    1000	   2268491 ns/op
BenchmarkG2-4        	     300	   7227637 ns/op
BenchmarkGT-4        	     100	  15121359 ns/op
BenchmarkPairing-4   	      50	  20296164 ns/op
```

### Note
The original Clouflare's repository includes a `lattice` branch for non-commercial 
use which benchmarks significantly faster than both the `official golang version` and the master branch. 
The `lattices` branch is for non-commercial use only.
