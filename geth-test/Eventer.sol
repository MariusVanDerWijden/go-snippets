pragma solidity >=0.6.0 <0.9.0;
pragma experimental ABIEncoderV2;

contract Eventer {
   
    event TestInt8(int8 indexed out1, int8 indexed out2);
    event AnonEvent(address, address);
    
    function getEvent() public {
        // set to 2,3 for functioning filter
        emit TestInt8(-2, -3);
    }

    function anonEvent() public {
        emit AnonEvent(msg.sender, msg.sender);
    }

    function fail() public {
        require(false, "error string");
    }
}

contract BigBoard {
    uint8 constant BOARD_WIDTH = 10;
	uint8 constant BOARD_HEIGHT = 9;

	uint8[9][10][3] board;

	function get_board_state() public view returns (uint8[9][10][3] memory) {
		return board;
	}

    function fill1() public {
        for(uint i = 0; i < 3; i++) {
            for(uint k = 0; k < 10; k++) {
                for(uint l = 0; l  < 9; l++) {
                    board[i][k][l] = uint8(i * (k + l));
                }
            }
        }
    }
}

contract Array {
    function get_array() public pure returns (uint8[2][4] memory) {
        uint8[2][4] memory board;
        board[0][0] = 0;
        board[1][0] = 1;
        board[2][0] = 2;
        board[3][0] = 3;
        board[0][1] = 4;
        board[1][1] = 5;
        board[2][1] = 6;
        board[3][1] = 7;
        return board;
    }
}

contract ReceiveFallback {
    event Fallback(address sender);
    event Receive(address sender);
    fallback() external {
        emit Fallback(msg.sender);
    }
    receive() external payable {
        emit Receive(msg.sender);
    }
}



contract TupleTest {
    struct S { uint a; uint[] b; T[] c; }
    event evF(S, T, uint);
    function f(S memory s, T memory t, uint u) public {
        emit evF(s,t,u);
    }
    function g() public pure returns (S memory, T memory, uint) {}
}


struct T { uint x; uint y; }
contract TupleTest2 {
    
    struct S { uint a; uint[] b; T[] c; }
    struct A {uint a; uint[]b;}
    event evF(S, T, uint);
    
    function a(T memory t) public view returns (T memory) {
        return t;
    }
    function b(T[] memory t) public view returns (T[] memory) {
        return t;
    }
    
    function method(A memory a, uint b) view public {}
    
    function f(S memory s, T memory t, uint u) view public {
        //emit evF(s,t,u);
    }
    
    function g() public view returns (S memory, T memory, uint) {
        T memory t = T(1, 2);
        uint[] memory b = new uint[](3);
        b[0] = 1;
        b[1] = 2;
        b[2] = 3;
        T[] memory ts = new T[](2);
        ts[0] = t;
        ts[1] = t;
        S memory a = S(1, b, ts);
       // emit evF(a, t, 3);
        return (a, t, 3);
    }
}


contract VoidTest {
     function method(uint a) view public {}
}
