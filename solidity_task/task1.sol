// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting{

  mapping(string =>uint256) public  votMap;
  string[] public names;

function vote( string calldata name, uint256   acount) public {
    // votMap[name]=acount;
    bool flag=true;
    for (uint i=0; i<names.length; i++) 
    {         
      if (keccak256(bytes(name)) == keccak256(bytes(names[i]))){
        flag=false;
        votMap[name]=acount+votMap[name];
        break ;
      }
    }
    if(flag){
    names.push(name);
    votMap[name]=acount;
    }
    
}

function getVotes(string calldata name)public view returns (uint256){
    return votMap[name];
}
 
function resetVotes()public  {
    for (uint i=0;i<names.length;i++){
        votMap[names[i]]=0;
    }
}

}