// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

/// @title contract for anonymous vote
/// @author 0xlahh
/// @notice vote anonymously
/// @dev nothing

contract Vote {
    struct Candidate {
        uint id;
        string name;
        uint voteCount;
    }
    mapping(uint => Candidate) public candidates;
    mapping(address => bool) public hasVoted;
    uint public candidateCount;

    function addCandidate(string memory _name) private {
        candidateCount++;
        candidates[candidateCount] = Candidate(candidateCount, _name, 0);
    }

    function vote(uint _candidateID) public {
        require(!hasVoted[msg.sender], "You already voted before !");
        require(
            _candidateID > 0 && _candidateID <= candidateCount,
            "Candidate doesn't exist"
        );
        hasVoted[msg.sender] = true;
        candidates[_candidateID].voteCount++;
    }

    constructor() {
        addCandidate("lahh");
        addCandidate("another_lahh");
    }
}
