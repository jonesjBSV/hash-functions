package logicalFunctions

import "math/bits"

// ∑0(x): (ROTR 2) XOR (ROTR 13) XOR (ROTR 22)
func BigSigmaZero(input uint32) uint32 {

	return bits.RotateLeft32(input, -2) ^ bits.RotateLeft32(input, -13) ^ bits.RotateLeft32(input, -22)
}

// ∑1(x): (ROTR 6) XOR (ROTR 11) XOR (ROTR 25)
func BigSigmaOne(input uint32) uint32 {

	return bits.RotateLeft32(input, -6) ^ bits.RotateLeft32(input, -11) ^ bits.RotateLeft32(input, -25)
}

// σ0(x): (ROTR 7) XOR (ROTR 18) XOR (SHR 3)
func SmallSigmaZero(input uint32) uint32 {

	return bits.RotateLeft32(input, -7) ^ bits.RotateLeft32(input, -18) ^ input>>3
}

// σ1(x): (ROTR 17) XOR (ROTR 19) XOR (SHR 10)
func SmallSigmaOne(input uint32) uint32 {

	return bits.RotateLeft32(input, -17) ^ bits.RotateLeft32(input, -19) ^ input>>10
}

// The Choice Function:
// if(x) then y else z
func Choice(x, y, z uint32) uint32 {

	return (x & y) ^ (^x & z)
}

//The Majority Function:
//if(x && y || x && z || y && z) then 1 else 0*/
func Majority(x, y, z uint32) uint32 {

	return (x & y) ^ (x & z) ^ (y & z)
}
