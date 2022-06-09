pragma solidity >=0.7.0 <0.9.0;

contract MySmartContract {
    function Hello() public view returns (string memory) {
        return "Hello, World!";
    }
    function Greet(string memory str) public view returns (string memory) {
        return str;
    }
}