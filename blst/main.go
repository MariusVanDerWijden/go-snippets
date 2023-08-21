package test

import (
	"crypto/rand"
	"fmt"

	blst "github.com/supranational/blst/bindings/go"
)

type PublicKey = blst.P1Affine
type Signature = blst.P2Affine
type AggregateSignature = blst.P2Aggregate
type AggregatePublicKey = blst.P1Aggregate

func main() {
	Fuzz([]byte{})
}

func Fuzz(data []byte) int {
	var ikm [32]byte
	_, _ = rand.Read(ikm[:])
	sk := blst.KeyGen(ikm[:])
	pk := new(PublicKey).From(sk)

	var dst = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")
	msg := []byte("hello foo")
	sig := new(Signature).Sign(sk, msg, dst)

	if !sig.Verify(true, pk, true, msg, dst) {
		fmt.Println("ERROR: Invalid!")
	} else {
		fmt.Println("Valid!")
	}
	return 0
}
