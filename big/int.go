// Copyright 2020 Marius van der Wijden. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This file implements a better interface for signed mult-precision integers.

package big

import (
	"math/big"
	"math/rand"
)

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
//
func Sign(x *big.Int) int {
	return x.Sign()
}

// SetInt64 sets z to x and returns z.
func SetInt64(z *big.Int, x int64) *big.Int {
	return z.SetInt64(x)
}

// SetUint64 sets z to x and returns z.
func SetUint64(z *big.Int, x uint64) *big.Int {
	return z.SetUint64(x)
}

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *big.Int {
	return big.NewInt(x)
}

// Set sets z to x and returns z.
func Set(z, x *big.Int) *big.Int {
	return z.Set(x)
}

// Bits provides raw (unchecked but fast) access to x by returning its
// absolute value as a little-endian Word slice. The result and x share
// the same underlying array.
// Bits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func Bits(x *big.Int) []big.Word {
	return x.Bits()
}

// SetBits provides raw (unchecked but fast) access to z by setting its
// value to abs, interpreted as a little-endian Word slice, and returning
// z. The result and abs share the same underlying array.
// SetBits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func SetBits(z *big.Int, abs []big.Word) *big.Int {
	return z.SetBits(abs)
}

// Abs sets z to |x| (the absolute value of x) and returns z.
func Abs(z, x *big.Int) *big.Int {
	return z.Abs(x)
}

// Neg sets z to -x and returns z.
func Neg(z, x *big.Int) *big.Int {
	return z.Neg(x)
}

// Add returns the sum of x+y.
func Add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

// Sub returns the difference of x-y.
func Sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

// Mul returns the product x*y.
func Mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

// MulRange returns the product of all integers
// in the range [a, b] inclusively.
// If a > b (empty range), the result is 1.
func MulRange(a, b int64) *big.Int {
	return new(big.Int).MulRange(a, b)
}

// Binomial returns the binomial coefficient of (n, k).
func Binomial(n, k int64) *big.Int {
	return new(big.Int).Binomial(n, k)
}

// Quo returns the quotient x/y for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
// Quo implements truncated division (like Go); see QuoRem for more details.
func Quo(x, y *big.Int) *big.Int {
	return new(big.Int).Quo(x, y)
}

// Rem returns the remainder x%y for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
// Rem implements truncated modulus (like Go); see QuoRem for more details.
func Rem(x, y *big.Int) *big.Int {
	return new(big.Int).Rem(x, y)
}

// QuoRem returns the quotient x/y and the remainder x%y
// and returns the pair (z, r) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// QuoRem implements T-division and modulus (like Go):
//
//	q = x/y      with the result truncated to zero
//	r = x - y*q
//
// (See Daan Leijen, ``Division and Modulus for Computer Scientists''.)
// See DivMod for Euclidean division and modulus (unlike Go).
//
func QuoRem(x, y, r *big.Int) (*big.Int, *big.Int) {
	return new(big.Int).QuoRem(x, y, r)
}

// Div returns the quotient x/y for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see DivMod for more details.
func Div(x, y *big.Int) *big.Int {
	return new(big.Int).Div(x, y)
}

// Mod sets z to the modulus x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean modulus (unlike Go); see DivMod for more details.
func Mod(x, y *big.Int) *big.Int {
	return new(big.Int).Mod(x, y)
}

// DivMod sets z to the quotient x div y and m to the modulus x mod y
// and returns the pair (z, m) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// DivMod implements Euclidean division and modulus (unlike Go):
//
//	q = x div y  such that
//	m = x - y*q  with 0 <= m < |y|
//
// (See Raymond T. Boute, ``The Euclidean definition of the functions
// div and mod''. ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.)
// See QuoRem for T-division and modulus (like Go).
//
func DivMod(x, y, m *big.Int) (*big.Int, *big.Int) {
	return new(big.Int).DivMod(x, y, m)
}

// Cmp compares x and y and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func Cmp(x, y *big.Int) (r int) {
	return x.Cmp(y)
}

// CmpAbs compares the absolute values of x and y and returns:
//
//   -1 if |x| <  |y|
//    0 if |x| == |y|
//   +1 if |x| >  |y|
//
func CmpAbs(x, y *big.Int) int {
	return x.CmpAbs(y)
}

// Int64 returns the int64 representation of x.
// If x cannot be represented in an int64, the result is undefined.
func Int64(x *big.Int) int64 {
	return x.Int64()
}

// Uint64 returns the uint64 representation of x.
// If x cannot be represented in a uint64, the result is undefined.
func Uint64(x *big.Int) uint64 {
	return x.Uint64()
}

// IsInt64 reports whether x can be represented as an int64.
func IsInt64(x *big.Int) bool {
	return x.IsInt64()
}

// IsUint64 reports whether x can be represented as a uint64.
func IsUint64(x *big.Int) bool {
	return x.IsUint64()
}

// SetString sets z to the value of s, interpreted in the given base,
// and returns z and a boolean indicating success. The entire string
// (not just a prefix) must be valid for success. If SetString fails,
// the value of z is undefined but the returned value is nil.
//
// The base argument must be 0 or a value between 2 and MaxBase.
// For base 0, the number prefix determines the actual base: A prefix of
// ``0b'' or ``0B'' selects base 2, ``0'', ``0o'' or ``0O'' selects base 8,
// and ``0x'' or ``0X'' selects base 16. Otherwise, the selected base is 10
// and no prefix is accepted.
//
// For bases <= 36, lower and upper case letters are considered the same:
// The letters 'a' to 'z' and 'A' to 'Z' represent digit values 10 to 35.
// For bases > 36, the upper case letters 'A' to 'Z' represent the digit
// values 36 to 61.
//
// For base 0, an underscore character ``_'' may appear between a base
// prefix and an adjacent digit, and between successive digits; such
// underscores do not change the value of the number.
// Incorrect placement of underscores is reported as an error if there
// are no other errors. If base != 0, underscores are not recognized
// and act like any other character that is not a valid digit.
//
func SetString(z *big.Int, s string, base int) (*big.Int, bool) {
	return z.SetString(s, base)
}

// SetBytes interprets buf as the bytes of a big-endian unsigned
// integer, sets z to that value, and returns z.
func SetBytes(z *big.Int, buf []byte) *big.Int {
	return z.SetBytes(buf)
}

// Bytes returns the absolute value of x as a big-endian byte slice.
func Bytes(z *big.Int) []byte {
	return z.Bytes()
}

// BitLen returns the length of the absolute value of x in bits.
// The bit length of 0 is 0.
func BitLen(z *big.Int) int {
	return z.BitLen()
}

// TrailingZeroBits returns the number of consecutive least significant zero
// bits of |x|.
func TrailingZeroBits(x *big.Int) uint {
	return x.TrailingZeroBits()
}

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
// If m == nil or m == 0, z = x**y unless y <= 0 then z = 1. If m > 0, y < 0,
// and x and n are not relatively prime, z is unchanged and nil is returned.
//
// Modular exponentation of inputs of a particular size is not a
// cryptographically constant-time operation.
func Exp(x, y, m *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, m)
}

// GCD sets z to the greatest common divisor of a and b, which both must
// be > 0, and returns z.
// If x or y are not nil, GCD sets their value such that z = a*x + b*y.
// If either a or b is <= 0, GCD sets z = x = y = 0.
func GCD(x, y, a, b *big.Int) *big.Int {
	return new(big.Int).GCD(x, y, a, b)
}

// Rand sets z to a pseudo-random number in [0, n) and returns z.
//
// As this uses the math/rand package, it must not be used for
// security-sensitive work. Use crypto/rand.Int instead.
func Rand(rnd *rand.Rand, n *big.Int) *big.Int {
	return new(big.Int).Rand(rnd, n)
}

// ModInverse sets z to the multiplicative inverse of g in the ring ℤ/nℤ
// and returns z. If g and n are not relatively prime, g has no multiplicative
// inverse in the ring ℤ/nℤ.  In this case, z is unchanged and the return value
// is nil.
func ModInverse(g, n *big.Int) *big.Int {
	return new(big.Int).ModInverse(g, n)
}

// Jacobi returns the Jacobi symbol (x/y), either +1, -1, or 0.
// The y argument must be an odd integer.
func Jacobi(x, y *big.Int) int {
	return big.Jacobi(x, y)
}

// ModSqrt sets z to a square root of x mod p if such a square root exists, and
// returns z. The modulus p must be an odd prime. If x is not a square mod p,
// ModSqrt leaves z unchanged and returns nil. This function panics if p is
// not an odd integer.
func ModSqrt(x, p *big.Int) *big.Int {
	return new(big.Int).ModSqrt(x, p)
}

// Lsh sets z = x << n and returns z.
func Lsh(x *big.Int, n uint) *big.Int {
	return new(big.Int).Lsh(x, n)
}

// Rsh sets z = x >> n and returns z.
func Rsh(x *big.Int, n uint) *big.Int {
	return new(big.Int).Rsh(x, n)
}

// Bit returns the value of the i'th bit of x. That is, it
// returns (x>>i)&1. The bit index i must be >= 0.
func Bit(x *big.Int, i int) uint {
	return x.Bit(i)
}

// SetBit sets z to x, with x's i'th bit set to b (0 or 1).
// That is, if b is 1 SetBit sets z = x | (1 << i);
// if b is 0 SetBit sets z = x &^ (1 << i). If b is not 0 or 1,
// SetBit will panic.
func SetBit(x *big.Int, i int, b uint) *big.Int {
	return new(big.Int).SetBit(x, i, b)
}

// And sets z = x & y and returns z.
func And(x, y *big.Int) *big.Int {
	return new(big.Int).And(x, y)
}

// AndNot sets z = x &^ y and returns z.
func AndNot(x, y *big.Int) *big.Int {
	return new(big.Int).AndNot(x, y)
}

// Or sets z = x | y and returns z.
func Or(x, y *big.Int) *big.Int {
	return new(big.Int).Or(x, y)
}

// Xor sets z = x ^ y and returns z.
func Xor(x, y *big.Int) *big.Int {
	return new(big.Int).Xor(x, y)
}

// Not sets z = ^x and returns z.
func Not(x *big.Int) *big.Int {
	return new(big.Int).Not(x)
}

// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
// It panics if x is negative.
func Sqrt(x *big.Int) *big.Int {
	return new(big.Int).Sqrt(x)
}
