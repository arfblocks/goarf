pragma solidity ^0.4.0;

import "./owned.sol";

contract validatorsF is owned {

    struct addressStatus {
	bool isOK;
	uint index;
    }

    address[] validators;
    address[] pending;
    mapping(address => addressStatus) status;

    function addValidator(address _validator)
	external
	onlyOwner
// isNotValidator(_validator)
    {
	status[_validator].isIn = true;
	status[_validator].index = pending.length;
	pending.push(_validator);
	triggerChange();
    }
}
