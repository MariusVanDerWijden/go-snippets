// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.8.0;

/************************************************************************************************
Originally from https://github.com/balancer-labs/balancer-core/blob/master/contracts/BNum.sol

This source code has been modified from the original, which was copied from the github repository
at commit hash f4ed5d65362a8d6cec21662fb6eae233b0babc1f.

Subject to the GPL-3.0 license
*************************************************************************************************/

contract BConst {
  uint256 public constant VERSION_NUMBER = 1;

/* ---  Weight Updates  --- */

  // Minimum time passed between each weight update for a token.
  uint256 public constant WEIGHT_UPDATE_DELAY = 1 hours;

  // Maximum percent by which a weight can adjust at a time
  // relative to the current weight.
  // The number of iterations needed to move from weight A to weight B is the floor of:
  // (A > B): (ln(A) - ln(B)) / ln(1.01)
  // (B > A): (ln(A) - ln(B)) / ln(0.99)
  uint256 public constant WEIGHT_CHANGE_PCT = BONE/100;

  uint256 public constant BONE = 10**18;

  uint256 public constant MIN_BOUND_TOKENS = 2;
  uint256 public constant MAX_BOUND_TOKENS = 10;

  // Minimum swap fee.
  uint256 public constant MIN_FEE = BONE / 10**6;
  // Maximum swap or exit fee.
  uint256 public constant MAX_FEE = BONE / 10;
  // Actual exit fee.
  uint256 public constant EXIT_FEE = 0;

  // Default total of all desired weights. Can differ by up to BONE.
  uint256 public constant DEFAULT_TOTAL_WEIGHT = BONE * 25;
  // Minimum weight for any token (1/100).
  uint256 public constant MIN_WEIGHT = BONE / 4;
  uint256 public constant MAX_WEIGHT = BONE * 25;
  // Maximum total weight.
  uint256 public constant MAX_TOTAL_WEIGHT = BONE * 26;
  // Minimum balance for a token (only applied at initialization)
  uint256 public constant MIN_BALANCE = BONE / 10**12;
  // Initial pool tokens
  uint256 public constant INIT_POOL_SUPPLY = BONE * 100;

  uint256 public constant MIN_BPOW_BASE = 1 wei;
  uint256 public constant MAX_BPOW_BASE = (2 * BONE) - 1 wei;
  uint256 public constant BPOW_PRECISION = BONE / 10**10;

  // Maximum ratio of input tokens to balance for swaps.
  uint256 public constant MAX_IN_RATIO = BONE / 2;
  // Maximum ratio of output tokens to balance for swaps.
  uint256 public constant MAX_OUT_RATIO = (BONE / 3) + 1 wei;
}

contract BNum is BConst {
  function btoi(uint256 a) public pure returns (uint256) {
    return a / BONE;
  }

  function bfloor(uint256 a) public pure returns (uint256) {
    return btoi(a) * BONE;
  }

  function badd(uint256 a, uint256 b) public pure returns (uint256) {
    uint256 c = a + b;
    require(c >= a, "ERR_ADD_OVERFLOW");
    return c;
  }

  function bsub(uint256 a, uint256 b) public pure returns (uint256) {
    (uint256 c, bool flag) = bsubSign(a, b);
    require(!flag, "ERR_SUB_UNDERFLOW");
    return c;
  }

  function bsubSign(uint256 a, uint256 b)
    public
    pure
    returns (uint256, bool)
  {
    if (a >= b) {
      return (a - b, false);
    } else {
      return (b - a, true);
    }
  }

  function bmul(uint256 a, uint256 b) public pure returns (uint256) {
    uint256 c0 = a * b;
    require(a == 0 || c0 / a == b, "ERR_MUL_OVERFLOW");
    uint256 c1 = c0 + (BONE / 2);
    require(c1 >= c0, "ERR_MUL_OVERFLOW");
    uint256 c2 = c1 / BONE;
    return c2;
  }

  function bdiv(uint256 a, uint256 b) public pure returns (uint256) {
    require(b != 0, "ERR_DIV_ZERO");
    uint256 c0 = a * BONE;
    require(a == 0 || c0 / a == BONE, "ERR_DIV_INTERNAL"); // bmul overflow
    uint256 c1 = c0 + (b / 2);
    require(c1 >= c0, "ERR_DIV_INTERNAL"); //  badd require
    uint256 c2 = c1 / b;
    return c2;
  }

  // DSMath.wpow
  function bpowi(uint256 a, uint256 n) public pure returns (uint256) {
    uint256 z = n % 2 != 0 ? a : BONE;

    for (n /= 2; n != 0; n /= 2) {
      a = bmul(a, a);

      if (n % 2 != 0) {
        z = bmul(z, a);
      }
    }
    return z;
  }

  // Compute b^(e.w) by splitting it into (b^e)*(b^0.w).
  // Use `bpowi` for `b^e` and `bpowK` for k iterations
  // of approximation of b^0.w
  function bpow(uint256 base, uint256 exp) public pure returns (uint256) {
    require(base >= MIN_BPOW_BASE, "ERR_BPOW_BASE_TOO_LOW");
    require(base <= MAX_BPOW_BASE, "ERR_BPOW_BASE_TOO_HIGH");

    uint256 whole = bfloor(exp);
    uint256 remain = bsub(exp, whole);

    uint256 wholePow = bpowi(base, btoi(whole));

    if (remain == 0) {
      return wholePow;
    }

    uint256 partialResult = bpowApprox(base, remain, BPOW_PRECISION);
    return bmul(wholePow, partialResult);
  }

  function bpowApprox(
    uint256 base,
    uint256 exp,
    uint256 precision
  ) public pure returns (uint256) {
    // term 0:
    uint256 a = exp;
    (uint256 x, bool xneg) = bsubSign(base, BONE);
    uint256 term = BONE;
    uint256 sum = term;
    bool negative = false;

    // term(k) = numer / denom
    //         = (product(a - i - 1, i=1-->k) * x^k) / (k!)
    // each iteration, multiply previous term by (a-(k-1)) * x / k
    // continue until term is less than precision
    for (uint256 i = 1; term >= precision; i++) {
      uint256 bigK = i * BONE;
      (uint256 c, bool cneg) = bsubSign(a, bsub(bigK, BONE));
      term = bmul(term, bmul(c, x));
      term = bdiv(term, bigK);
      if (term == 0) break;

      if (xneg) negative = !negative;
      if (cneg) negative = !negative;
      if (negative) {
        sum = bsub(sum, term);
      } else {
        sum = badd(sum, term);
      }
    }

    return sum;
  }
}