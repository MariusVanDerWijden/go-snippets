pragma solidity >=0.4.22 <0.7.0;


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