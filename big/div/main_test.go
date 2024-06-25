package big

import (
	"math/big"
	"testing"

	"fmt"
)

func TestMain(t *testing.T) {
	u := NewInt(0x0c)
	u.Lsh(u, 8*1336)
	v := NewInt(0x08)
	v.Lsh(v, 8*448)
	v.Add(v, NewInt(0x0f))
	v.Lsh(v, 8*440)
	z := new(Int).Div(u, v)
	uB := new(big.Int).SetBytes(u.Bytes())
	vB := new(big.Int).SetBytes(v.Bytes())
	zB := new(big.Int).SetBytes(z.Bytes())
	fmt.Println("u / v = z")
	fmt.Printf("with u = %s\n", uB)
	fmt.Printf("with v = %s\n", vB)
	fmt.Printf("with z = %s\n", zB)
	t.Fail()
}
