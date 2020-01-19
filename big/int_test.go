// Copyright 2020 Marius van der Wijden. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This file implements a better interface for signed mult-precision integers.

package big

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	a := big.NewInt(1234)
	b := big.NewInt(4321)
	c := big.NewInt(5555)
	res := Add(a, b)
	assert.Equal(t, c, res, "add should be equal")
}
