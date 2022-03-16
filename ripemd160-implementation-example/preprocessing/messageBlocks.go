package preprocessing

import "encoding/binary"

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

func BuildMessageBlocks(message []byte) []byte {

	var messageBlock []byte

	messageBlock = append(messageBlock, message...)

	messageBlock = AddZeroPaddingBits(messageBlock, message)

	messageBlock = AddMessageLength(messageBlock, message)

	return messageBlock
}
