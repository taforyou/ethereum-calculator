pragma solidity >=0.4.24;

contract Calculator {
  modifier divideByZero(uint y) {
      require(y > 0);
      _;
  }
  function plus(uint8 x, uint y) public view returns (uint) {
      return x + y;
  }
  function minus(int x, int y) public view returns (int) {
      return  x - y;
  }
  function multiply(uint x, uint y) public view returns (uint) {
      return x * y;
  }
  function divide(uint x, uint y) public divideByZero(y) view returns (uint) {
      return x/y;
  }
}
