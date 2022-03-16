package preprocessing

import "testing"
import "encoding/binary"

import "fmt"

func TestBuildMessageBlocks(t *testing.T) {

	message := []byte("abc")

	var messageBlock []byte

	fmt.Printf("Output: %b\n", messageBlock)

	messageBlock = append(messageBlock, message...)

	fmt.Printf("Output: %b\n", messageBlock)

	messageBlock = addZeroPaddingBits(messageBlock, message)

	fmt.Printf("Output: %b\n", messageBlock)

	messageBlock = addMessageLength(messageBlock, message)

	fmt.Printf("Output: %b\n", messageBlock)

	var expected [64]byte

	expected[0] = 'a'
	expected[1] = 'b'
	expected[2] = 'c'
	expected[3] = 0x80
	expected[63] = 24

	var expected1 []byte

	expected1 = append(expected1, 'a')
	expected1 = append(expected1, 'b')
	expected1 = append(expected1, 'c')

	zeros1 := make([]byte, 53)
	zeros1[0] += 0x80

	messageLength := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLength[:], uint64(24))

	expected1 = append(expected1, zeros1...)
	expected1 = append(expected1, messageLength...)

	if expected1[0] != messageBlock[0] {
		t.Errorf("\nExpected: %32b\n Result: %32b\n length: %d", expected1, messageBlock, len(messageBlock))
	}

	if expected1[63] != messageBlock[63] {
		t.Errorf("\nExpected: %32b\n Result: %32b\n length: %d", expected1, messageBlock, len(messageBlock))
	}

}

func addZeroPaddingBits(messageBlock []byte, message []byte) []byte {

	zeros := make([]byte, howMuchPadding(message))

	//add separator bit
	zeros[0] += 0x80

	return append(messageBlock, zeros...)
}

func howMuchPadding(message []byte) int {

	messageLength := len(message)
	result := 0

	if messageLength%64 < 56 {
		result = 56 - messageLength%64
	} else {
		result = 64 + 56 - messageLength%64
	}

	return result
}

func addMessageLength(messageBlock []byte, message []byte) []byte {

	messageLength := len(message) << 3

	messageLengthInBits := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLengthInBits[:], uint64(messageLength))

	return append(messageBlock, messageLengthInBits...)
}
