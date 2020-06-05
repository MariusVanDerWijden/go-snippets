pragma solidity >=0.4.22 <0.7.0;

contract CallBLS {
    
    event Bytes(bytes);
    event Uint(uint256);
    
    function callPrec(uint256 addr, bytes memory input) public {
        uint256 a;
        uint256 b;
        uint256 c;
        uint256 d;
        uint256 len = input.length;
        uint256 outLen;
        if (addr == 0x0a || addr == 0x0b || addr == 0x0c || addr == 0x11) {
            outLen = 128;
        } else if (addr == 0x0d || addr == 0x0e || addr == 0x0f || addr == 0x12) {
            outLen = 256;
        } else if (addr == 0x10) {
            outLen = 32;
        } else {
            require(false, "Invalid precompile");
        }
        assembly {
            let res := mload(0x40)
            // call precompile
            if iszero(staticcall(gas(), addr, input, len, res, outLen)) {
                revert(0, 0)
            }
            a := mload(res)
            b := mload(add(res,32))
            c := mload(add(res,64))
            d := mload(add(res,96))
        }
        emit Uint(a);
        emit Uint(b);
        emit Uint(c);
        emit Uint(d);
    }
}
//"0x11","0x000000000000000000000000000000000dbb997ef4970a472bfcf03e959acb90bb13671a3d27c91698975a407856505e93837f46afc965363f21c35a3d194ec0"
// 0x51eF031Ba1B2128971bd67A3ba5dda9BfF8706c8