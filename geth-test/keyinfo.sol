/**
 *Submitted for verification at Etherscan.io on 2020-10-27
*/

// @title dKeys Key Info
// @author Atomrigs Lab
// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract KeyInfo {

  struct Shares {
    bytes party1;
    bytes party2;
    uint256 timestamp;    
  }
  
  mapping(bytes => Shares) public pubkeys;
  
  mapping(bytes32 => bytes[]) public userkeys;
  
  mapping(bytes => bytes) public pubkeyForwards;

  event KeyShareUpdated(bytes indexed pubkey, bytes32 indexed user, 
    bytes party1, bytes party2, uint256 timestamp);
    
  event PubkeyForwarded(bytes indexed pubkey, bytes indexed newPubkey, 
    bytes32 indexed user, uint256 timestamp);

  modifier onlyOwner(bytes memory pubkey) {
    //require(address(uint160(uint256(keccak256(pubkey)))) == msg.sender);
    _;
  } 

  function addKeyShareInfo(bytes32 user, bytes memory pubkey, bytes memory party1, bytes memory party2)
    public onlyOwner(pubkey) returns (bool success) {

    pubkeys[pubkey] = Shares({
      party1: party1,
      party2: party2,
      timestamp: block.timestamp
    });
    userkeys[user].push(pubkey);
    emit KeyShareUpdated(pubkey, user, party1, party2, block.timestamp);
    return true;
  }

  function forwardPubkey(bytes memory pubkey, bytes memory newPubkey, bytes32 user) 
    public onlyOwner(pubkey) returns (bool success) {

    pubkeyForwards[pubkey] = newPubkey;
    emit PubkeyForwarded(pubkey, newPubkey, user, block.timestamp);

    return true;    
  }
}