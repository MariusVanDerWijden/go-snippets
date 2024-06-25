pragma solidity >=0.7.0 <0.9.0;

contract WarmSStore {
    mapping(uint => uint) public values;
    uint public run = 0;

    function Sstore(uint runs) public {
        run += 1;
        for(uint i = 0; i < runs; i++) {
            values[i] = run;
        }
    }
}