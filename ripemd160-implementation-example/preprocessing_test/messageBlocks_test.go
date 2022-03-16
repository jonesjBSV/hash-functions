package preprocessing

import "testing"
import "encoding/binary"
import "fmt"

func HowMuchPadding(message []byte) int {

	messageLength := len(message)

	result := 0

	if messageLength%64 < 56 {
		result = 56 - messageLength%64
	} else {
		result = 64 + 56 - messageLength%64
	}

	return result
}

func AddZeroPaddingBits(messageBlock []byte, message []byte) []byte {

	zeros := make([]byte, HowMuchPadding(message))

	//add separator bit
	zeros[0] += 0x80

	return append(messageBlock, zeros...)
}

func AddMessageLength(messageBlock []byte, message []byte) []byte {

	messageLength := len(message) << 3

	messageLengthInBits := make([]byte, 8)
	binary.LittleEndian.PutUint64(messageLengthInBits, uint64(messageLength))

	return append(messageBlock, messageLengthInBits...)
}

func TestBuildMessageBlocks(t *testing.T) {
	message := []byte("abc")

	var messageBlock []byte

	fmt.Printf("Output: %b\n", messageBlock)

	messageBlock = append(messageBlock, message...)

	fmt.Printf("Output: %b\n", messageBlock)

	messageBlock = AddZeroPaddingBits(messageBlock, message)

	fmt.Printf("Output: %b\n", messageBlock)

	messageBlock = AddMessageLength(messageBlock, message)

	fmt.Printf("Output: %b\n", messageBlock)

	var expected1 []byte

	expected1 = append(expected1, 'a')
	expected1 = append(expected1, 'b')
	expected1 = append(expected1, 'c')

	zeros1 := make([]byte, 53)
	zeros1[0] += 0x80

	messageLength := make([]byte, 8)
	binary.LittleEndian.PutUint64(messageLength[:], uint64(24))

	expected1 = append(expected1, zeros1...)
	expected1 = append(expected1, messageLength...)

	if fmt.Sprintf("%x", expected1) != fmt.Sprintf("%x", messageBlock) {
		t.Errorf("\nExpected: %b\n Result: %b\n length: %d", expected1, messageBlock, len(messageBlock))
	}

}
