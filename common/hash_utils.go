package common

import (
	big "github.com/binance-chain/tss-lib/common/int"
)

// SHA512/256
const hashBitLen = 256

// RejectionSample implements the rejection sampling logic for converting a
// SHA512/256 hash to a value between 0-q
func RejectionSample(q *big.Int, eHash *big.Int) *big.Int { // e' = eHash
	// e = the first |q| bits of e'
	qBits := q.BitLen()
	e := eHash
	// optimisation to skip firstBitsOf for secp256k1 and nist256p1
	if qBits != hashBitLen {
		e = firstBitsOf(qBits, eHash)
	}
	// while e is not between 0-q
	for e.Cmp(q) > -1 {
		eHash = SHA512_256iOne(eHash)
		e = firstBitsOf(qBits, eHash)
	}
	return e
}

func firstBitsOf(bits int, v *big.Int) *big.Int {
	e := big.NewInt(0)
	for i := 0; i < bits; i++ {
		bit := v.Bit(i)
		if 0 < bit {
			e.SetBit(e, i, bit)
		}
	}
	return e
}
