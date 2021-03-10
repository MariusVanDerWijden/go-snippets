contract IO{
    
    // write writes 'value' to 'slot'
    function write( bytes32 slot,  bytes32 value) public payable{
        assembly{
            sstore(slot, value)
        }
    }

    // write writes 'value' to 'slot', then errors out
    function writeRevert( bytes32 slot,  bytes32 value) public payable{
        assembly{
            sstore(slot, value)
            revert(0,0)
        }
    }

    // read reads and returns the value at 'slot'
    function read(bytes32 slot) public payable returns(bytes32){
        assembly{
            mstore(0, sload(slot))
            return (0,32)
        }
    }

    // read reads and returns the value at 'slot', then errors out
    function readRevert(bytes32 slot) public payable returns(bytes32){
        assembly{
            mstore(0, sload(slot))
            revert(0,0)
        }
    }
}

contract DcallProxy {
    address target;

    function setTarget(address newTarget) public {
        target = newTarget;
    }

    fallback() external payable {
        assembly {
            let _target := sload(0)
            calldatacopy(0x0, 0x0, calldatasize())
            let result := delegatecall(gas(), _target, 0x0, calldatasize(), 0x0, 0)
            returndatacopy(0x0, 0x0, returndatasize())
            return (0, returndatasize())
        }
    }
}

contract CallProxy {
    address target;

    function setTarget(address newTarget) public {
        target = newTarget;
    }

    fallback() external payable {
        assembly {
            let _target := sload(0)
            calldatacopy(0x0, 0x0, calldatasize())
            let result := call(gas(), _target, callvalue(), 0x0, calldatasize(), 0x0, 0)
            returndatacopy(0x0, 0x0, returndatasize())
            return (0, returndatasize())
        }
    }
}

contract CcallProxy {
    address target;

    function setTarget(address newTarget) public {
        target = newTarget;
    }

    fallback() external payable {
        assembly {
            let _target := sload(0)
            calldatacopy(0x0, 0x0, calldatasize())
            let result := callcode(gas(), _target, callvalue(), 0x0, calldatasize(), 0x0, 0)
            returndatacopy(0x0, 0x0, returndatasize())
            return (0, returndatasize())
        }
    }
}

contract ScallProxy {
    address target;

    function setTarget(address newTarget) public {
        target = newTarget;
    }

    fallback() external payable {
        assembly {
            let _target := sload(0)
            calldatacopy(0x0, 0x0, calldatasize())
            let result := staticcall(gas(), _target, 0x0, calldatasize(), 0x0, 0)
            returndatacopy(0x0, 0x0, returndatasize())
            return (0, returndatasize())
        }
    }
}

contract Deployer{
    function deploy() public returns (address[5] memory){
       address[5] memory ret ;
        ret[0] = address(new IO());
        ret[1] = address( new DcallProxy());
        ret[2] = address(new CallProxy());
        ret[3] = address(new CcallProxy());
        ret[4] = address(new ScallProxy());
        return ret;
    }
}