package logicalFunctions

import "testing"

import "math/bits"

//import "fmt"

// σ0(x): (ROTR 7) XOR (ROTR 18) XOR (SHR 3)
// input: 01010001 00001110 01010010 01111111
func TestSmallSigmaZero(t *testing.T) {

	input := uint32(0x510e527f)

	// input:    01010001 00001110 01010010 01111111
	// expected: 11111110 10100010 00011100 10100100
	expectedX := uint32(0xFEA21CA4)
	x := bits.RotateLeft32(input, -7)

	if x != expectedX {
		t.Errorf("\nexpectedX value:        %32b\nx:                      %32b\n", expectedX, x)
	}

	// input:    01010001 00001110 01010010 01111111
	// expected: 10010100 10011111 11010100 01000011
	expectedY := uint32(0x949FD443)
	y := bits.RotateLeft32(input, -18)

	if y != expectedY {
		t.Errorf("\nexpectedY value:        %32b\ny:                      %32b\n", expectedY, y)
	}

	// input:    01010001 00001110 01010010 01111111
	// expected: 00001010 00100001 11001010 01001111
	expectedZ := uint32(0xA21CA4F)
	z := input >> 3

	if z != expectedZ {
		t.Errorf("\nexpectedZ value:        %32b\nz:                      %32b\n", expectedZ, z)
	}

	result := x ^ y ^ z

	expected := expectedX ^ expectedY ^ expectedZ

	if result != expected {
		t.Errorf("\nexpected value:         %32b\nresult:                 %32b\n", expected, result)
	}

}

// σ1(x): (ROTR 17) XOR (ROTR 19) XOR (SHR 10)
// input: 01010001 00001110 01010010 01111111
func TestSmallSigmaOne(t *testing.T) {

	input := uint32(0x510e527f)

	expectedX := uint32(0x293FA887)
	x := bits.RotateLeft32(input, -17)

	// input:    01010001 00001110 01010010 01111111
	// expected: 00101001 00111111 10101000 10000111
	if x != expectedX {
		t.Errorf("\nexpectedX value:        %32b\nx:                      %32b\n", expectedX, x)
	}

	expectedY := uint32(0xCA4FEA21)
	y := bits.RotateLeft32(input, -19)

	// input:    01010001 00001110 01010010 01111111
	// expected: 11001010 01001111 11101010 00100001
	if y != expectedY {
		t.Errorf("\nexpectedY value:        %32b\ny:                      %32b\n", expectedY, y)
	}

	expectedZ := uint32(0x144394)
	z := input >> 10

	// input:    01010001 00001110 01010010 01111111
	// expected: 00000000 00010100 01000011 10010100
	if z != expectedZ {
		t.Errorf("\nexpectedZ value:        %32b\nz:                      %32b\n", expectedZ, z)
	}

	result := x ^ y ^ z

	expected := expectedX ^ expectedY ^ expectedZ

	if result != expected {
		t.Errorf("\nexpected value:         %32b\nresult:                 %32b\n", expected, result)
	}

}

// ∑0(x): (ROTR 2) XOR (ROTR 13) XOR (ROTR 22)
// input: 01010001 00001110 01010010 01111111
func TestBigSigmaZero(t *testing.T) {

	input := uint32(0x510e527f)

	expectedX := uint32(0xD443949F)
	x := bits.RotateLeft32(input, -2)

	// input:    01010001 00001110 01010010 01111111
	// expected: 11010100 01000011 10010100 10011111
	if x != expectedX {
		t.Errorf("\nexpectedX value:        %32b\nx:                      %32b\n", expectedX, x)
	}

	expectedY := uint32(0x93FA8872)
	y := bits.RotateLeft32(input, -13)

	// input:    01010001 00001110 01010010 01111111
	// expected: 10010011 11111010 10001000 01110010
	if y != expectedY {
		t.Errorf("\nexpectedY value:        %32b\ny:                      %32b\n", expectedY, y)
	}

	expectedZ := uint32(0x3949FD44)
	z := bits.RotateLeft32(input, -22)

	// input:    01010001 00001110 01010010 01111111
	// expected: 00111001 01001001 11111101 01000100
	if z != expectedZ {
		t.Errorf("\nexpectedZ value:        %32b\nz:                      %32b\n", expectedZ, z)
	}

	expected := expectedX ^ expectedY ^ expectedZ

	result := x ^ y ^ z

	if result != expected {
		t.Errorf("\nexpected value:         %32b\nresult:                 %32b\n", expected, result)
	}
}

// ∑1(x): (ROTR 6) XOR (ROTR 11) XOR (ROTR 25)
// input: 01010001 00001110 01010010 01111111
func TestBigSigmaOne(t *testing.T) {

	//input 01100001 01100010 01100011
	input := uint32('a' << 16)
	input += ('b' << 8)
	input += 'c'

	//expected:
	expectedX := input
	x := bits.RotateLeft32(input, -6)

	// input:    01010001 00001110 01010010 01111111
	// expected: 11111101 01000100 00111001 01001001
	if x != expectedX {
		t.Errorf("\nexpectedX value:        %32b\nx:                      %32b\n", expectedX, x)
	}

	expectedY := uint32(0x4FEA21CA)
	y := bits.RotateLeft32(input, -11)

	// input:    01010001 00001110 01010010 01111111
	// expected: 01001111 11101010 00100001 11001010
	if y != expectedY {
		t.Errorf("\nexpectedY value:        %32b\ny:                      %32b\n", expectedY, y)
	}

	// input:    01010001 00001110 01010010 01111111
	// expected: 10000111 00101001 00111111 10101000
	expectedZ := uint32(0x87293FA8)
	z := bits.RotateLeft32(input, -25)

	if z != expectedZ {
		t.Errorf("\nexpectedZ value:        %32b\nz:                      %32b\n", expectedZ, z)
	}

	expected := expectedX ^ expectedY ^ expectedZ

	result := x ^ y ^ z

	if result != expected {
		t.Errorf("\nexpected value:         %32b\nresult:                 %32b\n", expected, result)
	}
}

// The Choice Function:
// if(x) then y else z
func TestChoice(t *testing.T) {

	// expected: 01100001
	x := uint32('a')

	// expected: 01100010
	y := uint32('b')

	// expected: 01100011
	z := uint32('c')

	result := (x & y) ^ (^x & z)

	// expected:  01100010
	expected := uint32('b')

	if result != expected {
		t.Errorf("\nexpected value:         %32b\nresult:                 %32b\n", expected, result)
	}
}

//The Majority Function:
//if(x && y || x && z || y && z) then 1 else 0*/
func TestMajority(t *testing.T) {

	// expected: 01100001
	x := uint32('a')

	// expected: 01100010
	y := uint32('b')

	// expected: 01100011
	z := uint32('c')

	/*var result uint32 = 0

	for i := uint32(1); i < uint32(1)<<31; i = i << 1 {
		if x&i == i && y&i == i ||
			x&i == i && z&i == i ||
			y&i == i && z&i == i {
			result += i
		}
	}*/

	result := (x & y) ^ (x & z) ^ (y & z)

	// expected:  01100011
	expected := uint32('c')

	if result != expected {
		t.Errorf("\nexpected value:         %32b\nresult:                 %32b\n", expected, result)
	}
}
