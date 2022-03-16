package preprocessing

import (
	"encoding/binary"
	"sha256-example/logicalFunctions"
	"testing"
)

func TestBuildMessageSchedule(t *testing.T) {

	messageSchedule := getFirst16Words(BuildMessageBlocks([]byte("abc")))

	t.Logf("Output: %b\n", messageSchedule)

	messageSchedule = getRemaining48Words(messageSchedule)

	t.Logf("Output: %b\n", messageSchedule)

}

func getFirst16Words(messageBlock []byte) [64]uint32 {

	var schedule [64]uint32

	j := 0
	for i := 0; i < 64; i += 4 {

		schedule[j] = binary.BigEndian.Uint32(messageBlock[i:])
		j++
	}

	return schedule

}

func getRemaining48Words(schedule [64]uint32) [64]uint32 {

	for i := 16; i < 64; i++ {
		schedule[i] = logicalFunctions.SmallSigmaOne(schedule[i-2]) + schedule[i-7] + logicalFunctions.SmallSigmaZero(schedule[i-15]) + schedule[i-16]
	}

	return schedule

}
