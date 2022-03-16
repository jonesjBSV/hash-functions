package preprocessing

import "encoding/binary"
import "sha256-example/logicalFunctions"

func BuildMessageSchedule(messageBlock []byte) [64]uint32 {

	schedule := GetFirst16Words(messageBlock)

	schedule = GetRemaining48Words(schedule)

	return schedule
}

func GetFirst16Words(messageBlock []byte) [64]uint32 {

	var schedule [64]uint32

	j := 0
	for i := 0; i < 64; i += 4 {

		schedule[j] = binary.BigEndian.Uint32(messageBlock[i:])
		j++
	}

	return schedule

}

func GetRemaining48Words(schedule [64]uint32) [64]uint32 {

	for i := 16; i < 64; i++ {
		schedule[i] = logicalFunctions.SmallSigmaOne(schedule[i-2]) +
			schedule[i-7] +
			logicalFunctions.SmallSigmaZero(schedule[i-15]) +
			schedule[i-16]
	}

	return schedule

}
