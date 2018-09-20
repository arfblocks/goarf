pragma solidity ^0.4.0;

contract owned {
    address public owner = msg.sender;
    event newOwner(address indexed old, address indexed current);
    
    modifier onlyOwner() {
        if (msg.sender == owner) {
            _;
        }
    }

    function setOwner(address _new)
	external
	onlyOwner
    {
	emit newOwner(owner, _new);
	owner = _new;
    }
}
