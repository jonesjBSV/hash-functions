package compression

import "math/bits"

var word []int = []int{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	7, 4, 13, 1, 10, 6, 15, 3, 12, 0, 9, 5, 2, 14, 11, 8,
	3, 10, 14, 4, 9, 15, 8, 1, 2, 7, 0, 6, 13, 11, 5, 12,
	1, 9, 11, 10, 0, 8, 12, 4, 13, 3, 7, 15, 14, 5, 6, 2,
	4, 0, 5, 9, 7, 12, 2, 10, 14, 1, 3, 8, 11, 6, 15, 13}

var rotationLeft []int = []int{
	11, 14, 15, 12, 5, 8, 7, 9, 11, 13, 14, 15, 6, 7, 9, 8,
	7, 6, 8, 13, 11, 9, 7, 15, 7, 12, 15, 9, 11, 7, 13, 12,
	11, 13, 6, 7, 14, 9, 13, 15, 14, 8, 13, 6, 5, 12, 7, 5,
	11, 12, 14, 15, 14, 15, 9, 8, 9, 14, 5, 6, 8, 6, 5, 12,
	9, 15, 5, 11, 6, 8, 13, 12, 5, 12, 13, 14, 11, 8, 5, 6}

var wordPrime []int = []int{
	5, 14, 7, 0, 9, 2, 11, 4, 13, 6, 15, 8, 1, 10, 3, 12,
	6, 11, 3, 7, 0, 13, 5, 10, 14, 15, 8, 12, 4, 9, 1, 2,
	15, 5, 1, 3, 7, 14, 6, 9, 11, 8, 12, 2, 10, 0, 4, 13,
	8, 6, 4, 1, 3, 11, 15, 0, 5, 12, 2, 13, 9, 7, 10, 14,
	12, 15, 10, 4, 1, 5, 8, 7, 6, 2, 13, 14, 0, 3, 9, 11}

var rotationLeftPrime []int = []int{
	8, 9, 9, 11, 13, 15, 15, 5, 7, 7, 8, 11, 14, 14, 12, 6,
	9, 13, 15, 7, 12, 8, 9, 11, 7, 7, 12, 7, 6, 15, 13, 11,
	9, 7, 15, 11, 8, 6, 6, 14, 12, 13, 5, 14, 13, 13, 7, 5,
	15, 5, 8, 11, 14, 14, 6, 14, 6, 9, 12, 9, 12, 5, 15, 8,
	8, 5, 12, 9, 12, 5, 14, 6, 8, 13, 6, 5, 15, 13, 11, 11}

var InitialValues []uint32 = []uint32{0x67452301, 0xefcdab89, 0x98badcfe,
	0x10325476, 0xc3d2e1f0}

func Compression(messageSchedule [16]uint32, chainingValues []uint32) []uint32 {

	var h0, h1, h2, h3, h4 uint32 = chainingValues[0], chainingValues[1],
		chainingValues[2], chainingValues[3], chainingValues[4]

	var a, b, c, d, e = h0, h1, h2, h3, h4
	var aa, bb, cc, dd, ee = h0, h1, h2, h3, h4

	var tempWord, tempWordPrime uint32

	//first round
	for i := 0; i < 16; i++ {

		tempWord = a + (b ^ c ^ d) + messageSchedule[word[i]] + 0x00000000
		tempWord = bits.RotateLeft32(tempWord, rotationLeft[i]) + e
		a, e, d, c, b = e, d, bits.RotateLeft32(c, 10), b, tempWord

		//first round prime
		tempWordPrime = aa + (bb ^ (cc | ^dd)) + messageSchedule[wordPrime[i]] + 0x50a28be6
		tempWordPrime = bits.RotateLeft32(tempWordPrime, rotationLeftPrime[i]) + ee
		aa, ee, dd, cc, bb = ee, dd, bits.RotateLeft32(cc, 10), bb, tempWordPrime
	}

	//second round
	for i := 16; i < 32; i++ {
		tempWord = a + (b&c | ^b&d) + messageSchedule[word[i]] + 0x5a827999
		tempWord = bits.RotateLeft32(tempWord, rotationLeft[i]) + e
		a, e, d, c, b = e, d, bits.RotateLeft32(c, 10), b, tempWord

		//second round prime
		tempWordPrime = aa + (bb&dd | cc&^dd) + messageSchedule[wordPrime[i]] + 0x5c4dd124
		tempWordPrime = bits.RotateLeft32(tempWordPrime, rotationLeftPrime[i]) + ee
		aa, ee, dd, cc, bb = ee, dd, bits.RotateLeft32(cc, 10), bb, tempWordPrime
	}

	//third round
	for i := 32; i < 48; i++ {
		tempWord = a + (b | ^c ^ d) + messageSchedule[word[i]] + 0x6ed9eba1
		tempWord = bits.RotateLeft32(tempWord, rotationLeft[i]) + e
		a, e, d, c, b = e, d, bits.RotateLeft32(c, 10), b, tempWord

		//third round prime
		tempWordPrime = aa + (bb | ^cc ^ dd) + messageSchedule[wordPrime[i]] + 0x6d703ef3
		tempWordPrime = bits.RotateLeft32(tempWordPrime, rotationLeftPrime[i]) + ee
		aa, ee, dd, cc, bb = ee, dd, bits.RotateLeft32(cc, 10), bb, tempWordPrime
	}

	//fourth round
	for i := 48; i < 64; i++ {
		tempWord = a + (b&d | c&^d) + messageSchedule[word[i]] + 0x8f1bbcdc
		tempWord = bits.RotateLeft32(tempWord, rotationLeft[i]) + e
		a, e, d, c, b = e, d, bits.RotateLeft32(c, 10), b, tempWord

		//fourth round prime
		tempWordPrime = aa + (bb&cc | ^bb&dd) + messageSchedule[wordPrime[i]] + 0x7a6d76e9
		tempWordPrime = bits.RotateLeft32(tempWordPrime, rotationLeftPrime[i]) + ee
		aa, ee, dd, cc, bb = ee, dd, bits.RotateLeft32(cc, 10), bb, tempWordPrime
	}

	//fifth round
	for i := 64; i < 80; i++ {
		tempWord = a + (b ^ (c | ^d)) + messageSchedule[word[i]] + 0xa953fd4e
		tempWord = bits.RotateLeft32(tempWord, rotationLeft[i]) + e
		a, e, d, c, b = e, d, bits.RotateLeft32(c, 10), b, tempWord

		//fifthround prime
		tempWordPrime = aa + (bb ^ cc ^ dd) + messageSchedule[wordPrime[i]] + 0x00000000
		tempWordPrime = bits.RotateLeft32(tempWordPrime, rotationLeftPrime[i]) + ee
		aa, ee, dd, cc, bb = ee, dd, bits.RotateLeft32(cc, 10), bb, tempWordPrime
	}

	tempValue := h1 + c + dd
	h1 = h2 + d + ee
	h2 = h3 + e + aa
	h3 = h4 + a + bb
	h4 = h0 + b + cc
	h0 = tempValue

	chainingValues[0] = h0
	chainingValues[1] = h1
	chainingValues[2] = h2
	chainingValues[3] = h3
	chainingValues[4] = h4

	return chainingValues

}
