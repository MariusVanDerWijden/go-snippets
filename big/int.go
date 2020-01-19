// Copyright 2020 Marius van der Wijden. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This file implements a better interface for signed mult-precision integers.

package big

import (
	"big"
)

func Add(a, b big.Int) *big.Int {
	return a.Add(b)
}