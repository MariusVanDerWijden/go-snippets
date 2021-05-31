
pragma solidity ^0.5.16;
contract reverter {
    function revert() public{
        require(false, "revert reason 123");
    }
}
