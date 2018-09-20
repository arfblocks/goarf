pragma solidity ^0.4.0;

interface validators {

    function getValidators()
	external
	view
	returns (address[]);
}
